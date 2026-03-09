// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {RecoveryFacet} from "../../src/facets/rwa/RecoveryFacet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {FreezeFacet} from "../../src/facets/rwa/FreezeFacet.sol";
import {IdentityRegistryFacet} from "../../src/facets/identity/IdentityRegistryFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

contract RecoveryFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal aliceNew = makeAddr("aliceNew");
    address internal bob = makeAddr("bob");
    address internal recoveryAgent = makeAddr("recoveryAgent");
    address internal attacker = makeAddr("attacker");

    RecoveryFacet internal recovery;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    FreezeFacet internal freeze;
    IdentityRegistryFacet internal ir;
    AccessControlFacet internal ac;

    uint256 internal constant TOKEN_1 = 1;
    uint256 internal constant TOKEN_2 = 2;

    bytes32 internal constant RECOVERY_AGENT = keccak256("RECOVERY_AGENT");
    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    function setUp() public {
        d = deployDiamond(owner);
        recovery = RecoveryFacet(address(d.diamond));
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        freeze = FreezeFacet(address(d.diamond));
        ir = IdentityRegistryFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);

        // Register two assets
        am.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_1,
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModule: address(0),
                issuer: owner,
                allowedCountries: countries
            })
        );
        am.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_2,
                name: "Bond B",
                symbol: "BNDB",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModule: address(0),
                issuer: owner,
                allowedCountries: countries
            })
        );

        // Roles
        ac.grantRole(ISSUER_ROLE, owner);
        ac.grantRole(RECOVERY_AGENT, recoveryAgent);

        // Register alice identity
        ir.registerIdentity(alice, makeAddr("aliceONCHAINID"), 76);

        // Mint tokens to alice
        supply.mint(TOKEN_1, alice, 1000);
        supply.mint(TOKEN_2, alice, 500);

        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        BASIC RECOVERY
    //////////////////////////////////////////////////////////////*/

    function test_RecoverWallet_MigratesBalances() public {
        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        // Old wallet zeroed
        assertEq(token.balanceOf(alice, TOKEN_1), 0);
        assertEq(token.balanceOf(alice, TOKEN_2), 0);

        // New wallet has all balances
        assertEq(token.balanceOf(aliceNew, TOKEN_1), 1000);
        assertEq(token.balanceOf(aliceNew, TOKEN_2), 500);
    }

    function test_RecoverWallet_MigratesIdentity() public {
        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertEq(ir.getCountry(aliceNew), 76);
        assertEq(ir.getIdentity(aliceNew), makeAddr("aliceONCHAINID"));

        // Old wallet identity cleared
        assertEq(ir.getCountry(alice), 0);
        assertEq(ir.getIdentity(alice), address(0));
    }

    function test_RecoverWallet_EmitsEvents() public {
        vm.prank(recoveryAgent);

        vm.expectEmit(true, true, true, true);
        emit RecoveryFacet.WalletRecovered(alice, aliceNew, recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);
    }

    function test_RecoverWallet_EmitsTokensRecoveredPerAsset() public {
        vm.prank(recoveryAgent);

        vm.expectEmit(true, true, true, true);
        emit RecoveryFacet.TokensRecovered(TOKEN_1, alice, aliceNew, 1000, 0, 0, 0);
        vm.expectEmit(true, true, true, true);
        emit RecoveryFacet.TokensRecovered(TOKEN_2, alice, aliceNew, 500, 0, 0, 0);
        recovery.recoverWallet(alice, aliceNew);
    }

    function test_RecoverWallet_ByOwner() public {
        vm.prank(owner);
        recovery.recoverWallet(alice, aliceNew);
        assertEq(token.balanceOf(aliceNew, TOKEN_1), 1000);
    }

    /*//////////////////////////////////////////////////////////////
                    HOLDER TRACKING
    //////////////////////////////////////////////////////////////*/

    function test_RecoverWallet_UpdatesHolderCount() public {
        assertEq(supply.holderCount(TOKEN_1), 1);
        assertTrue(supply.isHolder(TOKEN_1, alice));

        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertEq(supply.holderCount(TOKEN_1), 1);
        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertTrue(supply.isHolder(TOKEN_1, aliceNew));
    }

    function test_RecoverWallet_NewWalletAlreadyHoldsTokens() public {
        // Give bob some TOKEN_1
        vm.prank(owner);
        supply.mint(TOKEN_1, bob, 200);

        // bob has TOKEN_1, alice has TOKEN_1 + TOKEN_2
        assertEq(supply.holderCount(TOKEN_1), 2);

        // Recover alice → bob (bob already holds TOKEN_1)
        // Need bob to not have identity registered for this to work
        // Actually bob doesn't have identity registered, so this should pass
        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, bob);

        // bob now has 200 + 1000 = 1200 for TOKEN_1
        assertEq(token.balanceOf(bob, TOKEN_1), 1200);
        // holder count should be 1 (only bob), not 2
        assertEq(supply.holderCount(TOKEN_1), 1);
    }

    /*//////////////////////////////////////////////////////////////
                    FREEZE STATE MIGRATION
    //////////////////////////////////////////////////////////////*/

    function test_RecoverWallet_MigratesGlobalFreeze() public {
        vm.prank(owner);
        freeze.setWalletFrozen(alice, true);

        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertTrue(freeze.isWalletFrozen(aliceNew));
        assertFalse(freeze.isWalletFrozen(alice));
    }

    function test_RecoverWallet_MigratesAssetFreeze() public {
        vm.prank(owner);
        freeze.setAssetWalletFrozen(TOKEN_1, alice, true);

        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertTrue(freeze.isAssetWalletFrozen(TOKEN_1, aliceNew));
        assertFalse(freeze.isAssetWalletFrozen(TOKEN_1, alice));
    }

    function test_RecoverWallet_MigratesFrozenAmount() public {
        vm.prank(owner);
        freeze.setFrozenAmount(TOKEN_1, alice, 300);

        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertEq(freeze.getFrozenAmount(TOKEN_1, aliceNew), 300);
        assertEq(freeze.getFrozenAmount(TOKEN_1, alice), 0);
    }

    function test_RecoverWallet_MigratesLockup() public {
        uint64 expiry = uint64(block.timestamp + 365 days);
        vm.prank(owner);
        freeze.setLockupExpiry(TOKEN_1, alice, expiry);

        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        assertEq(freeze.getLockupExpiry(TOKEN_1, aliceNew), expiry);
        assertEq(freeze.getLockupExpiry(TOKEN_1, alice), 0);
    }

    /*//////////////////////////////////////////////////////////////
                    WALLET WITH NO BALANCES
    //////////////////////////////////////////////////////////////*/

    function test_RecoverWallet_NoBalances() public {
        address emptyWallet = makeAddr("emptyWallet");
        address newEmpty = makeAddr("newEmpty");

        vm.prank(owner);
        ir.registerIdentity(emptyWallet, makeAddr("emptyONCHAINID"), 840);

        vm.prank(recoveryAgent);
        recovery.recoverWallet(emptyWallet, newEmpty);

        // Identity migrated even with no tokens
        assertEq(ir.getCountry(newEmpty), 840);
        assertEq(ir.getIdentity(newEmpty), makeAddr("emptyONCHAINID"));
    }

    /*//////////////////////////////////////////////////////////////
                        ACCESS CONTROL
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_Unauthorized() public {
        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("RecoveryFacet__Unauthorized()"));
        recovery.recoverWallet(alice, aliceNew);
    }

    /*//////////////////////////////////////////////////////////////
                        VALIDATION
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ZeroAddress_Lost() public {
        vm.prank(recoveryAgent);
        vm.expectRevert(abi.encodeWithSignature("RecoveryFacet__ZeroAddress()"));
        recovery.recoverWallet(address(0), aliceNew);
    }

    function test_RevertWhen_ZeroAddress_New() public {
        vm.prank(recoveryAgent);
        vm.expectRevert(abi.encodeWithSignature("RecoveryFacet__ZeroAddress()"));
        recovery.recoverWallet(alice, address(0));
    }

    function test_RevertWhen_SameAddress() public {
        vm.prank(recoveryAgent);
        vm.expectRevert(abi.encodeWithSignature("RecoveryFacet__SameAddress()"));
        recovery.recoverWallet(alice, alice);
    }

    function test_RevertWhen_NewWalletAlreadyHasIdentity() public {
        vm.prank(owner);
        ir.registerIdentity(aliceNew, makeAddr("newONCHAINID"), 840);

        vm.prank(recoveryAgent);
        vm.expectRevert(abi.encodeWithSignature("RecoveryFacet__NewWalletAlreadyRegistered()"));
        recovery.recoverWallet(alice, aliceNew);
    }

    /*//////////////////////////////////////////////////////////////
                    POST-RECOVERY USABILITY
    //////////////////////////////////////////////////////////////*/

    function test_RecoveredWallet_CanTransfer() public {
        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        // aliceNew should be able to transfer normally
        vm.prank(aliceNew);
        token.safeTransferFrom(aliceNew, bob, TOKEN_1, 100, "");

        assertEq(token.balanceOf(aliceNew, TOKEN_1), 900);
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
    }

    function test_RecoveredWallet_CanReceiveMint() public {
        vm.prank(recoveryAgent);
        recovery.recoverWallet(alice, aliceNew);

        vm.prank(owner);
        supply.mint(TOKEN_1, aliceNew, 500);

        assertEq(token.balanceOf(aliceNew, TOKEN_1), 1500);
    }
}
