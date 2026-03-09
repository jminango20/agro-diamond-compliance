// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {CountryRestrictModule} from "../../src/compliance/modules/CountryRestrictModule.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {IdentityRegistryFacet} from "../../src/facets/identity/IdentityRegistryFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {LibReasonCodes} from "../../src/libraries/LibReasonCodes.sol";

contract CountryRestrictModuleTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");

    CountryRestrictModule internal module;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    IdentityRegistryFacet internal ir;
    AccessControlFacet internal ac;

    uint256 internal constant TOKEN_1 = 1;
    uint16 internal constant US = 840;
    uint16 internal constant BR = 76;
    uint16 internal constant CN = 156;

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ir = IdentityRegistryFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

        module = new CountryRestrictModule(address(d.diamond), owner);

        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);
        am.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_1,
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModule: address(module),
                issuer: owner,
                allowedCountries: countries
            })
        );
        ac.grantRole(keccak256("ISSUER_ROLE"), owner);
        vm.stopPrank();

        // Register identities with countries
        vm.startPrank(owner);
        ir.registerIdentity(alice, makeAddr("aliceONCHAINID"), BR);
        ir.registerIdentity(bob, makeAddr("bobONCHAINID"), US);
        vm.stopPrank();

        // Mint tokens to alice
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);
    }

    /*//////////////////////////////////////////////////////////////
                        RESTRICT / UNRESTRICT
    //////////////////////////////////////////////////////////////*/

    function test_RestrictCountry() public {
        vm.prank(owner);
        module.restrictCountry(TOKEN_1, US);
        assertTrue(module.isRestricted(TOKEN_1, US));
    }

    function test_UnrestrictCountry() public {
        vm.startPrank(owner);
        module.restrictCountry(TOKEN_1, US);
        module.unrestrictCountry(TOKEN_1, US);
        vm.stopPrank();
        assertFalse(module.isRestricted(TOKEN_1, US));
    }

    function test_BatchRestrictCountries() public {
        uint16[] memory countries = new uint16[](2);
        countries[0] = US;
        countries[1] = CN;

        vm.prank(owner);
        module.batchRestrictCountries(TOKEN_1, countries);

        assertTrue(module.isRestricted(TOKEN_1, US));
        assertTrue(module.isRestricted(TOKEN_1, CN));
        assertFalse(module.isRestricted(TOKEN_1, BR));
    }

    function test_RevertWhen_RestrictCountry_NotOwner() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("CountryRestrictModule__OnlyOwner()"));
        module.restrictCountry(TOKEN_1, US);
    }

    /*//////////////////////////////////////////////////////////////
                        canTransfer
    //////////////////////////////////////////////////////////////*/

    function test_CanTransfer_AllowedCountries() public view {
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertTrue(ok);
        assertEq(reason, LibReasonCodes.REASON_OK);
    }

    function test_CanTransfer_RejectsRestrictedReceiver() public {
        vm.prank(owner);
        module.restrictCountry(TOKEN_1, US);

        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertFalse(ok);
        assertEq(reason, LibReasonCodes.REASON_COUNTRY_RESTRICTED);
    }

    function test_CanTransfer_RejectsRestrictedSender() public {
        vm.prank(owner);
        module.restrictCountry(TOKEN_1, BR);

        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertFalse(ok);
        assertEq(reason, LibReasonCodes.REASON_COUNTRY_RESTRICTED);
    }

    function test_CanTransfer_AllowsUnregisteredWallet() public {
        address unknown = makeAddr("unknown");
        // country = 0 for unregistered wallet → not restricted
        (bool ok,) = module.canTransfer(TOKEN_1, alice, unknown, 100, "");
        assertTrue(ok);
    }

    /*//////////////////////////////////////////////////////////////
                    INTEGRATION — TRANSFER BLOCKED
    //////////////////////////////////////////////////////////////*/

    function test_Transfer_BlockedByCountryRestriction() public {
        vm.prank(owner);
        module.restrictCountry(TOKEN_1, US);

        vm.prank(alice);
        vm.expectRevert();
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_Transfer_AllowedAfterUnrestrict() public {
        vm.startPrank(owner);
        module.restrictCountry(TOKEN_1, US);
        module.unrestrictCountry(TOKEN_1, US);
        vm.stopPrank();

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
    }

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ZeroDiamond() public {
        vm.expectRevert(abi.encodeWithSignature("CountryRestrictModule__ZeroDiamond()"));
        new CountryRestrictModule(address(0), owner);
    }
}
