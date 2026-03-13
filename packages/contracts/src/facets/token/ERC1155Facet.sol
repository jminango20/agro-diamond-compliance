// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibERC1155Storage, ERC1155Storage, PartitionBalance} from "../../storage/LibERC1155Storage.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibAppStorage, AppStorage} from "../../libraries/LibAppStorage.sol";
import {LibFreezeStorage, FreezeStorage} from "../../storage/LibFreezeStorage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";
import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {IERC1155Receiver} from "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";

/**
 * @title ERC1155Facet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice ERC-1155 multi-token with integrated regulatory validation.
 *         All transfers funnel through `_validateMovement` which enforces
 *         the three-level regulatory model (global → per-tokenId → per-holder).
 * @dev Partition-aware: transfers deduct from `free` balance only.
 *      Frozen amounts, locked, custody, and pendingSettlement are excluded.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract ERC1155Facet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event TransferSingle(
        address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value
    );
    event TransferBatch(
        address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values
    );
    event ApprovalForAll(address indexed account, address indexed operator, bool approved);

    event RegulatoryTransfer(
        uint256 indexed tokenId, address indexed from, address indexed to, uint256 amount, bytes32 reasonCode
    );

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error ERC1155Facet__NotApprovedOrOwner();
    error ERC1155Facet__TransferToZeroAddress();
    error ERC1155Facet__InsufficientFreeBalance(uint256 tokenId, address account, uint256 available, uint256 required);
    error ERC1155Facet__ArrayLengthMismatch();
    error ERC1155Facet__ProtocolPaused();
    error ERC1155Facet__AssetPaused(uint256 tokenId);
    error ERC1155Facet__AssetNotRegistered(uint256 tokenId);
    error ERC1155Facet__WalletFrozenGlobal(address wallet);
    error ERC1155Facet__WalletFrozenAsset(uint256 tokenId, address wallet);
    error ERC1155Facet__LockupActive(uint256 tokenId, address wallet, uint64 expiry);
    error ERC1155Facet__ComplianceRejected(uint256 tokenId, bytes32 reason);
    error ERC1155Facet__SelfApproval();
    error ERC1155Facet__InvalidReceiver(address to);

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Transfers `amount` of `id` from `from` to `to`
    function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes calldata data) external {
        _validateAndTransferSingle(msg.sender, from, to, id, amount, data);
    }

    /// @notice Batch transfer of multiple token types
    function safeBatchTransferFrom(
        address from,
        address to,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        bytes calldata data
    ) external {
        if (ids.length != amounts.length) revert ERC1155Facet__ArrayLengthMismatch();
        _validateAndTransferBatch(msg.sender, from, to, ids, amounts, data);
    }

    /// @notice Sets approval for `operator` to manage all of caller's tokens
    function setApprovalForAll(address operator, bool approved) external {
        if (operator == msg.sender) revert ERC1155Facet__SelfApproval();
        LibERC1155Storage.layout().operatorApprovals[msg.sender][operator] = approved;
        emit ApprovalForAll(msg.sender, operator, approved);
    }

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the free balance of `account` for token `id`
    /// @dev Only the `free` partition is returned — locked/custody/pending are excluded
    function balanceOf(address account, uint256 id) external view returns (uint256) {
        return LibERC1155Storage.layout().partitions[id][account].free;
    }

    /// @notice Batch balance query
    function balanceOfBatch(address[] calldata accounts, uint256[] calldata ids)
        external
        view
        returns (uint256[] memory balances)
    {
        if (accounts.length != ids.length) revert ERC1155Facet__ArrayLengthMismatch();
        ERC1155Storage storage es = LibERC1155Storage.layout();
        balances = new uint256[](accounts.length);
        for (uint256 i; i < accounts.length; ++i) {
            balances[i] = es.partitions[ids[i]][accounts[i]].free;
        }
    }

    /// @notice Returns true if `operator` is approved to manage `account`'s tokens
    function isApprovedForAll(address account, address operator) external view returns (bool) {
        return LibERC1155Storage.layout().operatorApprovals[account][operator];
    }

    /// @notice Returns all four partition balances for a holder
    function partitionBalanceOf(address account, uint256 id)
        external
        view
        returns (uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement)
    {
        PartitionBalance storage pb = LibERC1155Storage.layout().partitions[id][account];
        free = pb.free;
        locked = pb.locked;
        custody = pb.custody;
        pendingSettlement = pb.pendingSettlement;
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — VALIDATE AND TRANSFER
    //////////////////////////////////////////////////////////////*/

    function _validateAndTransferSingle(
        address operator,
        address from,
        address to,
        uint256 id,
        uint256 amount,
        bytes calldata data
    ) internal {
        if (to == address(0)) revert ERC1155Facet__TransferToZeroAddress();
        _enforceApproval(operator, from);

        // Global checks (steps 1-3)
        _validateGlobal(from, to);

        // Per-tokenId checks (steps 7-13) + execute
        bytes32 reason = _validatePerToken(from, to, id, amount, data);

        // Execute
        _executeTransfer(from, to, id, amount);

        emit TransferSingle(operator, from, to, id, amount);
        emit RegulatoryTransfer(id, from, to, amount, reason);

        // Post-hook
        _compliancePostTransfer(id, from, to, amount);

        // ERC-1155 receiver callback
        _checkOnERC1155Received(operator, from, to, id, amount, data);
    }

    function _validateAndTransferBatch(
        address operator,
        address from,
        address to,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        bytes calldata data
    ) internal {
        if (to == address(0)) revert ERC1155Facet__TransferToZeroAddress();
        _enforceApproval(operator, from);

        // Global checks — ONCE per call (steps 1-3)
        _validateGlobal(from, to);

        // Per-tokenId checks — in loop (steps 7-13)
        for (uint256 i; i < ids.length; ++i) {
            bytes32 reason = _validatePerToken(from, to, ids[i], amounts[i], data);
            _executeTransfer(from, to, ids[i], amounts[i]);
            emit RegulatoryTransfer(ids[i], from, to, amounts[i], reason);
            _compliancePostTransfer(ids[i], from, to, amounts[i]);
        }

        emit TransferBatch(operator, from, to, ids, amounts);

        // ERC-1155 receiver callback
        _checkOnERC1155BatchReceived(operator, from, to, ids, amounts, data);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — EXECUTE
    //////////////////////////////////////////////////////////////*/

    /// @dev Mutates partition balances and updates holder tracking
    function _executeTransfer(address from, address to, uint256 id, uint256 amount) internal {
        ERC1155Storage storage es = LibERC1155Storage.layout();
        es.partitions[id][from].free -= amount;
        es.partitions[id][to].free += amount;

        // Update holder tracking
        SupplyStorage storage ss = LibSupplyStorage.layout();
        if (!ss.isHolder[id][to]) {
            ss.isHolder[id][to] = true;
            ss.holderCount[id] += 1;
        }
        PartitionBalance storage pb = es.partitions[id][from];
        if (pb.free == 0 && pb.locked == 0 && pb.custody == 0 && pb.pendingSettlement == 0) {
            ss.isHolder[id][from] = false;
            ss.holderCount[id] -= 1;
        }
    }

    function _compliancePostTransfer(uint256 tokenId, address from, address to, uint256 amount) internal {
        address[] storage modules = LibAssetStorage.layout().configs[tokenId].complianceModules;
        uint256 len = modules.length;
        for (uint256 i; i < len;) {
            IComplianceModule(modules[i]).transferred(tokenId, from, to, amount);
            unchecked { ++i; }
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — VALIDATION STEPS
    //////////////////////////////////////////////////////////////*/

    /// @dev Steps 1-3: global pause and freeze checks
    function _validateGlobal(address from, address to) internal view {
        AppStorage storage app = LibAppStorage.layout();
        if (app.globalPaused) revert ERC1155Facet__ProtocolPaused();

        FreezeStorage storage fs = LibFreezeStorage.layout();
        if (fs.globalFreeze[from]) revert ERC1155Facet__WalletFrozenGlobal(from);
        if (fs.globalFreeze[to]) revert ERC1155Facet__WalletFrozenGlobal(to);
    }

    /// @dev Steps 7-13: per-tokenId checks
    function _validatePerToken(address from, address to, uint256 tokenId, uint256 amount, bytes calldata data)
        internal
        view
        returns (bytes32 reason)
    {
        AssetStorage storage as_ = LibAssetStorage.layout();
        AssetConfig storage config = as_.configs[tokenId];

        // Step 7: asset must exist and not be paused
        if (!config.exists) revert ERC1155Facet__AssetNotRegistered(tokenId);
        if (config.paused) revert ERC1155Facet__AssetPaused(tokenId);

        FreezeStorage storage fs = LibFreezeStorage.layout();

        // Step 8-9: asset-level freeze
        if (fs.assetFreeze[tokenId][from]) revert ERC1155Facet__WalletFrozenAsset(tokenId, from);
        if (fs.assetFreeze[tokenId][to]) revert ERC1155Facet__WalletFrozenAsset(tokenId, to);

        // Step 10: lockup expiry
        uint64 expiry = fs.lockupExpiry[tokenId][from];
        if (expiry != 0 && expiry >= block.timestamp) {
            revert ERC1155Facet__LockupActive(tokenId, from, expiry);
        }

        // Step 11: free balance minus frozen amount
        ERC1155Storage storage es = LibERC1155Storage.layout();
        uint256 freeBalance = es.partitions[tokenId][from].free;
        uint256 frozen = fs.frozenAmount[tokenId][from];
        uint256 available = freeBalance > frozen ? freeBalance - frozen : 0;
        if (available < amount) {
            revert ERC1155Facet__InsufficientFreeBalance(tokenId, from, available, amount);
        }

        // Step 13: compliance modules check (early exit on first rejection)
        address[] storage modules = config.complianceModules;
        uint256 mLen = modules.length;
        for (uint256 j; j < mLen;) {
            bool ok;
            (ok, reason) = IComplianceModule(modules[j]).canTransfer(tokenId, from, to, amount, data);
            if (!ok) revert ERC1155Facet__ComplianceRejected(tokenId, reason);
            unchecked { ++j; }
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — HELPERS
    //////////////////////////////////////////////////////////////*/

    function _enforceApproval(address operator, address from) internal view {
        if (operator != from) {
            if (!LibERC1155Storage.layout().operatorApprovals[from][operator]) {
                revert ERC1155Facet__NotApprovedOrOwner();
            }
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — RECEIVER CALLBACKS
    //////////////////////////////////////////////////////////////*/

    /// @dev Calls onERC1155Received on `to` if it is a contract.
    ///      Reverts if the receiver rejects the transfer or does not implement the interface.
    function _checkOnERC1155Received(
        address operator,
        address from,
        address to,
        uint256 id,
        uint256 amount,
        bytes calldata data
    ) internal {
        if (to.code.length > 0) {
            try IERC1155Receiver(to).onERC1155Received(operator, from, id, amount, data) returns (bytes4 response) {
                if (response != IERC1155Receiver.onERC1155Received.selector) {
                    revert ERC1155Facet__InvalidReceiver(to);
                }
            } catch {
                revert ERC1155Facet__InvalidReceiver(to);
            }
        }
    }

    /// @dev Calls onERC1155BatchReceived on `to` if it is a contract.
    function _checkOnERC1155BatchReceived(
        address operator,
        address from,
        address to,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        bytes calldata data
    ) internal {
        if (to.code.length > 0) {
            try IERC1155Receiver(to).onERC1155BatchReceived(operator, from, ids, amounts, data) returns (
                bytes4 response
            ) {
                if (response != IERC1155Receiver.onERC1155BatchReceived.selector) {
                    revert ERC1155Facet__InvalidReceiver(to);
                }
            } catch {
                revert ERC1155Facet__InvalidReceiver(to);
            }
        }
    }
}
