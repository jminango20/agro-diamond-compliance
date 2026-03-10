// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

/// @title SupplyFuzz
/// @notice Fuzz tests for mint, burn, and forcedTransfer with randomized inputs.
contract SupplyFuzz is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");

    SupplyFacet internal supply;
    ERC1155Facet internal token;

    uint256 internal constant TOKEN_1 = 1;
    uint256 internal constant SUPPLY_CAP = 1_000_000;

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));

        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);
        AssetManagerFacet(address(d.diamond)).registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_1,
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: SUPPLY_CAP,
                identityProfileId: 0,
                complianceModule: address(0),
                issuer: owner,
                allowedCountries: countries
            })
        );
        AccessControlFacet(address(d.diamond)).grantRole(TRANSFER_AGENT, owner);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                            MINT FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_MintIncreasesBalanceAndSupply(uint256 amount) external {
        amount = bound(amount, 1, SUPPLY_CAP);

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, amount);

        assertEq(token.balanceOf(alice, TOKEN_1), amount);
        assertEq(supply.totalSupply(TOKEN_1), amount);
        assertTrue(supply.isHolder(TOKEN_1, alice));
        assertEq(supply.holderCount(TOKEN_1), 1);
    }

    function testFuzz_MintTwiceAccumulates(uint256 a, uint256 b) external {
        a = bound(a, 1, SUPPLY_CAP / 2);
        b = bound(b, 1, SUPPLY_CAP / 2);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, a);
        supply.mint(TOKEN_1, alice, b);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), a + b);
        assertEq(supply.totalSupply(TOKEN_1), a + b);
        assertEq(supply.holderCount(TOKEN_1), 1);
    }

    function testFuzz_MintToTwoHolders(uint256 a, uint256 b) external {
        a = bound(a, 1, SUPPLY_CAP / 2);
        b = bound(b, 1, SUPPLY_CAP / 2);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, a);
        supply.mint(TOKEN_1, bob, b);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), a);
        assertEq(token.balanceOf(bob, TOKEN_1), b);
        assertEq(supply.totalSupply(TOKEN_1), a + b);
        assertEq(supply.holderCount(TOKEN_1), 2);
    }

    function testFuzz_MintRevertsAboveSupplyCap(uint256 amount) external {
        amount = bound(amount, SUPPLY_CAP + 1, type(uint256).max / 2);

        vm.prank(owner);
        vm.expectRevert();
        supply.mint(TOKEN_1, alice, amount);
    }

    /*//////////////////////////////////////////////////////////////
                            BURN FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_BurnDecreasesBalanceAndSupply(uint256 mintAmt, uint256 burnAmt) external {
        mintAmt = bound(mintAmt, 1, SUPPLY_CAP);
        burnAmt = bound(burnAmt, 1, mintAmt);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, mintAmt);
        supply.burn(TOKEN_1, alice, burnAmt);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), mintAmt - burnAmt);
        assertEq(supply.totalSupply(TOKEN_1), mintAmt - burnAmt);
    }

    function testFuzz_BurnEntireBalance_RemovesHolder(uint256 amount) external {
        amount = bound(amount, 1, SUPPLY_CAP);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, amount);
        supply.burn(TOKEN_1, alice, amount);
        vm.stopPrank();

        assertEq(token.balanceOf(alice, TOKEN_1), 0);
        assertEq(supply.totalSupply(TOKEN_1), 0);
        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertEq(supply.holderCount(TOKEN_1), 0);
    }

    function testFuzz_BurnRevertsIfExceedsBalance(uint256 mintAmt, uint256 burnAmt) external {
        mintAmt = bound(mintAmt, 1, SUPPLY_CAP);
        burnAmt = bound(burnAmt, mintAmt + 1, type(uint256).max / 2);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, mintAmt);
        vm.expectRevert();
        supply.burn(TOKEN_1, alice, burnAmt);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        FORCED TRANSFER FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_ForcedTransferConservesSupply(uint256 mintAmt, uint256 xferAmt) external {
        mintAmt = bound(mintAmt, 1, SUPPLY_CAP);
        xferAmt = bound(xferAmt, 1, mintAmt);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, mintAmt);
        uint256 supplyBefore = supply.totalSupply(TOKEN_1);

        supply.forcedTransfer(TOKEN_1, alice, bob, xferAmt, bytes32("FUZZ"));
        vm.stopPrank();

        assertEq(supply.totalSupply(TOKEN_1), supplyBefore);
        assertEq(token.balanceOf(alice, TOKEN_1), mintAmt - xferAmt);
        assertEq(token.balanceOf(bob, TOKEN_1), xferAmt);
    }

    function testFuzz_ForcedTransferFullBalance_UpdatesHolders(uint256 amount) external {
        amount = bound(amount, 1, SUPPLY_CAP);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, amount);
        supply.forcedTransfer(TOKEN_1, alice, bob, amount, bytes32("FUZZ"));
        vm.stopPrank();

        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertTrue(supply.isHolder(TOKEN_1, bob));
        assertEq(supply.holderCount(TOKEN_1), 1);
    }

    /*//////////////////////////////////////////////////////////////
                        TRANSFER FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_TransferConservesSupply(uint256 mintAmt, uint256 xferAmt) external {
        mintAmt = bound(mintAmt, 1, SUPPLY_CAP);
        xferAmt = bound(xferAmt, 1, mintAmt);

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, mintAmt);

        uint256 supplyBefore = supply.totalSupply(TOKEN_1);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, xferAmt, "");

        assertEq(supply.totalSupply(TOKEN_1), supplyBefore);
        assertEq(token.balanceOf(alice, TOKEN_1) + token.balanceOf(bob, TOKEN_1), mintAmt);
    }

    function testFuzz_TransferFullBalance_UpdatesHolders(uint256 amount) external {
        amount = bound(amount, 1, SUPPLY_CAP);

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, amount);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, amount, "");

        assertFalse(supply.isHolder(TOKEN_1, alice));
        assertTrue(supply.isHolder(TOKEN_1, bob));
        assertEq(supply.holderCount(TOKEN_1), 1);
    }

    /*//////////////////////////////////////////////////////////////
                    MINT + BURN CYCLE FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_MintBurnCycleReturnsToZero(uint256 amount) external {
        amount = bound(amount, 1, SUPPLY_CAP);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, amount);
        supply.burn(TOKEN_1, alice, amount);
        vm.stopPrank();

        assertEq(supply.totalSupply(TOKEN_1), 0);
        assertEq(supply.holderCount(TOKEN_1), 0);
        assertEq(token.balanceOf(alice, TOKEN_1), 0);
    }
}
