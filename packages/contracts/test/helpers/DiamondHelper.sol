// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {Diamond} from "../../src/Diamond.sol";
import {DiamondCutFacet} from "../../src/facets/core/DiamondCutFacet.sol";
import {DiamondLoupeFacet} from "../../src/facets/core/DiamondLoupeFacet.sol";
import {OwnershipFacet} from "../../src/facets/core/OwnershipFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {PauseFacet} from "../../src/facets/security/PauseFacet.sol";
import {EmergencyFacet} from "../../src/facets/security/EmergencyFacet.sol";
import {FreezeFacet} from "../../src/facets/rwa/FreezeFacet.sol";
import {RecoveryFacet} from "../../src/facets/rwa/RecoveryFacet.sol";
import {SnapshotFacet} from "../../src/facets/rwa/SnapshotFacet.sol";
import {DividendFacet} from "../../src/facets/rwa/DividendFacet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {ClaimTopicsFacet} from "../../src/facets/identity/ClaimTopicsFacet.sol";
import {TrustedIssuerFacet} from "../../src/facets/identity/TrustedIssuerFacet.sol";
import {IdentityRegistryFacet} from "../../src/facets/identity/IdentityRegistryFacet.sol";
import {ComplianceRouterFacet} from "../../src/facets/compliance/ComplianceRouterFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {MetadataFacet} from "../../src/facets/token/MetadataFacet.sol";
import {IDiamond, IDiamondCut, IDiamondLoupe} from "../../src/interfaces/core/IDiamond.sol";
import {DiamondInit} from "../../src/initializers/DiamondInit.sol";

