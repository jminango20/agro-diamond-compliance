// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Diamond} from "../../src/Diamond.sol";
import {DiamondCutFacet} from "../../src/facets/core/DiamondCutFacet.sol";
import {DiamondLoupeFacet} from "../../src/facets/core/DiamondLoupeFacet.sol";
import {OwnershipFacet} from "../../src/facets/core/OwnershipFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {MetadataFacet} from "../../src/facets/token/MetadataFacet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {DiamondInit} from "../../src/initializers/DiamondInit.sol";
import {IDiamond, IDiamondCut} from "../../src/interfaces/core/IDiamond.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

/// @title EchidnaSupply
/// @notice Echidna property tests for Diamond ERC-3643 supply invariants.
contract EchidnaSupply {
    address internal diamond;
    uint256 internal tokenId;
    uint256 internal constant SUPPLY_CAP = 1_000_000;

    address internal constant ALICE = address(0x10000);
    address internal constant BOB = address(0x20000);
    address internal constant CAROL = address(0x30000);

    uint256 internal ghost_minted;
    uint256 internal ghost_burned;

    constructor() {
        DiamondCutFacet cutFacet = new DiamondCutFacet();
        Diamond d = new Diamond(address(this), address(cutFacet));
        diamond = address(d);

        IDiamond.FacetCut[] memory cuts = _buildCuts();
        DiamondInit diamondInit = new DiamondInit();

        IDiamondCut(diamond).diamondCut(
            cuts,
            address(diamondInit),
            abi.encodeCall(DiamondInit.init, ("Echidna RWA", "eRWA"))
        );

        // Grant roles
        AccessControlFacet(diamond).grantRole(keccak256("ISSUER_ROLE"), address(this));
        AccessControlFacet(diamond).grantRole(keccak256("TRANSFER_AGENT"), address(this));

        // Register asset
        uint16[] memory countries = new uint16[](0);
        address[] memory modules = new address[](0);
        tokenId = AssetManagerFacet(diamond).registerAsset(
            IAssetManager.RegisterAssetParams("Test Bond", "TBND", "", SUPPLY_CAP, 0, modules, address(this), countries)
        );
    }

    function _buildCuts() internal returns (IDiamond.FacetCut[] memory cuts) {
        cuts = new IDiamond.FacetCut[](6);
        uint256 ci;

        // Loupe
        {
            DiamondLoupeFacet f = new DiamondLoupeFacet();
            bytes4[] memory s = new bytes4[](5);
            s[0] = f.facets.selector;
            s[1] = f.facetFunctionSelectors.selector;
            s[2] = f.facetAddresses.selector;
            s[3] = f.facetAddress.selector;
            s[4] = f.supportsInterface.selector;
            cuts[ci++] = IDiamond.FacetCut(address(f), IDiamond.FacetCutAction.Add, s);
        }

        // Ownership + AccessControl
        {
            OwnershipFacet own = new OwnershipFacet();
            bytes4[] memory s = new bytes4[](4);
            s[0] = own.transferOwnership.selector;
            s[1] = own.acceptOwnership.selector;
            s[2] = own.owner.selector;
            s[3] = own.pendingOwner.selector;
            cuts[ci++] = IDiamond.FacetCut(address(own), IDiamond.FacetCutAction.Add, s);
        }
        {
            AccessControlFacet f = new AccessControlFacet();
            bytes4[] memory s = new bytes4[](6);
            s[0] = f.grantRole.selector;
            s[1] = f.revokeRole.selector;
            s[2] = f.renounceRole.selector;
            s[3] = f.setRoleAdmin.selector;
            s[4] = f.hasRole.selector;
            s[5] = f.getRoleAdmin.selector;
            cuts[ci++] = IDiamond.FacetCut(address(f), IDiamond.FacetCutAction.Add, s);
        }

        // ERC1155 + Supply
        {
            ERC1155Facet f = new ERC1155Facet();
            bytes4[] memory s = new bytes4[](7);
            s[0] = f.safeTransferFrom.selector;
            s[1] = f.safeBatchTransferFrom.selector;
            s[2] = f.setApprovalForAll.selector;
            s[3] = f.balanceOf.selector;
            s[4] = f.balanceOfBatch.selector;
            s[5] = f.isApprovedForAll.selector;
            s[6] = f.partitionBalanceOf.selector;
            cuts[ci++] = IDiamond.FacetCut(address(f), IDiamond.FacetCutAction.Add, s);
        }
        {
            SupplyFacet f = new SupplyFacet();
            bytes4[] memory s = new bytes4[](7);
            s[0] = f.mint.selector;
            s[1] = f.batchMint.selector;
            s[2] = f.burn.selector;
            s[3] = f.forcedTransfer.selector;
            s[4] = f.totalSupply.selector;
            s[5] = f.holderCount.selector;
            s[6] = f.isHolder.selector;
            cuts[ci++] = IDiamond.FacetCut(address(f), IDiamond.FacetCutAction.Add, s);
        }

        // AssetManager (for registerAsset + getComplianceModules)
        {
            AssetManagerFacet f = new AssetManagerFacet();
            bytes4[] memory s = new bytes4[](14);
            s[0] = f.registerAsset.selector;
            s[1] = f.addComplianceModule.selector;
            s[2] = f.removeComplianceModule.selector;
            s[3] = f.setComplianceModules.selector;
            s[4] = f.setIdentityProfile.selector;
            s[5] = f.setIssuer.selector;
            s[6] = f.setSupplyCap.selector;
            s[7] = f.setAllowedCountries.selector;
            s[8] = f.setAssetUri.selector;
            s[9] = f.getAssetConfig.selector;
            s[10] = f.getComplianceModules.selector;
            s[11] = f.getRegisteredTokenIds.selector;
            s[12] = f.assetExists.selector;
            s[13] = f.nextTokenId.selector;
            cuts[ci++] = IDiamond.FacetCut(address(f), IDiamond.FacetCutAction.Add, s);
        }
    }

    /*//////////////////////////////////////////////////////////////
                         ACTIONS (fuzzer calls these)
    //////////////////////////////////////////////////////////////*/

    function action_mint(uint256 amount) public {
        amount = (amount % 999) + 1;
        address to = _pickTarget(amount);
        try SupplyFacet(diamond).mint(tokenId, to, amount) {
            ghost_minted += amount;
        } catch {}
    }

    function action_burn(uint256 seed) public {
        address from = _pickTarget(seed);
        uint256 bal = ERC1155Facet(diamond).balanceOf(from, tokenId);
        if (bal == 0) return;
        uint256 amount = (seed % bal) + 1;
        try SupplyFacet(diamond).burn(tokenId, from, amount) {
            ghost_burned += amount;
        } catch {}
    }

    function action_transfer(uint256 fromSeed, uint256 toSeed, uint256 amount) public {
        address from = _pickTarget(fromSeed);
        address to = _pickTarget(toSeed);
        if (from == to) return;
        uint256 bal = ERC1155Facet(diamond).balanceOf(from, tokenId);
        if (bal == 0) return;
        amount = (amount % bal) + 1;
        try ERC1155Facet(diamond).safeTransferFrom(from, to, tokenId, amount, "") {} catch {}
    }

    function action_forcedTransfer(uint256 fromSeed, uint256 toSeed, uint256 amount) public {
        address from = _pickTarget(fromSeed);
        address to = _pickTarget(toSeed);
        if (from == to) return;
        uint256 bal = ERC1155Facet(diamond).balanceOf(from, tokenId);
        if (bal == 0) return;
        amount = (amount % bal) + 1;
        try SupplyFacet(diamond).forcedTransfer(tokenId, from, to, amount, bytes32("ECHIDNA")) {} catch {}
    }

    /*//////////////////////////////////////////////////////////////
                       ECHIDNA PROPERTIES (no args, return bool)
    //////////////////////////////////////////////////////////////*/

    function echidna_supply_conservation() public view returns (bool) {
        return SupplyFacet(diamond).totalSupply(tokenId) == ghost_minted - ghost_burned;
    }

    function echidna_supply_cap() public view returns (bool) {
        return SupplyFacet(diamond).totalSupply(tokenId) <= SUPPLY_CAP;
    }

    function echidna_balance_sum() public view returns (bool) {
        uint256 sum = ERC1155Facet(diamond).balanceOf(ALICE, tokenId)
            + ERC1155Facet(diamond).balanceOf(BOB, tokenId)
            + ERC1155Facet(diamond).balanceOf(CAROL, tokenId);
        return sum == SupplyFacet(diamond).totalSupply(tokenId);
    }

    function echidna_no_overflow_balance() public view returns (bool) {
        uint256 total = SupplyFacet(diamond).totalSupply(tokenId);
        return ERC1155Facet(diamond).balanceOf(ALICE, tokenId) <= total
            && ERC1155Facet(diamond).balanceOf(BOB, tokenId) <= total
            && ERC1155Facet(diamond).balanceOf(CAROL, tokenId) <= total;
    }

    function echidna_holder_count() public view returns (bool) {
        uint256 counted;
        if (ERC1155Facet(diamond).balanceOf(ALICE, tokenId) > 0) counted++;
        if (ERC1155Facet(diamond).balanceOf(BOB, tokenId) > 0) counted++;
        if (ERC1155Facet(diamond).balanceOf(CAROL, tokenId) > 0) counted++;
        return SupplyFacet(diamond).holderCount(tokenId) == counted;
    }

    /*//////////////////////////////////////////////////////////////
                          INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _pickTarget(uint256 seed) internal pure returns (address) {
        uint256 idx = seed % 3;
        if (idx == 0) return ALICE;
        if (idx == 1) return BOB;
        return CAROL;
    }
}
