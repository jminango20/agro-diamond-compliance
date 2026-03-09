// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {MockComplianceModule} from "../helpers/MockComplianceModule.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {FreezeFacet} from "../../src/facets/rwa/FreezeFacet.sol";
import {PauseFacet} from "../../src/facets/security/PauseFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {LibReasonCodes} from "../../src/libraries/LibReasonCodes.sol";

contract SupplyFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal issuer = makeAddr("issuer");
    address internal agent = makeAddr("agent");
    address internal attacker = makeAddr("attacker");

    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    FreezeFacet internal freeze;
    PauseFacet internal pause;
    AccessControlFacet internal ac;
    MockComplianceModule internal mockModule;

    uint256 internal constant TOKEN_1 = 1;
    uint256 internal constant TOKEN_2 = 2;

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        freeze = FreezeFacet(address(d.diamond));
        pause = PauseFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));
        mockModule = new MockComplianceModule();

        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);
        am.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_1,
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 10_000,
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
        ac.grantRole(ISSUER_ROLE, issuer);
        ac.grantRole(TRANSFER_AGENT, agent);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                            MINT
    //////////////////////////////////////////////////////////////*/

    function test_Mint() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 500);

        assertEq(token.balanceOf(alice, TOKEN_1), 500);
        assertEq(supply.totalSupply(TOKEN_1), 500);
        assertEq(supply.holderCount(TOKEN_1), 1);
        assertTrue(supply.isHolder(TOKEN_1, alice));
    }

    function test_Mint_ByIssuerRole() public {
        vm.prank(issuer);
        supply.mint(TOKEN_1, alice, 100);
        assertEq(token.balanceOf(alice, TOKEN_1), 100);
    }

    function test_Mint_EmitsEvent() public {
        vm.prank(owner);
        vm.expectEmit(true, true, false, true);
        emit SupplyFacet.Minted(TOKEN_1, alice, 500);
        supply.mint(TOKEN_1, alice, 500);
    }

    function test_Mint_MultipleTimes_AccumulatesSupply() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 300);
        supply.mint(TOKEN_1, bob, 200);
        vm.stopPrank();

        assertEq(supply.totalSupply(TOKEN_1), 500);
        assertEq(supply.holderCount(TOKEN_1), 2);
    }

    function test_Mint_UnlimitedSupply() public {
        vm.prank(owner);
        supply.mint(TOKEN_2, alice, type(uint128).max);
        assertEq(token.balanceOf(alice, TOKEN_2), type(uint128).max);
    }

    function test_RevertWhen_Mint_Unauthorized() public {
        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__Unauthorized()"));
        supply.mint(TOKEN_1, alice, 100);
    }

    function test_RevertWhen_Mint_ToZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__MintToZeroAddress()"));
        supply.mint(TOKEN_1, address(0), 100);
    }

    function test_RevertWhen_Mint_ExceedsSupplyCap() public {
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSignature("SupplyFacet__SupplyCapExceeded(uint256,uint256,uint256)", TOKEN_1, 0, 10_000)
        );
        supply.mint(TOKEN_1, alice, 10_001);
    }

    function test_RevertWhen_Mint_ExceedsSupplyCap_Accumulated() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 9_000);

        vm.expectRevert(
            abi.encodeWithSignature("SupplyFacet__SupplyCapExceeded(uint256,uint256,uint256)", TOKEN_1, 9_000, 10_000)
        );
        supply.mint(TOKEN_1, bob, 1_001);
        vm.stopPrank();
    }

    function test_RevertWhen_Mint_AssetNotRegistered() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__AssetNotRegistered(uint256)", 999));
        supply.mint(999, alice, 100);
    }

    function test_RevertWhen_Mint_ProtocolPaused() public {
        vm.prank(owner);
        pause.pauseProtocol();

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__ProtocolPaused()"));
        supply.mint(TOKEN_1, alice, 100);
    }

    function test_RevertWhen_Mint_AssetPaused() public {
        vm.prank(owner);
        pause.pauseAsset(TOKEN_1);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__AssetPaused(uint256)", TOKEN_1));
        supply.mint(TOKEN_1, alice, 100);
    }

    function test_RevertWhen_Mint_ReceiverGloballyFrozen() public {
        vm.prank(owner);
        freeze.setWalletFrozen(alice, true);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__WalletFrozenGlobal(address)", alice));
        supply.mint(TOKEN_1, alice, 100);
    }

    function test_RevertWhen_Mint_ReceiverAssetFrozen() public {
        vm.prank(owner);
        freeze.setAssetWalletFrozen(TOKEN_1, alice, true);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__WalletFrozenAsset(uint256,address)", TOKEN_1, alice));
        supply.mint(TOKEN_1, alice, 100);
    }

    /*//////////////////////////////////////////////////////////////
                        BATCH MINT
    //////////////////////////////////////////////////////////////*/

    function test_BatchMint() public {
        uint256[] memory ids = new uint256[](2);
        ids[0] = TOKEN_1;
        ids[1] = TOKEN_2;
        address[] memory recipients = new address[](2);
        recipients[0] = alice;
        recipients[1] = bob;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 100;
        amounts[1] = 200;

        vm.prank(owner);
        supply.batchMint(ids, recipients, amounts);

        assertEq(token.balanceOf(alice, TOKEN_1), 100);
        assertEq(token.balanceOf(bob, TOKEN_2), 200);
        assertEq(supply.holderCount(TOKEN_1), 1);
        assertEq(supply.holderCount(TOKEN_2), 1);
    }

    function test_RevertWhen_BatchMint_ArrayLengthMismatch() public {
        uint256[] memory ids = new uint256[](2);
        address[] memory recipients = new address[](1);
        uint256[] memory amounts = new uint256[](2);

        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__ArrayLengthMismatch()"));
        supply.batchMint(ids, recipients, amounts);
    }

    /*//////////////////////////////////////////////////////////////
                            BURN
    //////////////////////////////////////////////////////////////*/

    function test_Burn() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 1000);
        supply.burn(TOKEN_1, alice, 400);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), 600);
        assertEq(supply.totalSupply(TOKEN_1), 600);
        assertTrue(supply.isHolder(TOKEN_1, alice));
    }

    function test_Burn_EmitsEvent() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.expectEmit(true, true, false, true);
        emit SupplyFacet.Burned(TOKEN_1, alice, 300);
        supply.burn(TOKEN_1, alice, 300);
        vm.stopPrank();
    }

    function test_Burn_FullBalance_RemovesHolder() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 500);
        assertEq(supply.holderCount(TOKEN_1), 1);

        supply.burn(TOKEN_1, alice, 500);
        vm.stopPrank();

        assertEq(supply.holderCount(TOKEN_1), 0);
        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertEq(supply.totalSupply(TOKEN_1), 0);
    }

    function test_Burn_ByIssuerRole() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(issuer);
        supply.burn(TOKEN_1, alice, 200);
        assertEq(token.balanceOf(alice, TOKEN_1), 800);
    }

    function test_RevertWhen_Burn_Unauthorized() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__Unauthorized()"));
        supply.burn(TOKEN_1, alice, 100);
    }

    function test_RevertWhen_Burn_FromZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__BurnFromZeroAddress()"));
        supply.burn(TOKEN_1, address(0), 100);
    }

    function test_RevertWhen_Burn_InsufficientBalance() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 500);

        vm.expectRevert(
            abi.encodeWithSignature(
                "SupplyFacet__InsufficientFreeBalance(uint256,address,uint256,uint256)", TOKEN_1, alice, 500, 600
            )
        );
        supply.burn(TOKEN_1, alice, 600);
        vm.stopPrank();
    }

    function test_Burn_RespectsPartialFreeze() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 1000);
        freeze.setFrozenAmount(TOKEN_1, alice, 800);

        // Available = 1000 - 800 = 200; burning 200 should work
        supply.burn(TOKEN_1, alice, 200);
        assertEq(token.balanceOf(alice, TOKEN_1), 800);
        vm.stopPrank();
    }

    function test_RevertWhen_Burn_ExceedsAvailableAfterFreeze() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 1000);
        freeze.setFrozenAmount(TOKEN_1, alice, 800);

        vm.expectRevert(
            abi.encodeWithSignature(
                "SupplyFacet__InsufficientFreeBalance(uint256,address,uint256,uint256)", TOKEN_1, alice, 200, 300
            )
        );
        supply.burn(TOKEN_1, alice, 300);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                    FORCED TRANSFER
    //////////////////////////////////////////////////////////////*/

    function test_ForcedTransfer() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        bytes32 reason = keccak256("REGULATORY_SEIZURE");
        vm.prank(agent);
        supply.forcedTransfer(TOKEN_1, alice, bob, 400, reason);

        assertEq(token.balanceOf(alice, TOKEN_1), 600);
        assertEq(token.balanceOf(bob, TOKEN_1), 400);
        assertEq(supply.holderCount(TOKEN_1), 2);
    }

    function test_ForcedTransfer_EmitsEvent() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        bytes32 reason = keccak256("RECOVERY");
        vm.prank(agent);
        vm.expectEmit(true, true, true, true);
        emit SupplyFacet.ForcedTransfer(TOKEN_1, alice, bob, 500, reason);
        supply.forcedTransfer(TOKEN_1, alice, bob, 500, reason);
    }

    function test_ForcedTransfer_ByOwner() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(owner);
        supply.forcedTransfer(TOKEN_1, alice, bob, 300, bytes32(0));
        assertEq(token.balanceOf(bob, TOKEN_1), 300);
    }

    function test_ForcedTransfer_FullBalance_UpdatesHolders() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(agent);
        supply.forcedTransfer(TOKEN_1, alice, bob, 1000, bytes32(0));

        assertEq(supply.holderCount(TOKEN_1), 1);
        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertTrue(supply.isHolder(TOKEN_1, bob));
    }

    function test_RevertWhen_ForcedTransfer_Unauthorized() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__Unauthorized()"));
        supply.forcedTransfer(TOKEN_1, alice, bob, 100, bytes32(0));
    }

    function test_RevertWhen_ForcedTransfer_ToZeroAddress() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1000);

        vm.prank(agent);
        vm.expectRevert(abi.encodeWithSignature("SupplyFacet__TransferToZeroAddress()"));
        supply.forcedTransfer(TOKEN_1, alice, address(0), 100, bytes32(0));
    }

    function test_RevertWhen_ForcedTransfer_InsufficientBalance() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 500);

        vm.prank(agent);
        vm.expectRevert(
            abi.encodeWithSignature(
                "SupplyFacet__InsufficientFreeBalance(uint256,address,uint256,uint256)", TOKEN_1, alice, 500, 600
            )
        );
        supply.forcedTransfer(TOKEN_1, alice, bob, 600, bytes32(0));
    }

    /*//////////////////////////////////////////////////////////////
                    COMPLIANCE MODULE HOOKS
    //////////////////////////////////////////////////////////////*/

    function test_Mint_CallsComplianceMintedHook() public {
        vm.prank(owner);
        am.setComplianceModule(TOKEN_1, address(mockModule));

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 100);

        assertEq(mockModule.mintedCount(), 1);
    }

    function test_Burn_CallsComplianceBurnedHook() public {
        vm.prank(owner);
        am.setComplianceModule(TOKEN_1, address(mockModule));

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 500);
        supply.burn(TOKEN_1, alice, 200);
        vm.stopPrank();

        assertEq(mockModule.burnedCount(), 1);
    }

    function test_ForcedTransfer_CallsComplianceTransferredHook() public {
        vm.prank(owner);
        am.setComplianceModule(TOKEN_1, address(mockModule));

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 500);

        vm.prank(agent);
        supply.forcedTransfer(TOKEN_1, alice, bob, 200, bytes32(0));

        assertEq(mockModule.transferredCount(), 1);
    }

    /*//////////////////////////////////////////////////////////////
                        VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function test_TotalSupply_ZeroInitially() public view {
        assertEq(supply.totalSupply(TOKEN_1), 0);
    }

    function test_HolderCount_ZeroInitially() public view {
        assertEq(supply.holderCount(TOKEN_1), 0);
    }

    function test_IsHolder_FalseInitially() public view {
        assertFalse(supply.isHolder(TOKEN_1, alice));
    }

    /*//////////////////////////////////////////////////////////////
                            FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_MintAndBurn(uint256 mintAmount, uint256 burnAmount) public {
        vm.assume(mintAmount > 0 && mintAmount <= 10_000);
        vm.assume(burnAmount > 0 && burnAmount <= mintAmount);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, mintAmount);
        supply.burn(TOKEN_1, alice, burnAmount);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), mintAmount - burnAmount);
        assertEq(supply.totalSupply(TOKEN_1), mintAmount - burnAmount);
    }

    function testFuzz_ForcedTransfer(uint256 mintAmount, uint256 transferAmount) public {
        vm.assume(mintAmount > 0 && mintAmount <= 10_000);
        vm.assume(transferAmount > 0 && transferAmount <= mintAmount);

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, mintAmount);

        vm.prank(agent);
        supply.forcedTransfer(TOKEN_1, alice, bob, transferAmount, bytes32(0));

        assertEq(token.balanceOf(alice, TOKEN_1), mintAmount - transferAmount);
        assertEq(token.balanceOf(bob, TOKEN_1), transferAmount);
    }
}
