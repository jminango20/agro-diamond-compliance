// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {Diamond} from "../src/Diamond.sol";
import {DiamondInit} from "../src/initializers/DiamondInit.sol";
import {DiamondCutFacet} from "../src/facets/core/DiamondCutFacet.sol";
import {DiamondLoupeFacet} from "../src/facets/core/DiamondLoupeFacet.sol";
import {OwnershipFacet} from "../src/facets/core/OwnershipFacet.sol";
import {AccessControlFacet} from "../src/facets/security/AccessControlFacet.sol";
import {PauseFacet} from "../src/facets/security/PauseFacet.sol";
import {EmergencyFacet} from "../src/facets/security/EmergencyFacet.sol";
import {FreezeFacet} from "../src/facets/rwa/FreezeFacet.sol";
import {RecoveryFacet} from "../src/facets/rwa/RecoveryFacet.sol";
import {AssetManagerFacet} from "../src/facets/token/AssetManagerFacet.sol";
import {ClaimTopicsFacet} from "../src/facets/identity/ClaimTopicsFacet.sol";
import {TrustedIssuerFacet} from "../src/facets/identity/TrustedIssuerFacet.sol";
import {IdentityRegistryFacet} from "../src/facets/identity/IdentityRegistryFacet.sol";
import {ComplianceRouterFacet} from "../src/facets/compliance/ComplianceRouterFacet.sol";
import {ERC1155Facet} from "../src/facets/token/ERC1155Facet.sol";
import {SupplyFacet} from "../src/facets/token/SupplyFacet.sol";
import {MetadataFacet} from "../src/facets/token/MetadataFacet.sol";
import {IDiamond, IDiamondCut, IDiamondLoupe} from "../src/interfaces/core/IDiamond.sol";
// solhint-disable-next-line import-path-check
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";