contract DiamondHelper is Test {
    struct CoreFacets {
        DiamondCutFacet cutFacet;
        DiamondLoupeFacet loupeFacet;
        OwnershipFacet ownershipFacet;
        AccessControlFacet accessControlFacet;
        PauseFacet pauseFacet;
        EmergencyFacet emergencyFacet;
        FreezeFacet freezeFacet;
        RecoveryFacet recoveryFacet;
    }

    struct TokenFacets {
        AssetManagerFacet assetManagerFacet;
        ERC1155Facet erc1155Facet;
        SupplyFacet supplyFacet;
        MetadataFacet metadataFacet;
        SnapshotFacet snapshotFacet;
        DividendFacet dividendFacet;
    }

    struct ComplianceFacets {
        ClaimTopicsFacet claimTopicsFacet;
        TrustedIssuerFacet trustedIssuerFacet;
        IdentityRegistryFacet identityRegistryFacet;
        ComplianceRouterFacet complianceRouterFacet;
    }

    struct DeployedDiamond {
        Diamond diamond;
        CoreFacets core;
        TokenFacets token;
        ComplianceFacets compliance;
    }

    function deployDiamond(address owner) internal returns (DeployedDiamond memory d) {
        d.core = _deployCoreFacets();
        d.token = _deployTokenFacets();
        d.compliance = _deployComplianceFacets();
        DiamondInit diamondInit = new DiamondInit();

        d.diamond = new Diamond(owner, address(d.core.cutFacet));

        IDiamond.FacetCut[] memory cuts = new IDiamond.FacetCut[](17);
        _fillCoreCuts(cuts, d.core);
        _fillTokenCuts(cuts, d.token);
        _fillComplianceCuts(cuts, d.compliance);

        vm.prank(owner);
        IDiamondCut(address(d.diamond)).diamondCut(
            cuts, address(diamondInit), abi.encodeCall(DiamondInit.init, ())
        );
    }

    function _deployCoreFacets() internal returns (CoreFacets memory c) {
        c.cutFacet = new DiamondCutFacet();
        c.loupeFacet = new DiamondLoupeFacet();
        c.ownershipFacet = new OwnershipFacet();
        c.accessControlFacet = new AccessControlFacet();
        c.pauseFacet = new PauseFacet();
        c.emergencyFacet = new EmergencyFacet();
        c.freezeFacet = new FreezeFacet();
        c.recoveryFacet = new RecoveryFacet();
    }

    function _deployTokenFacets() internal returns (TokenFacets memory t) {
        t.assetManagerFacet = new AssetManagerFacet();
        t.erc1155Facet = new ERC1155Facet();
        t.supplyFacet = new SupplyFacet();
        t.metadataFacet = new MetadataFacet();
        t.snapshotFacet = new SnapshotFacet();
        t.dividendFacet = new DividendFacet();
    }

    function _deployComplianceFacets() internal returns (ComplianceFacets memory c) {
        c.claimTopicsFacet = new ClaimTopicsFacet();
        c.trustedIssuerFacet = new TrustedIssuerFacet();
        c.identityRegistryFacet = new IdentityRegistryFacet();
        c.complianceRouterFacet = new ComplianceRouterFacet();
    }

    function _fillCoreCuts(IDiamond.FacetCut[] memory cuts, CoreFacets memory c) internal pure {
        cuts[0] = _cut(address(c.loupeFacet), _loupeSelectors());
        cuts[1] = _cut(address(c.ownershipFacet), _ownershipSelectors());
        cuts[2] = _cut(address(c.accessControlFacet), _accessControlSelectors());
        cuts[3] = _cut(address(c.pauseFacet), _pauseSelectors());
        cuts[4] = _cut(address(c.emergencyFacet), _emergencySelectors());
        cuts[5] = _cut(address(c.freezeFacet), _freezeSelectors());
        cuts[14] = _cut(address(c.recoveryFacet), _recoverySelectors());
    }

    function _fillTokenCuts(IDiamond.FacetCut[] memory cuts, TokenFacets memory t) internal pure {
        cuts[6] = _cut(address(t.assetManagerFacet), _assetManagerSelectors());
        cuts[11] = _cut(address(t.erc1155Facet), _erc1155Selectors());
        cuts[12] = _cut(address(t.supplyFacet), _supplySelectors());
        cuts[13] = _cut(address(t.metadataFacet), _metadataSelectors());
        cuts[15] = _cut(address(t.snapshotFacet), _snapshotSelectors());
        cuts[16] = _cut(address(t.dividendFacet), _dividendSelectors());
    }

    function _fillComplianceCuts(IDiamond.FacetCut[] memory cuts, ComplianceFacets memory c) internal pure {
        cuts[7] = _cut(address(c.claimTopicsFacet), _claimTopicsSelectors());
        cuts[8] = _cut(address(c.trustedIssuerFacet), _trustedIssuerSelectors());
        cuts[9] = _cut(address(c.identityRegistryFacet), _identityRegistrySelectors());
        cuts[10] = _cut(address(c.complianceRouterFacet), _complianceRouterSelectors());
    }

    function _cut(address facet, bytes4[] memory sels)
        internal
        pure
        returns (IDiamond.FacetCut memory)
    {
        return IDiamond.FacetCut({
            facetAddress: facet,
            action: IDiamond.FacetCutAction.Add,
            functionSelectors: sels
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
        sels[4] = bytes4(keccak256("supportsInterface(bytes4)"));
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

    function _recoverySelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](1);
        sels[0] = RecoveryFacet.recoverWallet.selector;
    }

    function _snapshotSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](8);
        sels[0] = SnapshotFacet.createSnapshot.selector;
        sels[1] = SnapshotFacet.recordHolder.selector;
        sels[2] = SnapshotFacet.recordHoldersBatch.selector;
        sels[3] = SnapshotFacet.getSnapshot.selector;
        sels[4] = SnapshotFacet.getSnapshotBalance.selector;
        sels[5] = SnapshotFacet.getTokenSnapshots.selector;
        sels[6] = SnapshotFacet.getLatestSnapshotId.selector;
        sels[7] = SnapshotFacet.nextSnapshotId.selector;
    }

    function _dividendSelectors() internal pure returns (bytes4[] memory sels) {
        sels = new bytes4[](6);
        sels[0] = DividendFacet.createDividend.selector;
        sels[1] = DividendFacet.claimDividend.selector;
        sels[2] = DividendFacet.getDividend.selector;
        sels[3] = DividendFacet.hasClaimed.selector;
        sels[4] = DividendFacet.claimableAmount.selector;
        sels[5] = DividendFacet.getTokenDividends.selector;
    }
}
