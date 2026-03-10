// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibSnapshotStorage, SnapshotStorage, Snapshot, HolderSnapshot} from "../../storage/LibSnapshotStorage.sol";
import {LibAccessStorage, AccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibDiamond} from "../../libraries/LibDiamond.sol";

/// @dev Dividend distribution record
struct DividendStorage {
    /// dividendId → Dividend
    mapping(uint256 => Dividend) dividends;
    /// next dividend ID
    uint256 nextDividendId;
    /// tokenId → array of dividend IDs
    mapping(uint256 => uint256[]) tokenDividends;
}

struct Dividend {
    uint256 id;
    uint256 snapshotId;        // which snapshot to use for pro-rata
    uint256 tokenId;           // which asset class
    uint256 totalAmount;       // total distribution amount (in wei of payment token)
    address paymentToken;      // ERC-20 token used for payment (address(0) = ETH)
    uint256 claimedAmount;     // total claimed so far
    uint64 createdAt;
    bool exists;
    mapping(address => bool) claimed;  // holder → has claimed
}

/// @title LibDividendStorage
/// @notice Namespaced dividend storage for the Diamond.
///         slot = keccak256("diamond.rwa.dividend.storage") - 1
library LibDividendStorage {
    // solhint-disable-next-line no-inline-assembly
    bytes32 internal constant POSITION =
        0xa1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2;

    function layout() internal pure returns (DividendStorage storage s) {
        bytes32 position = POSITION;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            s.slot := position
        }
    }
}

