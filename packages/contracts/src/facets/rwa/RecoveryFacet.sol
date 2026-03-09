// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibERC1155Storage, ERC1155Storage, PartitionBalance} from "../../storage/LibERC1155Storage.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";
import {LibIdentityStorage, IdentityStorage} from "../../storage/LibIdentityStorage.sol";
import {LibFreezeStorage, FreezeStorage} from "../../storage/LibFreezeStorage.sol";
import {LibAccessStorage, AccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibDiamond} from "../../libraries/LibDiamond.sol";

/**
 * @title RecoveryFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Wallet recovery for regulated security tokens.
 *         When an investor loses access to their wallet, a RECOVERY_AGENT
 *         migrates all token balances, identity, and freeze state to a new
 *         wallet in a single atomic transaction.
 * @dev Recovery is the most privileged operation in the protocol — it moves
 *      assets without holder consent. Must be protected by multisig + timelock
 *      in production. Emits granular events for full auditability.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract RecoveryFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event WalletRecovered(address indexed lostWallet, address indexed newWallet, address indexed agent);
    event TokensRecovered(
        uint256 indexed tokenId,
        address indexed lostWallet,
        address indexed newWallet,
        uint256 free,
        uint256 locked,
        uint256 custody,
        uint256 pendingSettlement
    );

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error RecoveryFacet__Unauthorized();
    error RecoveryFacet__ZeroAddress();
    error RecoveryFacet__SameAddress();
    error RecoveryFacet__NewWalletAlreadyRegistered();

    /*//////////////////////////////////////////////////////////////
                            ROLE CONSTANT
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant RECOVERY_AGENT = keccak256("RECOVERY_AGENT");

    /*//////////////////////////////////////////////////////////////
                        EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Recovers all assets and identity from a lost wallet to a new wallet
    /// @param lostWallet The wallet that lost access
    /// @param newWallet The new wallet to receive everything
    function recoverWallet(address lostWallet, address newWallet) external {
        _enforceRecoveryAgentOrOwner();
        if (lostWallet == address(0) || newWallet == address(0)) revert RecoveryFacet__ZeroAddress();
        if (lostWallet == newWallet) revert RecoveryFacet__SameAddress();

        IdentityStorage storage ids = LibIdentityStorage.layout();
        if (ids.walletToIdentity[newWallet] != address(0)) {
            revert RecoveryFacet__NewWalletAlreadyRegistered();
        }

        _migrateIdentity(ids, lostWallet, newWallet);
        _migrateFreezeState(lostWallet, newWallet);
        _migrateAllTokenBalances(lostWallet, newWallet);

        emit WalletRecovered(lostWallet, newWallet, msg.sender);
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — IDENTITY MIGRATION
    //////////////////////////////////////////////////////////////*/

    /// @dev Moves ONCHAINID mapping and country code from lost to new wallet
    function _migrateIdentity(IdentityStorage storage ids, address lostWallet, address newWallet) internal {
        address onchainId = ids.walletToIdentity[lostWallet];
        uint16 country = ids.walletCountry[lostWallet];

        // Move identity to new wallet
        ids.walletToIdentity[newWallet] = onchainId;
        ids.walletCountry[newWallet] = country;
        ids.identityVersion[newWallet] = ids.identityVersion[lostWallet] + 1;

        // Clear old wallet
        delete ids.walletToIdentity[lostWallet];
        delete ids.walletCountry[lostWallet];
        delete ids.identityVersion[lostWallet];
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — FREEZE MIGRATION
    //////////////////////////////////////////////////////////////*/

    /// @dev Moves global and per-asset freeze state from lost to new wallet
    function _migrateFreezeState(address lostWallet, address newWallet) internal {
        FreezeStorage storage fs = LibFreezeStorage.layout();

        // Global freeze
        if (fs.globalFreeze[lostWallet]) {
            fs.globalFreeze[newWallet] = true;
            delete fs.globalFreeze[lostWallet];
        }

        // Per-asset freeze, frozen amounts, and lockups
        AssetStorage storage as_ = LibAssetStorage.layout();
        uint256[] storage tokenIds = as_.registeredTokenIds;
        for (uint256 i; i < tokenIds.length; ++i) {
            uint256 tokenId = tokenIds[i];

            if (fs.assetFreeze[tokenId][lostWallet]) {
                fs.assetFreeze[tokenId][newWallet] = true;
                delete fs.assetFreeze[tokenId][lostWallet];
            }

            uint256 frozenAmt = fs.frozenAmount[tokenId][lostWallet];
            if (frozenAmt != 0) {
                fs.frozenAmount[tokenId][newWallet] = frozenAmt;
                delete fs.frozenAmount[tokenId][lostWallet];
            }

            uint64 lockup = fs.lockupExpiry[tokenId][lostWallet];
            if (lockup != 0) {
                fs.lockupExpiry[tokenId][newWallet] = lockup;
                delete fs.lockupExpiry[tokenId][lostWallet];
            }
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — TOKEN BALANCE MIGRATION
    //////////////////////////////////////////////////////////////*/

    /// @dev Moves all partition balances across all registered tokenIds
    function _migrateAllTokenBalances(address lostWallet, address newWallet) internal {
        ERC1155Storage storage es = LibERC1155Storage.layout();
        SupplyStorage storage ss = LibSupplyStorage.layout();
        AssetStorage storage as_ = LibAssetStorage.layout();
        uint256[] storage tokenIds = as_.registeredTokenIds;

        for (uint256 i; i < tokenIds.length; ++i) {
            uint256 tokenId = tokenIds[i];
            PartitionBalance storage oldBal = es.partitions[tokenId][lostWallet];

            uint256 free = oldBal.free;
            uint256 locked = oldBal.locked;
            uint256 custody = oldBal.custody;
            uint256 pending = oldBal.pendingSettlement;

            // Skip tokenIds where lost wallet has no balance
            if (free == 0 && locked == 0 && custody == 0 && pending == 0) continue;

            // Move balances to new wallet
            PartitionBalance storage newBal = es.partitions[tokenId][newWallet];
            newBal.free += free;
            newBal.locked += locked;
            newBal.custody += custody;
            newBal.pendingSettlement += pending;

            // Clear old wallet
            delete es.partitions[tokenId][lostWallet];

            // Update holder tracking
            if (ss.isHolder[tokenId][lostWallet]) {
                ss.isHolder[tokenId][lostWallet] = false;
                if (!ss.isHolder[tokenId][newWallet]) {
                    ss.isHolder[tokenId][newWallet] = true;
                } else {
                    // newWallet was already a holder — net holder count decreases by 1
                    ss.holderCount[tokenId] -= 1;
                }
            }

            emit TokensRecovered(tokenId, lostWallet, newWallet, free, locked, custody, pending);
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INTERNAL — ACCESS CONTROL
    //////////////////////////////////////////////////////////////*/

    function _enforceRecoveryAgentOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isAgent = LibAccessStorage.layout().roles[RECOVERY_AGENT][msg.sender];
        if (!isOwner && !isAgent) revert RecoveryFacet__Unauthorized();
    }
}
