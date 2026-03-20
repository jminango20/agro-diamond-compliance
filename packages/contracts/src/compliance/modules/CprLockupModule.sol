// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {LibReasonCodes} from "../../libraries/LibReasonCodes.sol";

/// @title CprLockupModule
/// @notice Compliance module that enforces lock-up periods derived from
///         Cédula de Produto Rural (CPR) instruments registered with Brazilian
///         financial institutions.
///
///         A CPR is a rural credit note issued against a future harvest. Until the
///         CPR matures (harvest delivery date), the underlying tokenized commodity
///         cannot be freely transferred — it serves as collateral for the credit
///         operation. This module blocks all transfers for a tokenId until the
///         configured expiry timestamp passes.
///
///         Lock-up is per-tokenId: different coffee lots (same asset class, different
///         IDs) can have independent CPR maturity dates. A tokenId with expiry = 0
///         is unlocked (no CPR restriction).
///
/// @dev Implements IComplianceModule. Plugged into each tokenId via AssetManagerFacet.
///      Only the module owner (typically the Diamond owner or a compliance admin)
///      can set or clear lock-up periods.
contract CprLockupModule is IComplianceModule {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event LockupSet(uint256 indexed tokenId, uint256 expiry);
    event LockupCleared(uint256 indexed tokenId);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error CprLockupModule__OnlyOwner();
    error CprLockupModule__ExpiryInPast();

    /*//////////////////////////////////////////////////////////////
                                STATE
    //////////////////////////////////////////////////////////////*/

    address public owner;

    /// @notice tokenId → CPR maturity timestamp (unix seconds).
    ///         0 means no lock-up is active for that tokenId.
    mapping(uint256 => uint256) public lockupExpiry;

    /*//////////////////////////////////////////////////////////////
                            CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address owner_) {
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Sets a CPR lock-up expiry for a tokenId.
    ///         All transfers of that tokenId are blocked until block.timestamp > expiry.
    /// @param tokenId   The asset class to lock (e.g., a specific coffee lot).
    /// @param expiry    Unix timestamp of the CPR maturity date. Must be in the future.
    function setLockup(uint256 tokenId, uint256 expiry) external {
        _enforceOwner();
        if (expiry <= block.timestamp) revert CprLockupModule__ExpiryInPast();
        lockupExpiry[tokenId] = expiry;
        emit LockupSet(tokenId, expiry);
    }

    /// @notice Removes the CPR lock-up for a tokenId (e.g., after CPR settlement).
    /// @param tokenId The asset class to unlock.
    function clearLockup(uint256 tokenId) external {
        _enforceOwner();
        delete lockupExpiry[tokenId];
        emit LockupCleared(tokenId);
    }

    /// @notice Returns true if the tokenId is currently under a CPR lock-up.
    function isLocked(uint256 tokenId) external view returns (bool) {
        uint256 expiry = lockupExpiry[tokenId];
        return expiry != 0 && block.timestamp <= expiry;
    }

    /*//////////////////////////////////////////////////////////////
                    IComplianceModule — HOOKS
    //////////////////////////////////////////////////////////////*/

    function transferred(uint256, address, address, uint256) external {}

    function minted(uint256, address, uint256) external {}

    function burned(uint256, address, uint256) external {}

    /*//////////////////////////////////////////////////////////////
                    IComplianceModule — VALIDATION
    //////////////////////////////////////////////////////////////*/

    /// @notice Blocks transfers when the CPR lock-up period is active.
    ///         Minting by the issuer is always allowed (from == address(0)).
    function canTransfer(uint256 tokenId, address from, address, uint256, bytes calldata)
        external
        view
        returns (bool ok, bytes32 reason)
    {
        // Minting (from == address(0)) is exempt from CPR lock-up:
        // the issuer creates the token at CPR origination.
        if (from == address(0)) return (true, LibReasonCodes.REASON_OK);

        uint256 expiry = lockupExpiry[tokenId];
        if (expiry != 0 && block.timestamp <= expiry) {
            return (false, REASON_CPR_LOCKED);
        }

        return (true, LibReasonCodes.REASON_OK);
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL
    //////////////////////////////////////////////////////////////*/

    /// @dev Reason code returned when a transfer is blocked by CPR lock-up.
    bytes32 internal constant REASON_CPR_LOCKED = keccak256("REASON_CPR_LOCKED");

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert CprLockupModule__OnlyOwner();
    }
}