/**
 * @title DividendFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Distributes dividends to holders based on snapshot balances.
 *         Supports ETH and ERC-20 payment tokens. Holders claim their
 *         pro-rata share: `holderBalance / totalSupply * totalAmount`.
 * @dev Requires a snapshot to be created and holders recorded before
 *      creating a dividend. Uses pull-pattern (holders claim) to avoid
 *      gas issues with large holder sets.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract DividendFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event DividendCreated(
        uint256 indexed dividendId,
        uint256 indexed tokenId,
        uint256 indexed snapshotId,
        uint256 totalAmount,
        address paymentToken
    );
    event DividendClaimed(uint256 indexed dividendId, address indexed holder, uint256 amount);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error DividendFacet__Unauthorized();
    error DividendFacet__SnapshotNotFound(uint256 snapshotId);
    error DividendFacet__DividendNotFound(uint256 dividendId);
    error DividendFacet__ZeroAmount();
    error DividendFacet__AlreadyClaimed(uint256 dividendId, address holder);
    error DividendFacet__HolderNotRecorded(uint256 dividendId, address holder);
    error DividendFacet__NothingToClaim(uint256 dividendId, address holder);
    error DividendFacet__TransferFailed();
    error DividendFacet__InsufficientETH();

    /*//////////////////////////////////////////////////////////////
                            ROLE CONSTANT
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Creates a dividend distribution backed by a snapshot
    /// @dev For ETH dividends: send msg.value >= totalAmount.
    ///      For ERC-20: caller must have approved this Diamond for totalAmount.
    /// @param snapshotId The snapshot to use for pro-rata calculation
    /// @param totalAmount Total amount to distribute
    /// @param paymentToken ERC-20 address (address(0) for ETH)
    /// @return dividendId The created dividend ID
    function createDividend(uint256 snapshotId, uint256 totalAmount, address paymentToken)
        external
        payable
        returns (uint256 dividendId)
    {
        _enforceIssuerOrOwner();
        if (totalAmount == 0) revert DividendFacet__ZeroAmount();

        Snapshot storage snap = LibSnapshotStorage.layout().snapshots[snapshotId];
        if (!snap.exists) revert DividendFacet__SnapshotNotFound(snapshotId);

        // Collect funds
        if (paymentToken == address(0)) {
            if (msg.value < totalAmount) revert DividendFacet__InsufficientETH();
        } else {
            // Pull ERC-20 from caller
            (bool ok, bytes memory data) = paymentToken.call(
                abi.encodeWithSignature("transferFrom(address,address,uint256)", msg.sender, address(this), totalAmount)
            );
            if (!ok || (data.length != 0 && !abi.decode(data, (bool)))) {
                revert DividendFacet__TransferFailed();
            }
        }

        DividendStorage storage ds = LibDividendStorage.layout();
        dividendId = ds.nextDividendId++;

        Dividend storage div = ds.dividends[dividendId];
        div.id = dividendId;
        div.snapshotId = snapshotId;
        div.tokenId = snap.tokenId;
        div.totalAmount = totalAmount;
        div.paymentToken = paymentToken;
        div.createdAt = uint64(block.timestamp);
        div.exists = true;

        ds.tokenDividends[snap.tokenId].push(dividendId);

        emit DividendCreated(dividendId, snap.tokenId, snapshotId, totalAmount, paymentToken);
    }

    /// @notice Claims dividend for msg.sender
    /// @param dividendId The dividend to claim from
    function claimDividend(uint256 dividendId) external {
        DividendStorage storage ds = LibDividendStorage.layout();
        Dividend storage div = ds.dividends[dividendId];
        if (!div.exists) revert DividendFacet__DividendNotFound(dividendId);
        if (div.claimed[msg.sender]) revert DividendFacet__AlreadyClaimed(dividendId, msg.sender);

        Snapshot storage snap = LibSnapshotStorage.layout().snapshots[div.snapshotId];
        HolderSnapshot storage hs = snap.balances[msg.sender];
        if (!hs.recorded) revert DividendFacet__HolderNotRecorded(dividendId, msg.sender);
        if (hs.balance == 0) revert DividendFacet__NothingToClaim(dividendId, msg.sender);

        // Pro-rata: holderBalance * totalAmount / totalSupply
        uint256 amount = (hs.balance * div.totalAmount) / snap.totalSupply;
        if (amount == 0) revert DividendFacet__NothingToClaim(dividendId, msg.sender);

        div.claimed[msg.sender] = true;
        div.claimedAmount += amount;

        // Transfer
        if (div.paymentToken == address(0)) {
            (bool ok,) = msg.sender.call{value: amount}("");
            if (!ok) revert DividendFacet__TransferFailed();
        } else {
            (bool ok, bytes memory data) = div.paymentToken.call(
                abi.encodeWithSignature("transfer(address,uint256)", msg.sender, amount)
            );
            if (!ok || (data.length != 0 && !abi.decode(data, (bool)))) {
                revert DividendFacet__TransferFailed();
            }
        }

        emit DividendClaimed(dividendId, msg.sender, amount);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns dividend metadata
    function getDividend(uint256 dividendId)
        external
        view
        returns (
            uint256 snapshotId,
            uint256 tokenId,
            uint256 totalAmount,
            address paymentToken,
            uint256 claimedAmount,
            uint64 createdAt
        )
    {
        Dividend storage div = _enforceDividendExists(dividendId);
        snapshotId = div.snapshotId;
        tokenId = div.tokenId;
        totalAmount = div.totalAmount;
        paymentToken = div.paymentToken;
        claimedAmount = div.claimedAmount;
        createdAt = div.createdAt;
    }

    /// @notice Returns true if holder has claimed a dividend
    function hasClaimed(uint256 dividendId, address holder) external view returns (bool) {
        return _enforceDividendExists(dividendId).claimed[holder];
    }

    /// @notice Returns the claimable amount for a holder (0 if already claimed or not recorded)
    function claimableAmount(uint256 dividendId, address holder) external view returns (uint256) {
        Dividend storage div = _enforceDividendExists(dividendId);
        if (div.claimed[holder]) return 0;

        Snapshot storage snap = LibSnapshotStorage.layout().snapshots[div.snapshotId];
        HolderSnapshot storage hs = snap.balances[holder];
        if (!hs.recorded || hs.balance == 0 || snap.totalSupply == 0) return 0;

        return (hs.balance * div.totalAmount) / snap.totalSupply;
    }

    /// @notice Returns all dividend IDs for a tokenId
    function getTokenDividends(uint256 tokenId) external view returns (uint256[] memory) {
        return LibDividendStorage.layout().tokenDividends[tokenId];
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _enforceDividendExists(uint256 dividendId) internal view returns (Dividend storage div) {
        div = LibDividendStorage.layout().dividends[dividendId];
        if (!div.exists) revert DividendFacet__DividendNotFound(dividendId);
    }

    function _enforceIssuerOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isIssuer = LibAccessStorage.layout().roles[ISSUER_ROLE][msg.sender];
        if (!isOwner && !isIssuer) revert DividendFacet__Unauthorized();
    }

    receive() external payable {}
}
