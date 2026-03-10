// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {DividendFacet} from "../../src/facets/rwa/DividendFacet.sol";
import {SnapshotFacet} from "../../src/facets/rwa/SnapshotFacet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

/// @dev Minimal ERC-20 mock for testing dividend payments
contract MockERC20 {
    string public name = "Mock Token";
    string public symbol = "MOCK";
    uint8 public decimals = 18;
    uint256 public totalSupply;
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    function mint(address to, uint256 amount) external {
        balanceOf[to] += amount;
        totalSupply += amount;
    }

    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        return true;
    }

    function transfer(address to, uint256 amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        allowance[from][msg.sender] -= amount;
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        return true;
    }
}

contract DividendFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal issuer = makeAddr("issuer");
    address internal attacker = makeAddr("attacker");

    DividendFacet internal dividend;
    SnapshotFacet internal snapshot;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    AccessControlFacet internal ac;
    MockERC20 internal payToken;

    uint256 internal constant TOKEN_1 = 1;
    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    uint256 internal snapshotId;

    function setUp() public {
        d = deployDiamond(owner);
        dividend = DividendFacet(payable(address(d.diamond)));
        snapshot = SnapshotFacet(address(d.diamond));
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));
        payToken = new MockERC20();

        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);
        am.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: TOKEN_1,
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 100_000,
                identityProfileId: 0,
                complianceModule: address(0),
                issuer: owner,
                allowedCountries: countries
            })
        );
        ac.grantRole(ISSUER_ROLE, issuer);

        // Mint: alice=3000, bob=7000  → total 10000
        supply.mint(TOKEN_1, alice, 3000);
        supply.mint(TOKEN_1, bob, 7000);

        // Create snapshot and record holders
        snapshotId = snapshot.createSnapshot(TOKEN_1);
        address[] memory holders = new address[](2);
        holders[0] = alice;
        holders[1] = bob;
        snapshot.recordHoldersBatch(snapshotId, holders);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                    CREATE DIVIDEND — ETH
    //////////////////////////////////////////////////////////////*/

    function test_CreateDividend_ETH() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        (uint256 snapId, uint256 tokenId, uint256 totalAmount, address paymentToken, uint256 claimedAmount, uint64 createdAt) =
            dividend.getDividend(divId);
        assertEq(snapId, snapshotId);
        assertEq(tokenId, TOKEN_1);
        assertEq(totalAmount, 10 ether);
        assertEq(paymentToken, address(0));
        assertEq(claimedAmount, 0);
        assertGt(createdAt, 0);
    }

    function test_CreateDividend_IssuerCanCreate() public {
        vm.deal(issuer, 1 ether);
        vm.prank(issuer);
        uint256 divId = dividend.createDividend{value: 1 ether}(snapshotId, 1 ether, address(0));
        assertEq(divId, 0);
    }

    function test_RevertWhen_CreateDividend_Unauthorized() public {
        vm.deal(attacker, 1 ether);
        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__Unauthorized()"));
        dividend.createDividend{value: 1 ether}(snapshotId, 1 ether, address(0));
    }

    function test_RevertWhen_CreateDividend_ZeroAmount() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__ZeroAmount()"));
        dividend.createDividend(snapshotId, 0, address(0));
    }

    function test_RevertWhen_CreateDividend_SnapshotNotFound() public {
        vm.deal(owner, 1 ether);
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__SnapshotNotFound(uint256)", 999));
        dividend.createDividend{value: 1 ether}(999, 1 ether, address(0));
    }

    function test_RevertWhen_CreateDividend_InsufficientETH() public {
        vm.deal(owner, 1 ether);
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__InsufficientETH()"));
        dividend.createDividend{value: 0.5 ether}(snapshotId, 1 ether, address(0));
    }

    /*//////////////////////////////////////////////////////////////
                    CREATE DIVIDEND — ERC-20
    //////////////////////////////////////////////////////////////*/

    function test_CreateDividend_ERC20() public {
        payToken.mint(owner, 10_000e18);
        vm.startPrank(owner);
        payToken.approve(address(d.diamond), 10_000e18);
        uint256 divId = dividend.createDividend(snapshotId, 10_000e18, address(payToken));
        vm.stopPrank();

        (,, uint256 totalAmount, address pt,,) = dividend.getDividend(divId);
        assertEq(totalAmount, 10_000e18);
        assertEq(pt, address(payToken));
        assertEq(payToken.balanceOf(address(d.diamond)), 10_000e18);
    }

    /*//////////////////////////////////////////////////////////////
                        CLAIM DIVIDEND — ETH
    //////////////////////////////////////////////////////////////*/

    function test_ClaimDividend_ETH_ProRata() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        // Alice: 3000/10000 * 10 ETH = 3 ETH
        uint256 aliceBefore = alice.balance;
        vm.prank(alice);
        dividend.claimDividend(divId);
        assertEq(alice.balance - aliceBefore, 3 ether);

        // Bob: 7000/10000 * 10 ETH = 7 ETH
        uint256 bobBefore = bob.balance;
        vm.prank(bob);
        dividend.claimDividend(divId);
        assertEq(bob.balance - bobBefore, 7 ether);

        // Verify claimed amounts
        assertTrue(dividend.hasClaimed(divId, alice));
        assertTrue(dividend.hasClaimed(divId, bob));
        (,,, , uint256 claimed,) = dividend.getDividend(divId);
        assertEq(claimed, 10 ether);
    }

    /*//////////////////////////////////////////////////////////////
                        CLAIM DIVIDEND — ERC-20
    //////////////////////////////////////////////////////////////*/

    function test_ClaimDividend_ERC20_ProRata() public {
        payToken.mint(owner, 10_000e18);
        vm.startPrank(owner);
        payToken.approve(address(d.diamond), 10_000e18);
        uint256 divId = dividend.createDividend(snapshotId, 10_000e18, address(payToken));
        vm.stopPrank();

        // Alice: 3000/10000 * 10000 = 3000 tokens
        vm.prank(alice);
        dividend.claimDividend(divId);
        assertEq(payToken.balanceOf(alice), 3000e18);

        // Bob: 7000/10000 * 10000 = 7000 tokens
        vm.prank(bob);
        dividend.claimDividend(divId);
        assertEq(payToken.balanceOf(bob), 7000e18);
    }

    /*//////////////////////////////////////////////////////////////
                        CLAIM REVERTS
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ClaimDividend_NotFound() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__DividendNotFound(uint256)", 999));
        dividend.claimDividend(999);
    }

    function test_RevertWhen_ClaimDividend_AlreadyClaimed() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        vm.prank(alice);
        dividend.claimDividend(divId);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__AlreadyClaimed(uint256,address)", divId, alice));
        dividend.claimDividend(divId);
    }

    function test_RevertWhen_ClaimDividend_HolderNotRecorded() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        address unknown = makeAddr("unknown");
        vm.prank(unknown);
        vm.expectRevert(abi.encodeWithSignature("DividendFacet__HolderNotRecorded(uint256,address)", divId, unknown));
        dividend.claimDividend(divId);
    }

    /*//////////////////////////////////////////////////////////////
                            VIEWS
    //////////////////////////////////////////////////////////////*/

    function test_ClaimableAmount() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        assertEq(dividend.claimableAmount(divId, alice), 3 ether);
        assertEq(dividend.claimableAmount(divId, bob), 7 ether);
    }

    function test_ClaimableAmount_ZeroAfterClaim() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        vm.prank(alice);
        dividend.claimDividend(divId);

        assertEq(dividend.claimableAmount(divId, alice), 0);
    }

    function test_ClaimableAmount_ZeroForUnrecordedHolder() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        assertEq(dividend.claimableAmount(divId, makeAddr("unknown")), 0);
    }

    function test_HasClaimed_FalseBeforeClaim() public {
        vm.deal(owner, 10 ether);
        vm.prank(owner);
        uint256 divId = dividend.createDividend{value: 10 ether}(snapshotId, 10 ether, address(0));

        assertFalse(dividend.hasClaimed(divId, alice));
    }

    function test_GetTokenDividends() public {
        vm.deal(owner, 2 ether);
        vm.startPrank(owner);
        dividend.createDividend{value: 1 ether}(snapshotId, 1 ether, address(0));
        dividend.createDividend{value: 1 ether}(snapshotId, 1 ether, address(0));
        vm.stopPrank();

        uint256[] memory ids = dividend.getTokenDividends(TOKEN_1);
        assertEq(ids.length, 2);
        assertEq(ids[0], 0);
        assertEq(ids[1], 1);
    }

    function test_GetTokenDividends_EmptyForUnusedToken() public view {
        uint256[] memory ids = dividend.getTokenDividends(999);
        assertEq(ids.length, 0);
    }
}