/// @title Deploy
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Deploys the Diamond proxy with all facets and initializes ERC-165 interfaces.
///
/// Usage (local):
///   anvil &
///   forge script script/Deploy.s.sol --rpc-url http://localhost:8545 --broadcast -vvvv
///
/// Usage (production — never use plain --private-key):
///   cast wallet import deployer --interactive
///   forge script script/Deploy.s.sol --rpc-url $RPC_URL --account deployer --broadcast -vvvv
///
/// Environment:
///   OWNER_ADDRESS — Diamond owner (defaults to msg.sender)
contract Deploy is Script {
    function run() external {
        address owner = vm.envOr("OWNER_ADDRESS", msg.sender);

        vm.startBroadcast();

        // ── 1. Deploy facets ────────────────────────────────────────

        DiamondCutFacet cutFacet = new DiamondCutFacet();
        DiamondLoupeFacet loupeFacet = new DiamondLoupeFacet();
        OwnershipFacet ownershipFacet = new OwnershipFacet();
        AccessControlFacet accessControlFacet = new AccessControlFacet();
        PauseFacet pauseFacet = new PauseFacet();
        EmergencyFacet emergencyFacet = new EmergencyFacet();
        FreezeFacet freezeFacet = new FreezeFacet();
        RecoveryFacet recoveryFacet = new RecoveryFacet();
        AssetManagerFacet assetManagerFacet = new AssetManagerFacet();
        ClaimTopicsFacet claimTopicsFacet = new ClaimTopicsFacet();
        TrustedIssuerFacet trustedIssuerFacet = new TrustedIssuerFacet();
        IdentityRegistryFacet identityRegistryFacet = new IdentityRegistryFacet();
        ComplianceRouterFacet complianceRouterFacet = new ComplianceRouterFacet();
        ERC1155Facet erc1155Facet = new ERC1155Facet();
        SupplyFacet supplyFacet = new SupplyFacet();
        MetadataFacet metadataFacet = new MetadataFacet();
        DiamondInit diamondInit = new DiamondInit();

        // ── 2. Deploy Diamond ───────────────────────────────────────

        Diamond diamond = new Diamond(owner, address(cutFacet));

        // ── 3. Build facet cuts ─────────────────────────────────────

        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](15);

        cuts[0] = _cut(address(loupeFacet), _loupeSelectors());
        cuts[1] = _cut(address(ownershipFacet), _ownershipSelectors());
        cuts[2] = _cut(address(accessControlFacet), _accessControlSelectors());
        cuts[3] = _cut(address(pauseFacet), _pauseSelectors());
        cuts[4] = _cut(address(emergencyFacet), _emergencySelectors());
        cuts[5] = _cut(address(freezeFacet), _freezeSelectors());
        cuts[6] = _cut(address(recoveryFacet), _recoverySelectors());
        cuts[7] = _cut(address(assetManagerFacet), _assetManagerSelectors());
        cuts[8] = _cut(address(claimTopicsFacet), _claimTopicsSelectors());
        cuts[9] = _cut(address(trustedIssuerFacet), _trustedIssuerSelectors());
        cuts[10] = _cut(address(identityRegistryFacet), _identityRegistrySelectors());
        cuts[11] = _cut(address(complianceRouterFacet), _complianceRouterSelectors());
        cuts[12] = _cut(address(erc1155Facet), _erc1155Selectors());
        cuts[13] = _cut(address(supplyFacet), _supplySelectors());
        cuts[14] = _cut(address(metadataFacet), _metadataSelectors());

        // ── 4. Execute diamond cut + init ───────────────────────────

        IDiamondCut(address(diamond)).diamondCut(
            cuts, address(diamondInit), abi.encodeCall(DiamondInit.init, ())
        );

        vm.stopBroadcast();

        // ── 5. Log addresses ────────────────────────────────────────

        console2.log("");
        console2.log("=== Diamond ERC-3643 Protocol ===");
        console2.log("");
        console2.log("Diamond              :", address(diamond));
        console2.log("Owner                :", owner);
        console2.log("");
        console2.log("--- Core ---");
        console2.log("DiamondCutFacet      :", address(cutFacet));
        console2.log("DiamondLoupeFacet    :", address(loupeFacet));
        console2.log("OwnershipFacet       :", address(ownershipFacet));
        console2.log("");
        console2.log("--- Security ---");
        console2.log("AccessControlFacet   :", address(accessControlFacet));
        console2.log("PauseFacet           :", address(pauseFacet));
        console2.log("EmergencyFacet       :", address(emergencyFacet));
        console2.log("");
        console2.log("--- RWA ---");
        console2.log("FreezeFacet          :", address(freezeFacet));
        console2.log("RecoveryFacet        :", address(recoveryFacet));
        console2.log("");
        console2.log("--- Token ---");
        console2.log("AssetManagerFacet    :", address(assetManagerFacet));
        console2.log("ERC1155Facet         :", address(erc1155Facet));
        console2.log("SupplyFacet          :", address(supplyFacet));
        console2.log("MetadataFacet        :", address(metadataFacet));
        console2.log("");
        console2.log("--- Identity ---");
        console2.log("ClaimTopicsFacet     :", address(claimTopicsFacet));
        console2.log("TrustedIssuerFacet   :", address(trustedIssuerFacet));
        console2.log("IdentityRegistryFacet:", address(identityRegistryFacet));
        console2.log("");
        console2.log("--- Compliance ---");
        console2.log("ComplianceRouterFacet:", address(complianceRouterFacet));
        console2.log("");
        console2.log("DiamondInit          :", address(diamondInit));
        console2.log("");

        // ── 6. Verify ───────────────────────────────────────────────

        _verify(diamond, owner);
    }

    /*//////////////////////////////////////////////////////////////
                            HELPERS
    //////////////////////////////////////////////////////////////*/

    function _cut(address facet, bytes4[] memory selectors)
        internal
        pure
        returns (IDiamond.FacetCut memory)
    {
        return IDiamond.FacetCut({
            facetAddress: facet,
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: selectors
        });
    }

    /*//////////////////////////////////////////////////////////////
                        SELECTOR BUILDERS
    //////////////////////////////////////////////////////////////*/

    function _loupeSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](5);
        sels[0] = IDiamondLoupe.facets.selector;
        sels[1] = IDiamondLoupe.facetFunctionSelectors.selector;
        sels[2] = IDiamondLoupe.facetAddresses.selector;
        sels[3] = IDiamondLoupe.facetAddress.selector;
        sels[4] = IERC165.supportsInterface.selector;
    }

    function _ownershipSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](4);
        sels[0] = OwnershipFacet.transferOwnership.selector;
        sels[1] = OwnershipFacet.acceptOwnership.selector;
        sels[2] = OwnershipFacet.owner.selector;
        sels[3] = OwnershipFacet.pendingOwner.selector;
    }

    function _accessControlSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](6);
        sels[0] = AccessControlFacet.grantRole.selector;
        sels[1] = AccessControlFacet.revokeRole.selector;
        sels[2] = AccessControlFacet.renounceRole.selector;
        sels[3] = AccessControlFacet.setRoleAdmin.selector;
        sels[4] = AccessControlFacet.hasRole.selector;
        sels[5] = AccessControlFacet.getRoleAdmin.selector;
    }

    function _pauseSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](6);
        sels[0] = PauseFacet.pauseProtocol.selector;
        sels[1] = PauseFacet.unpauseProtocol.selector;
        sels[2] = PauseFacet.pauseAsset.selector;
        sels[3] = PauseFacet.unpauseAsset.selector;
        sels[4] = PauseFacet.isProtocolPaused.selector;
        sels[5] = PauseFacet.isAssetPaused.selector;
    }

    function _emergencySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](2);
        sels[0] = EmergencyFacet.emergencyPause.selector;
        sels[1] = EmergencyFacet.isEmergencyPaused.selector;
    }

    function _freezeSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](8);
        sels[0] = FreezeFacet.setWalletFrozen.selector;
        sels[1] = FreezeFacet.setAssetWalletFrozen.selector;
        sels[2] = FreezeFacet.setFrozenAmount.selector;
        sels[3] = FreezeFacet.setLockupExpiry.selector;
        sels[4] = FreezeFacet.isWalletFrozen.selector;
        sels[5] = FreezeFacet.isAssetWalletFrozen.selector;
        sels[6] = FreezeFacet.getFrozenAmount.selector;
        sels[7] = FreezeFacet.getLockupExpiry.selector;
    }

    function _recoverySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](1);
        sels[0] = RecoveryFacet.recoverWallet.selector;
    }

    function _assetManagerSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](10);
        sels[0] = AssetManagerFacet.registerAsset.selector;
        sels[1] = AssetManagerFacet.setComplianceModule.selector;
        sels[2] = AssetManagerFacet.setIdentityProfile.selector;
        sels[3] = AssetManagerFacet.setIssuer.selector;
        sels[4] = AssetManagerFacet.setSupplyCap.selector;
        sels[5] = AssetManagerFacet.setAllowedCountries.selector;
        sels[6] = AssetManagerFacet.setAssetUri.selector;
        sels[7] = AssetManagerFacet.getAssetConfig.selector;
        sels[8] = AssetManagerFacet.getRegisteredTokenIds.selector;
        sels[9] = AssetManagerFacet.assetExists.selector;
    }

    function _claimTopicsSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](5);
        sels[0] = ClaimTopicsFacet.createProfile.selector;
        sels[1] = ClaimTopicsFacet.setProfileClaimTopics.selector;
        sels[2] = ClaimTopicsFacet.getProfileClaimTopics.selector;
        sels[3] = ClaimTopicsFacet.getProfileVersion.selector;
        sels[4] = ClaimTopicsFacet.profileExists.selector;
    }

    function _trustedIssuerSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](3);
        sels[0] = TrustedIssuerFacet.addTrustedIssuer.selector;
        sels[1] = TrustedIssuerFacet.removeTrustedIssuer.selector;
        sels[2] = TrustedIssuerFacet.isTrustedIssuer.selector;
    }

    function _identityRegistrySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](9);
        sels[0] = IdentityRegistryFacet.registerIdentity.selector;
        sels[1] = IdentityRegistryFacet.deleteIdentity.selector;
        sels[2] = IdentityRegistryFacet.updateIdentity.selector;
        sels[3] = IdentityRegistryFacet.updateCountry.selector;
        sels[4] = IdentityRegistryFacet.batchRegisterIdentity.selector;
        sels[5] = IdentityRegistryFacet.isVerified.selector;
        sels[6] = IdentityRegistryFacet.getIdentity.selector;
        sels[7] = IdentityRegistryFacet.getCountry.selector;
        sels[8] = IdentityRegistryFacet.contains.selector;
    }

    function _complianceRouterSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](5);
        sels[0] = ComplianceRouterFacet.canTransfer.selector;
        sels[1] = ComplianceRouterFacet.transferred.selector;
        sels[2] = ComplianceRouterFacet.minted.selector;
        sels[3] = ComplianceRouterFacet.burned.selector;
        sels[4] = ComplianceRouterFacet.getComplianceModule.selector;
    }

    function _erc1155Selectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](7);
        sels[0] = ERC1155Facet.safeTransferFrom.selector;
        sels[1] = ERC1155Facet.safeBatchTransferFrom.selector;
        sels[2] = ERC1155Facet.setApprovalForAll.selector;
        sels[3] = ERC1155Facet.balanceOf.selector;
        sels[4] = ERC1155Facet.balanceOfBatch.selector;
        sels[5] = ERC1155Facet.isApprovedForAll.selector;
        sels[6] = ERC1155Facet.partitionBalanceOf.selector;
    }

    function _supplySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](7);
        sels[0] = SupplyFacet.mint.selector;
        sels[1] = SupplyFacet.batchMint.selector;
        sels[2] = SupplyFacet.burn.selector;
        sels[3] = SupplyFacet.forcedTransfer.selector;
        sels[4] = SupplyFacet.totalSupply.selector;
        sels[5] = SupplyFacet.holderCount.selector;
        sels[6] = SupplyFacet.isHolder.selector;
    }

    function _metadataSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](7);
        sels[0] = MetadataFacet.uri.selector;
        sels[1] = MetadataFacet.name.selector;
        sels[2] = MetadataFacet.symbol.selector;
        sels[3] = MetadataFacet.supplyCap.selector;
        sels[4] = MetadataFacet.issuer.selector;
        sels[5] = MetadataFacet.allowedCountries.selector;
        sels[6] = MetadataFacet.tokenInfo.selector;
    }

    /*//////////////////////////////////////////////////////////////
                        POST-DEPLOY VERIFICATION
    //////////////////////////////////////////////////////////////*/

    function _verify(Diamond diamond, address expectedOwner) internal view {
        address[] memory facetAddrs = IDiamondLoupe(address(diamond)).facetAddresses();
        require(facetAddrs.length == 16, "Deploy: expected 16 facets");

        require(
            OwnershipFacet(address(diamond)).owner() == expectedOwner,
            "Deploy: wrong owner"
        );

        require(
            IERC165(address(diamond)).supportsInterface(type(IERC165).interfaceId),
            "Deploy: missing IERC165"
        );
        require(
            IERC165(address(diamond)).supportsInterface(type(IDiamondCut).interfaceId),
            "Deploy: missing IDiamondCut"
        );
        require(
            IERC165(address(diamond)).supportsInterface(type(IDiamondLoupe).interfaceId),
            "Deploy: missing IDiamondLoupe"
        );

        console2.log("=== Verification OK ===");
        console2.log("Facets registered    :", facetAddrs.length);
    }
}
