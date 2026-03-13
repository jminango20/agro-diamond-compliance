// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibERC1155Storage, ERC1155Storage, PartitionBalance} from "../../storage/LibERC1155Storage.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";
import {LibAccessStorage, AccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibAppStorage, AppStorage} from "../../libraries/LibAppStorage.sol";
import {LibFreezeStorage, FreezeStorage} from "../../storage/LibFreezeStorage.sol";
import {LibDiamond} from "../../libraries/LibDiamond.sol";
import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";

/**
 * @title SupplyFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Mint, burn, and forced transfer for regulated ERC-1155 security tokens.
 *         All supply mutations enforce the regulatory model:
 *         - mint: supplyCap, asset pause, receiver freeze checks
 *         - burn: asset pause, sufficient free balance
 *         - forcedTransfer: regulatory override that bypasses compliance but respects pause
 * @dev Only ISSUER_ROLE (or Diamond owner) can mint/burn.
 *      Only TRANSFER_AGENT (or Diamond owner) can forcedTransfer.
 *      Holder count tracking via SupplyStorage for on-chain analytics.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract SupplyFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    /// @dev ERC-1155 standard event — emitted on mint, burn, and forcedTransfer
    ///      so that explorers and indexers recognise token movements.
    event TransferSingle(
        address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value
    );

    event Minted(uint256 indexed tokenId, address indexed to, uint256 amount);
    event Burned(uint256 indexed tokenId, address indexed from, uint256 amount);
    event ForcedTransfer(
        uint256 indexed tokenId, address indexed from, address indexed to, uint256 amount, bytes32 reasonCode
    );

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error SupplyFacet__Unauthorized();
    error SupplyFacet__AssetNotRegistered(uint256 tokenId);
    error SupplyFacet__AssetPaused(uint256 tokenId);
    error SupplyFacet__ProtocolPaused();
    error SupplyFacet__MintToZeroAddress();
    error SupplyFacet__BurnFromZeroAddress();
    error SupplyFacet__TransferToZeroAddress();
    error SupplyFacet__SupplyCapExceeded(uint256 tokenId, uint256 currentSupply, uint256 cap);
    error SupplyFacet__InsufficientFreeBalance(uint256 tokenId, address account, uint256 available, uint256 required);
    error SupplyFacet__WalletFrozenGlobal(address wallet);
    error SupplyFacet__WalletFrozenAsset(uint256 tokenId, address wallet);
    error SupplyFacet__ArrayLengthMismatch();

    /*//////////////////////////////////////////////////////////////
                            ROLE CONSTANTS
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Mints `amount` of `tokenId` to `to`
    /// @param tokenId The asset class to mint
    /// @param to Receiver address
    /// @param amount Amount to mint
    function mint(uint256 tokenId, address to, uint256 amount) external {
        _enforceIssuerOrOwner();
        _validateMint(tokenId, to, amount);
        _executeMint(tokenId, to, amount);
    }

    /// @notice Batch mint: multiple tokenIds to multiple receivers
    /// @param tokenIds Array of asset classes
    /// @param recipients Array of receiver addresses
    /// @param amounts Array of amounts to mint
    function batchMint(uint256[] calldata tokenIds, address[] calldata recipients, uint256[] calldata amounts)
        external
    {
        if (tokenIds.length != recipients.length || tokenIds.length != amounts.length) {
            revert SupplyFacet__ArrayLengthMismatch();
        }
        _enforceIssuerOrOwner();
        for (uint256 i; i < tokenIds.length; ++i) {
            _validateMint(tokenIds[i], recipients[i], amounts[i]);
            _executeMint(tokenIds[i], recipients[i], amounts[i]);
        }
    }

    /// @notice Burns `amount` of `tokenId` from `from`
    /// @param tokenId The asset class to burn
    /// @param from Address to burn from
    /// @param amount Amount to burn
    function burn(uint256 tokenId, address from, uint256 amount) external {
        _enforceIssuerOrOwner();
        _validateBurn(tokenId, from, amount);
        _executeBurn(tokenId, from, amount);
    }

    /// @notice Forced transfer for regulatory purposes (recovery, seizure)
    /// @dev Bypasses compliance module canTransfer check but respects pause
    /// @param tokenId The asset class to transfer
    /// @param from Source address
    /// @param to Destination address
    /// @param amount Amount to transfer
    /// @param reasonCode Why this forced transfer is happening
    function forcedTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes32 reasonCode) external {
        _enforceTransferAgentOrOwner();
        if (to == address(0)) revert SupplyFacet__TransferToZeroAddress();

        AssetConfig storage config = _enforceAssetActive(tokenId);

        // Forced transfers only check free balance (no freeze/lockup bypass needed —
        // the agent IS the authority that imposed those restrictions)
        ERC1155Storage storage es = LibERC1155Storage.layout();
        uint256 freeBalance = es.partitions[tokenId][from].free;
        if (freeBalance < amount) {
            revert SupplyFacet__InsufficientFreeBalance(tokenId, from, freeBalance, amount);
        }

        es.partitions[tokenId][from].free -= amount;
        es.partitions[tokenId][to].free += amount;

        _updateHolderTracking(tokenId, from, to);

        emit TransferSingle(msg.sender, from, to, tokenId, amount);
        emit ForcedTransfer(tokenId, from, to, amount, reasonCode);

        // Post-hook: all compliance modules get notified even for forced transfers
        address[] storage modules = config.complianceModules;
        uint256 mLen = modules.length;
        for (uint256 i; i < mLen;) {
            IComplianceModule(modules[i]).transferred(tokenId, from, to, amount);
            unchecked { ++i; }
        }
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns total minted supply for a tokenId
    function totalSupply(uint256 tokenId) external view returns (uint256) {
        return LibAssetStorage.layout().configs[tokenId].totalSupply;
    }

    /// @notice Returns unique holder count for a tokenId
    function holderCount(uint256 tokenId) external view returns (uint256) {
        return LibSupplyStorage.layout().holderCount[tokenId];
    }

    /// @notice Returns true if `account` holds any balance of `tokenId`
    function isHolder(uint256 tokenId, address account) external view returns (bool) {
        return LibSupplyStorage.layout().isHolder[tokenId][account];
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — VALIDATION
    //////////////////////////////////////////////////////////////*/

    function _validateMint(uint256 tokenId, address to, uint256 amount) internal view {
        if (to == address(0)) revert SupplyFacet__MintToZeroAddress();

        AssetConfig storage config = _enforceAssetActive(tokenId);

        // Supply cap check (0 = unlimited)
        if (config.supplyCap != 0 && config.totalSupply + amount > config.supplyCap) {
            revert SupplyFacet__SupplyCapExceeded(tokenId, config.totalSupply, config.supplyCap);
        }

        // Receiver freeze checks
        FreezeStorage storage fs = LibFreezeStorage.layout();
        if (fs.globalFreeze[to]) revert SupplyFacet__WalletFrozenGlobal(to);
        if (fs.assetFreeze[tokenId][to]) revert SupplyFacet__WalletFrozenAsset(tokenId, to);
    }

    function _validateBurn(uint256 tokenId, address from, uint256 amount) internal view {
        if (from == address(0)) revert SupplyFacet__BurnFromZeroAddress();

        _enforceAssetActive(tokenId);

        // Only burn from free balance
        ERC1155Storage storage es = LibERC1155Storage.layout();
        uint256 freeBalance = es.partitions[tokenId][from].free;
        FreezeStorage storage fs = LibFreezeStorage.layout();
        uint256 frozen = fs.frozenAmount[tokenId][from];
        uint256 available = freeBalance > frozen ? freeBalance - frozen : 0;
        if (available < amount) {
            revert SupplyFacet__InsufficientFreeBalance(tokenId, from, available, amount);
        }
    }

    /// @dev Checks protocol pause + asset existence + asset pause
    function _enforceAssetActive(uint256 tokenId) internal view returns (AssetConfig storage config) {
        AppStorage storage app = LibAppStorage.layout();
        if (app.globalPaused) revert SupplyFacet__ProtocolPaused();

        AssetStorage storage as_ = LibAssetStorage.layout();
        config = as_.configs[tokenId];
        if (!config.exists) revert SupplyFacet__AssetNotRegistered(tokenId);
        if (config.paused) revert SupplyFacet__AssetPaused(tokenId);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — EXECUTE
    //////////////////////////////////////////////////////////////*/

    function _executeMint(uint256 tokenId, address to, uint256 amount) internal {
        // Update balances
        ERC1155Storage storage es = LibERC1155Storage.layout();
        es.partitions[tokenId][to].free += amount;

        // Update supply
        AssetConfig storage config = LibAssetStorage.layout().configs[tokenId];
        config.totalSupply += amount;

        // Holder tracking
        SupplyStorage storage ss = LibSupplyStorage.layout();
        if (!ss.isHolder[tokenId][to]) {
            ss.isHolder[tokenId][to] = true;
            ss.holderCount[tokenId] += 1;
        }

        emit TransferSingle(msg.sender, address(0), to, tokenId, amount);
        emit Minted(tokenId, to, amount);

        // Compliance post-hooks
        address[] storage modules = config.complianceModules;
        uint256 mLen = modules.length;
        for (uint256 i; i < mLen;) {
            IComplianceModule(modules[i]).minted(tokenId, to, amount);
            unchecked { ++i; }
        }
    }

    function _executeBurn(uint256 tokenId, address from, uint256 amount) internal {
        // Update balances
        ERC1155Storage storage es = LibERC1155Storage.layout();
        es.partitions[tokenId][from].free -= amount;

        // Update supply
        AssetConfig storage config = LibAssetStorage.layout().configs[tokenId];
        config.totalSupply -= amount;

        // Holder tracking — remove if all partitions are zero
        PartitionBalance storage pb = es.partitions[tokenId][from];
        if (pb.free == 0 && pb.locked == 0 && pb.custody == 0 && pb.pendingSettlement == 0) {
            SupplyStorage storage ss = LibSupplyStorage.layout();
            ss.isHolder[tokenId][from] = false;
            ss.holderCount[tokenId] -= 1;
        }

        emit TransferSingle(msg.sender, from, address(0), tokenId, amount);
        emit Burned(tokenId, from, amount);

        // Compliance post-hooks
        address[] storage modules = config.complianceModules;
        uint256 mLen = modules.length;
        for (uint256 i; i < mLen;) {
            IComplianceModule(modules[i]).burned(tokenId, from, amount);
            unchecked { ++i; }
        }
    }

    /// @dev Updates holder tracking for forced transfers
    function _updateHolderTracking(uint256 tokenId, address from, address to) internal {
        ERC1155Storage storage es = LibERC1155Storage.layout();
        SupplyStorage storage ss = LibSupplyStorage.layout();

        // Add receiver as holder if new
        if (!ss.isHolder[tokenId][to]) {
            ss.isHolder[tokenId][to] = true;
            ss.holderCount[tokenId] += 1;
        }

        // Remove sender if all partitions are zero
        PartitionBalance storage pb = es.partitions[tokenId][from];
        if (pb.free == 0 && pb.locked == 0 && pb.custody == 0 && pb.pendingSettlement == 0) {
            ss.isHolder[tokenId][from] = false;
            ss.holderCount[tokenId] -= 1;
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — ACCESS CONTROL
    //////////////////////////////////////////////////////////////*/

    function _enforceIssuerOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isIssuer = LibAccessStorage.layout().roles[ISSUER_ROLE][msg.sender];
        if (!isOwner && !isIssuer) revert SupplyFacet__Unauthorized();
    }

    function _enforceTransferAgentOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isAgent = LibAccessStorage.layout().roles[TRANSFER_AGENT][msg.sender];
        if (!isOwner && !isAgent) revert SupplyFacet__Unauthorized();
    }
}
