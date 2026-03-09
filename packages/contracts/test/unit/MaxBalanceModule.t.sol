// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {MaxBalanceModule} from "../../src/compliance/modules/MaxBalanceModule.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {LibReasonCodes} from "../../src/libraries/LibReasonCodes.sol";

contract MaxBalanceModuleTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");

    MaxBalanceModule internal module;
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

        module = new MaxBalanceModule(address(d.diamond), owner);

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

        // Mint tokens to alice
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 5000);
    }

    /*//////////////////////////////////////////////////////////////
                        CONFIGURATION
    //////////////////////////////////////////////////////////////*/

    function test_SetMaxBalance() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 1000);
        assertEq(module.maxBalance(TOKEN_1), 1000);
    }

    function test_RevertWhen_SetMaxBalance_NotOwner() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("MaxBalanceModule__OnlyOwner()"));
        module.setMaxBalance(TOKEN_1, 1000);
    }

    /*//////////////////////////////////////////////////////////////
                        canTransfer
    //////////////////////////////////////////////////////////////*/

    function test_CanTransfer_UnlimitedByDefault() public view {
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 5000, "");
        assertTrue(ok);
        assertEq(reason, LibReasonCodes.REASON_OK);
    }

    function test_CanTransfer_WithinLimit() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 1000);

        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 500, "");
        assertTrue(ok);
        assertEq(reason, LibReasonCodes.REASON_OK);
    }

    function test_CanTransfer_ExactLimit() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 1000);

        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 1000, "");
        assertTrue(ok);
        assertEq(reason, LibReasonCodes.REASON_OK);
    }

    function test_CanTransfer_RejectsOverLimit() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 1000);

        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 1001, "");
        assertFalse(ok);
        assertEq(reason, LibReasonCodes.REASON_HOLDING_LIMIT);
    }

    function test_CanTransfer_AccountsForExistingBalance() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 1000);

        // Give bob 600 first
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 600, "");

        // bob has 600, transferring 401 would push to 1001 > 1000
        (bool ok, bytes32 reason) = module.canTransfer(TOKEN_1, alice, bob, 401, "");
        assertFalse(ok);
        assertEq(reason, LibReasonCodes.REASON_HOLDING_LIMIT);

        // 400 should be fine (600 + 400 = 1000)
        (ok, reason) = module.canTransfer(TOKEN_1, alice, bob, 400, "");
        assertTrue(ok);
    }

    /*//////////////////////////////////////////////////////////////
                    INTEGRATION — TRANSFER BLOCKED
    //////////////////////////////////////////////////////////////*/

    function test_Transfer_BlockedByMaxBalance() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 500);

        vm.prank(alice);
        vm.expectRevert();
        token.safeTransferFrom(alice, bob, TOKEN_1, 501, "");
    }

    function test_Transfer_AllowedWithinMaxBalance() public {
        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, 500);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 500, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 500);
    }

    /*//////////////////////////////////////////////////////////////
                        FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_MaxBalanceEnforcement(uint256 limit, uint256 amount) public {
        vm.assume(limit > 0 && limit <= 5000);
        vm.assume(amount > 0 && amount <= 5000);

        vm.prank(owner);
        module.setMaxBalance(TOKEN_1, limit);

        (bool ok,) = module.canTransfer(TOKEN_1, alice, bob, amount, "");
        assertEq(ok, amount <= limit);
    }

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ZeroDiamond() public {
        vm.expectRevert(abi.encodeWithSignature("MaxBalanceModule__ZeroDiamond()"));
        new MaxBalanceModule(address(0), owner);
    }
}
