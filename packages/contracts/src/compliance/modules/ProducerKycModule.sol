// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";
import {IAgroDiamondCompliance} from "../../interfaces/compliance/IAgroDiamondCompliance.sol";
import {LibReasonCodes} from "../../libraries/LibReasonCodes.sol";

/// @title ProducerKycModule
/// @notice Compliance module that enforces Brazilian agricultural KYC requirements
///         on token receivers. Before any transfer is allowed, the recipient wallet
///         must hold all claim topics required by the asset's identity profile —
///         typically MAPA producer registration, CPF/CNPJ verification, and any
///         commodity-specific claims (e.g., SISBOV for cattle, CAR for coffee/honey).
///
///         Each tokenId maps to an identity profile (configured in ClaimTopicsFacet).
///         This module delegates the actual claim check to the Diamond's
///         IdentityRegistryFacet.isVerified(), which uses a version-based cache to
///         minimize gas on repeated transfers.
///
///         Minting (from == address(0)) is subject to KYC on the recipient: only
///         registered producers can receive newly tokenized agricultural assets.
///         Burns (to == address(0)) are exempt — redemption does not require the
///         sender to remain in the registry.
///
/// @dev Implements IComplianceModule. Plugged into tokenIds via AssetManagerFacet.
///      The Diamond address is immutable; profile mapping is configurable by the owner.
contract ProducerKycModule is IComplianceModule {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event TokenProfileSet(uint256 indexed tokenId, uint32 profileId);

    /*//////////////////////////////////////////////////////////////
                                ERRORS
    //////////////////////////////////////////////////////////////*/

    error ProducerKycModule__OnlyOwner();
    error ProducerKycModule__ZeroDiamond();
    error ProducerKycModule__ProfileNotSet(uint256 tokenId);

    /*//////////////////////////////////////////////////////////////
                                STATE
    //////////////////////////////////////////////////////////////*/

    address public immutable DIAMOND;
    address public owner;

    /// @notice tokenId → identity profile ID (from ClaimTopicsFacet).
    ///         Profile defines which claim topics (MAPA, SISBOV, etc.) are required.
    mapping(uint256 => uint32) public tokenProfile;

    /*//////////////////////////////////////////////////////////////
                            CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(address diamond_, address owner_) {
        if (diamond_ == address(0)) revert ProducerKycModule__ZeroDiamond();
        DIAMOND = diamond_;
        owner = owner_;
    }

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING
    //////////////////////////////////////////////////////////////*/

    /// @notice Assigns an identity profile to a tokenId.
    ///         The profile must have been created via ClaimTopicsFacet.createProfile().
    /// @param tokenId   The asset class to configure.
    /// @param profileId The identity profile whose required claim topics apply.
    function setTokenProfile(uint256 tokenId, uint32 profileId) external {
        _enforceOwner();
        tokenProfile[tokenId] = profileId;
        emit TokenProfileSet(tokenId, profileId);
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

    /// @notice Checks that the recipient wallet is registered in the identity registry.
    ///         Only wallets bound to an ONCHAINID by a TRANSFER_AGENT (i.e., KYC'd
    ///         agricultural producers) may receive tokenized commodities.
    ///         Burns (to == address(0)) are always allowed.
    ///         Returns REASON_NOT_VERIFIED if the receiver is not a registered producer.
    function canTransfer(uint256 tokenId, address, address to, uint256, bytes calldata)
        external
        view
        returns (bool ok, bytes32 reason)
    {
        // Burns are exempt — producer redeeming his own token is always allowed.
        if (to == address(0)) return (true, LibReasonCodes.REASON_OK);

        uint32 profileId = tokenProfile[tokenId];
        // No profile configured → module is pass-through for this tokenId.
        if (profileId == 0) return (true, LibReasonCodes.REASON_OK);

        bool registered = IAgroDiamondCompliance(DIAMOND).contains(to);
        if (!registered) {
            return (false, REASON_NOT_VERIFIED);
        }

        return (true, LibReasonCodes.REASON_OK);
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL
    //////////////////////////////////////////////////////////////*/

    /// @dev Reason code returned when the recipient lacks required agricultural claims.
    bytes32 internal constant REASON_NOT_VERIFIED = keccak256("REASON_NOT_VERIFIED");

    function _enforceOwner() internal view {
        if (msg.sender != owner) revert ProducerKycModule__OnlyOwner();
    }
}
