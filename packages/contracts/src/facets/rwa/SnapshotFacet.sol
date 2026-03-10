// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibSnapshotStorage, SnapshotStorage, Snapshot, HolderSnapshot} from "../../storage/LibSnapshotStorage.sol";
import {LibERC1155Storage, ERC1155Storage, PartitionBalance} from "../../storage/LibERC1155Storage.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";
import {LibAccessStorage, AccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibDiamond} from "../../libraries/LibDiamond.sol";

/**
 * @title SnapshotFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Creates balance snapshots for registered tokenIds.
 *         Snapshots capture holder balances at a point in time for
 *         dividend distribution, voting, and regulatory reporting.
 * @dev Snapshot balances are recorded lazily via `recordHolder` or
 *      eagerly via `recordHoldersBatch`. The snapshot stores total
 *      balance (free + locked + custody + pending) since all partitions
 *      are entitled to corporate actions.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract SnapshotFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event SnapshotCreated(uint256 indexed snapshotId, uint256 indexed tokenId, uint256 totalSupply, uint64 timestamp);
    event HolderRecorded(uint256 indexed snapshotId, address indexed holder, uint256 balance);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error SnapshotFacet__Unauthorized();
    error SnapshotFacet__AssetNotRegistered(uint256 tokenId);
    error SnapshotFacet__SnapshotNotFound(uint256 snapshotId);
    error SnapshotFacet__HolderAlreadyRecorded(uint256 snapshotId, address holder);

    /*//////////////////////////////////////////////////////////////
                            ROLE CONSTANT
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Creates a new snapshot for a tokenId
    /// @param tokenId The asset class to snapshot
    /// @return snapshotId The ID of the created snapshot
    function createSnapshot(uint256 tokenId) external returns (uint256 snapshotId) {
        _enforceIssuerOrOwner();

        AssetConfig storage config = LibAssetStorage.layout().configs[tokenId];
        if (!config.exists) revert SnapshotFacet__AssetNotRegistered(tokenId);

        SnapshotStorage storage ss = LibSnapshotStorage.layout();
        snapshotId = ss.nextSnapshotId++;

        Snapshot storage snap = ss.snapshots[snapshotId];
        snap.id = snapshotId;
        snap.tokenId = tokenId;
        snap.totalSupply = config.totalSupply;
        snap.timestamp = uint64(block.timestamp);
        snap.holderCount = LibSupplyStorage.layout().holderCount[tokenId];
        snap.exists = true;

        ss.tokenSnapshots[tokenId].push(snapshotId);

        emit SnapshotCreated(snapshotId, tokenId, config.totalSupply, uint64(block.timestamp));
    }

    /// @notice Records a holder's balance in a snapshot
    /// @dev Captures total balance (all partitions) at call time.
    ///      Should be called shortly after createSnapshot.
    /// @param snapshotId The snapshot to record into
    /// @param holder The holder address to record
    function recordHolder(uint256 snapshotId, address holder) external {
        _enforceIssuerOrOwner();
        _recordHolder(snapshotId, holder);
    }

    /// @notice Records multiple holders in a snapshot (batch)
    /// @param snapshotId The snapshot to record into
    /// @param holders Array of holder addresses
    function recordHoldersBatch(uint256 snapshotId, address[] calldata holders) external {
        _enforceIssuerOrOwner();
        for (uint256 i; i < holders.length; ++i) {
            _recordHolder(snapshotId, holders[i]);
        }
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns snapshot metadata
    function getSnapshot(uint256 snapshotId)
        external
        view
        returns (uint256 tokenId, uint256 totalSupply, uint64 timestamp, uint256 holderCount)
    {
        Snapshot storage snap = _enforceSnapshotExists(snapshotId);
        tokenId = snap.tokenId;
        totalSupply = snap.totalSupply;
        timestamp = snap.timestamp;
        holderCount = snap.holderCount;
    }

    /// @notice Returns a holder's balance at snapshot time
    function getSnapshotBalance(uint256 snapshotId, address holder)
        external
        view
        returns (uint256 balance, bool recorded)
    {
        _enforceSnapshotExists(snapshotId);
        HolderSnapshot storage hs = LibSnapshotStorage.layout().snapshots[snapshotId].balances[holder];
        balance = hs.balance;
        recorded = hs.recorded;
    }

    /// @notice Returns all snapshot IDs for a tokenId
    function getTokenSnapshots(uint256 tokenId) external view returns (uint256[] memory) {
        return LibSnapshotStorage.layout().tokenSnapshots[tokenId];
    }

    /// @notice Returns the latest snapshot ID for a tokenId (or reverts if none)
    function getLatestSnapshotId(uint256 tokenId) external view returns (uint256) {
        uint256[] storage ids = LibSnapshotStorage.layout().tokenSnapshots[tokenId];
        if (ids.length == 0) revert SnapshotFacet__SnapshotNotFound(0);
        return ids[ids.length - 1];
    }

    /// @notice Returns the next snapshot ID that will be assigned
    function nextSnapshotId() external view returns (uint256) {
        return LibSnapshotStorage.layout().nextSnapshotId;
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _recordHolder(uint256 snapshotId, address holder) internal {
        Snapshot storage snap = _enforceSnapshotExists(snapshotId);
        if (snap.balances[holder].recorded) {
            revert SnapshotFacet__HolderAlreadyRecorded(snapshotId, holder);
        }

        ERC1155Storage storage es = LibERC1155Storage.layout();
        PartitionBalance storage pb = es.partitions[snap.tokenId][holder];
        uint256 totalBalance = pb.free + pb.locked + pb.custody + pb.pendingSettlement;

        snap.balances[holder] = HolderSnapshot({balance: totalBalance, recorded: true});

        emit HolderRecorded(snapshotId, holder, totalBalance);
    }

    function _enforceSnapshotExists(uint256 snapshotId) internal view returns (Snapshot storage snap) {
        snap = LibSnapshotStorage.layout().snapshots[snapshotId];
        if (!snap.exists) revert SnapshotFacet__SnapshotNotFound(snapshotId);
    }

    function _enforceIssuerOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isIssuer = LibAccessStorage.layout().roles[ISSUER_ROLE][msg.sender];
        if (!isOwner && !isIssuer) revert SnapshotFacet__Unauthorized();
    }
}
