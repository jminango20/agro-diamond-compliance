// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {MaxHoldersModule, REASON_MAX_HOLDERS} from "../../src/compliance/modules/MaxHoldersModule.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {LibReasonCodes} from "../../src/libraries/LibReasonCodes.sol";

contract MaxHoldersModuleTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal carol = makeAddr("carol");
    address internal dave = makeAddr("dave");

    MaxHoldersModule internal module;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    AccessControlFacet internal ac;

    uint256 internal constant TOKEN_1 = 1;

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

        module = new MaxHoldersModule(address(d.diamond), owner);

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

        // Mint tokens to alice (1 holder)
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);
    }

    /*//////////////////////////////////////////////////////////////
                        CONFIGURATION
    //////////////////////////////////////////////////////////////*/

    function test_SetMaxHolders() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 5);
        assertEq(module.maxHolders(TOKEN_1), 5);
    }

    function test_RevertWhen_SetMaxHolders_NotOwner() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("MaxHoldersModule__OnlyOwner()"));
        module.setMaxHolders(TOKEN_1, 5);
    }

    /*//////////////////////////////////////////////////////////////
                        canTransfer
    //////////////////////////////////////////////////////////////*/

    function test_CanTransfer_UnlimitedByDefault() public view {
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertTrue(ok);
        assertEq(reason, LibReasonCodes.REASON_OK);
    }

    function test_CanTransfer_NewHolderWithinLimit() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 3);

        // alice is holder 1, bob would be holder 2
        (bool ok,) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertTrue(ok);
    }

    function test_CanTransfer_ExistingHolderDoesNotCount() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 2);

        // Transfer to bob (holder 2)
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");

        // bob already holds tokens — transfer to bob again should pass
        (bool ok,) = module.canTransfer(TOKEN_1, alice, bob, 50, "");
        assertTrue(ok);
    }

    function test_CanTransfer_RejectsNewHolderOverLimit() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 2);

        // alice=holder1, bob=holder2
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");

        // carol would be holder 3 → rejected
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, carol, 100, "");
        assertFalse(ok);
        assertEq(reason, REASON_MAX_HOLDERS);
    }

    function test_CanTransfer_LimitOfOne() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 1);

        // alice is holder 1, bob would be holder 2 → rejected
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 100, "");
        assertFalse(ok);
        assertEq(reason, REASON_MAX_HOLDERS);
    }

    /*//////////////////////////////////////////////////////////////
                    INTEGRATION — TRANSFER BLOCKED
    //////////////////////////////////////////////////////////////*/

    function test_Transfer_BlockedWhenMaxHoldersReached() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 2);

        // alice→bob OK (2 holders)
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");

        // alice→carol BLOCKED (would be 3 holders)
        vm.prank(alice);
        vm.expectRevert();
        token.safeTransferFrom(alice, carol, TOKEN_1, 100, "");
    }

    function test_Transfer_AllowedToExistingHolder() public {
        vm.prank(owner);
        module.setMaxHolders(TOKEN_1, 2);

        vm.startPrank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
        // bob already holds → should pass even at limit
        token.safeTransferFrom(alice, bob, TOKEN_1, 200, "");
        vm.stopPrank();

        assertEq(token.balanceOf(bob, TOKEN_1), 300);
    }

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ZeroDiamond() public {
        vm.expectRevert(abi.encodeWithSignature("MaxHoldersModule__ZeroDiamond()"));
        new MaxHoldersModule(address(0), owner);
    }
}
