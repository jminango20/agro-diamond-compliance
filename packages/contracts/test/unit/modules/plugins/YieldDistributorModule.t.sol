// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../../../helpers/DiamondHelper.sol";
import {MockERC20, MockNonStandardERC20, MockFailingERC20} from "../../../helpers/MockERC20.sol";
import {MockERC1155} from "../../../helpers/MockERC1155.sol";
import {YieldDistributorModule} from "../../../../src/plugins/modules/YieldDistributorModule.sol";
import {SupplyFacet} from "../../../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../../../src/interfaces/token/IAssetManager.sol";
import {IHookablePlugin} from "../../../../src/interfaces/plugins/IHookablePlugin.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract YieldDistributorModuleTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal carol = makeAddr("carol");
    address internal dave = makeAddr("dave");

    YieldDistributorModule internal yieldModule;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    AccessControlFacet internal ac;

    // ERC-20 reward tokens
    MockERC20 internal usdc;   // 6 decimals
    MockERC20 internal weth;   // 18 decimals
    MockERC20 internal wbtc;   // 8 decimals
    MockERC20 internal dai;    // 18 decimals

    // External ERC-1155 reward token
    MockERC1155 internal extNft;

    uint256 internal TOKEN_1;

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    // Convenience: RewardAsset aliases
    function _erc20Asset(address t) internal pure returns (YieldDistributorModule.RewardAsset memory) {
        return YieldDistributorModule.RewardAsset({
            token: t,
            id: 0,
            assetType: YieldDistributorModule.RewardType.ERC20
        });
    }

    function _erc1155Asset(address t, uint256 id) internal pure returns (YieldDistributorModule.RewardAsset memory) {
        return YieldDistributorModule.RewardAsset({
            token: t,
            id: id,
            assetType: YieldDistributorModule.RewardType.ERC1155
        });
    }

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

        // Deploy reward tokens
        usdc = new MockERC20("USD Coin", "USDC", 6);
        weth = new MockERC20("Wrapped Ether", "WETH", 18);
        wbtc = new MockERC20("Wrapped Bitcoin", "WBTC", 8);
        dai = new MockERC20("Dai", "DAI", 18);
        extNft = new MockERC1155();

        // Deploy yield module
        yieldModule = new YieldDistributorModule(address(d.diamond), owner);

        // Register asset with yield module as plugin module
        address[] memory compModules = new address[](0);
        uint16[] memory countries = new uint16[](0);

        vm.startPrank(owner);
        TOKEN_1 = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModules: compModules,
                issuer: owner,
                allowedCountries: countries
            })
        );

        // Plug yield module into the asset
        am.addPluginModule(TOKEN_1, address(yieldModule));

        // Grant roles
        ac.grantRole(ISSUER_ROLE, owner);
        ac.grantRole(TRANSFER_AGENT, owner);

        // Register USDC as default reward asset
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(usdc)));

        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    function test_Constructor_SetsState() public view {
        assertEq(yieldModule.DIAMOND(), address(d.diamond));
        assertEq(yieldModule.owner(), owner);
    }

    function test_RevertWhen_Constructor_ZeroDiamond() public {
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__ZeroDiamond()"));
        new YieldDistributorModule(address(0), owner);
    }

    function test_Name() public view {
        assertEq(yieldModule.name(), "YieldDistributor");
    }

    /*//////////////////////////////////////////////////////////////
                    REWARD ASSET MANAGEMENT — ERC-20
    //////////////////////////////////////////////////////////////*/

    function test_AddRewardAsset_ERC20() public {
        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));

        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(weth)));
        assertTrue(yieldModule.isRewardAsset(TOKEN_1, key));

        YieldDistributorModule.RewardAsset[] memory assets = yieldModule.getRewardAssets(TOKEN_1);
        assertEq(assets.length, 2); // usdc + weth
        assertEq(assets[1].token, address(weth));
    }

    function test_RevertWhen_AddRewardAsset_NotOwner() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__OnlyOwner()"));
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));
    }

    function test_RevertWhen_AddRewardAsset_ZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__ZeroAddress()"));
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(0)));
    }

    function test_RevertWhen_AddRewardAsset_AlreadyAdded() public {
        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(usdc)));
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(
                YieldDistributorModule.YieldDistributorModule__RewardAssetAlreadyAdded.selector,
                TOKEN_1,
                key
            )
        );
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(usdc)));
    }

    function test_RevertWhen_AddRewardAsset_InvalidERC20Id() public {
        YieldDistributorModule.RewardAsset memory badAsset = YieldDistributorModule.RewardAsset({
            token: address(usdc),
            id: 42,
            assetType: YieldDistributorModule.RewardType.ERC20
        });
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__InvalidERC20Id()"));
        yieldModule.addRewardAsset(TOKEN_1, badAsset);
    }

    function test_RevertWhen_AddRewardAsset_TooMany() public {
        vm.startPrank(owner);
        // usdc already added in setUp (1)
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));  // 2
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(wbtc)));  // 3
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(dai)));   // 4
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(makeAddr("token5"))); // 5

        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__TooManyRewardAssets()"));
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(makeAddr("token6"))); // 6 → revert
        vm.stopPrank();
    }

    function test_RemoveRewardAsset_ERC20() public {
        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(usdc)));
        vm.prank(owner);
        yieldModule.removeRewardAsset(TOKEN_1, _erc20Asset(address(usdc)));

        assertFalse(yieldModule.isRewardAsset(TOKEN_1, key));
        YieldDistributorModule.RewardAsset[] memory assets = yieldModule.getRewardAssets(TOKEN_1);
        assertEq(assets.length, 0);
    }

    function test_RevertWhen_RemoveRewardAsset_NotFound() public {
        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(weth)));
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(
                YieldDistributorModule.YieldDistributorModule__RewardAssetNotFound.selector,
                TOKEN_1,
                key
            )
        );
        yieldModule.removeRewardAsset(TOKEN_1, _erc20Asset(address(weth)));
    }

    /*//////////////////////////////////////////////////////////////
                    REWARD ASSET MANAGEMENT — ERC-1155
    //////////////////////////////////////////////////////////////*/

    function test_AddRewardAsset_ERC1155External() public {
        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), 42));

        bytes32 key = yieldModule.rewardKey(_erc1155Asset(address(extNft), 42));
        assertTrue(yieldModule.isRewardAsset(TOKEN_1, key));
    }

    function test_AddRewardAsset_ERC1155_TokenIdZero() public {
        // ERC-1155 tokenId 0 IS valid (unlike ERC-20 where id must be 0)
        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), 0));

        bytes32 key = yieldModule.rewardKey(_erc1155Asset(address(extNft), 0));
        assertTrue(yieldModule.isRewardAsset(TOKEN_1, key));
    }

    function test_RemoveRewardAsset_ERC1155() public {
        vm.startPrank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), 42));
        yieldModule.removeRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), 42));
        vm.stopPrank();

        bytes32 key = yieldModule.rewardKey(_erc1155Asset(address(extNft), 42));
        assertFalse(yieldModule.isRewardAsset(TOKEN_1, key));
    }

    /*//////////////////////////////////////////////////////////////
                    DEPOSIT YIELD — ERC-20 (USDC)
    //////////////////////////////////////////////////////////////*/

    function test_DepositYield_USDC() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);

        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(usdc)));
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(usdc.balanceOf(address(yieldModule)), depositAmount);
        assertGt(yieldModule.accRewardPerShare(TOKEN_1, key), 0);
    }

    function test_RevertWhen_DepositYield_ZeroAmount() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__ZeroAmount()"));
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 0);
    }

    function test_RevertWhen_DepositYield_ZeroSupply() public {
        usdc.mint(owner, 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 1e6);

        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__ZeroSupply(uint256)", TOKEN_1));
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 1e6);
        vm.stopPrank();
    }

    function test_RevertWhen_DepositYield_UnregisteredReward() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(weth)));
        vm.prank(owner);
        vm.expectRevert(
            abi.encodeWithSelector(
                YieldDistributorModule.YieldDistributorModule__RewardAssetNotFound.selector,
                TOKEN_1,
                key
            )
        );
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(weth)), 1e18);
    }

    function test_RevertWhen_DepositYield_NotOwner() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__OnlyOwner()"));
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 1e6);
    }

    /*//////////////////////////////////////////////////////////////
                    DEPOSIT YIELD — ERC-1155 EXTERNAL
    //////////////////////////////////////////////////////////////*/

    function test_DepositYield_ERC1155External() public {
        uint256 rewardId = 42;

        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), rewardId));

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 500;
        extNft.mint(owner, rewardId, depositAmount);

        vm.startPrank(owner);
        extNft.setApprovalForAll(address(yieldModule), true);
        yieldModule.depositYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), depositAmount);
        vm.stopPrank();

        assertEq(extNft.balanceOf(address(yieldModule), rewardId), depositAmount);
    }

    /*//////////////////////////////////////////////////////////////
                    CLAIM — SINGLE HOLDER, USDC
    //////////////////////////////////////////////////////////////*/

    function test_ClaimYield_SingleHolder_USDC() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        uint256 claimable = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice);
        assertEq(claimable, depositAmount);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));

        assertEq(usdc.balanceOf(alice), depositAmount);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 0);
    }

    /*//////////////////////////////////////////////////////////////
                    CLAIM — ERC-1155 EXTERNAL
    //////////////////////////////////////////////////////////////*/

    function test_ClaimYield_ERC1155External() public {
        uint256 rewardId = 42;

        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), rewardId));

        // Alice holds staked tokens
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        // Deposit ERC-1155 reward
        uint256 depositAmount = 100;
        extNft.mint(owner, rewardId, depositAmount);
        vm.startPrank(owner);
        extNft.setApprovalForAll(address(yieldModule), true);
        yieldModule.depositYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), depositAmount);
        vm.stopPrank();

        // Alice claims
        uint256 claimable = yieldModule.claimableYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), alice);
        assertEq(claimable, depositAmount);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId));

        assertEq(extNft.balanceOf(alice, rewardId), depositAmount);
    }

    function test_ClaimYield_ERC1155_ProportionalTwoHolders() public {
        uint256 rewardId = 7;

        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), rewardId));

        // Alice: 7500, Bob: 2500
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 7500);
        supply.mint(TOKEN_1, bob, 2500);
        vm.stopPrank();

        // Deposit 1000 NFT units
        uint256 depositAmount = 1000;
        extNft.mint(owner, rewardId, depositAmount);
        vm.startPrank(owner);
        extNft.setApprovalForAll(address(yieldModule), true);
        yieldModule.depositYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), depositAmount);
        vm.stopPrank();

        // Alice: 75% = 750, Bob: 25% = 250
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), alice), 750);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), bob), 250);
    }

    /*//////////////////////////////////////////////////////////////
                    CLAIM — PROPORTIONAL DISTRIBUTION (ERC-20)
    //////////////////////////////////////////////////////////////*/

    function test_ClaimYield_Proportional_TwoHolders() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 7500);
        supply.mint(TOKEN_1, bob, 2500);
        vm.stopPrank();

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 750 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 250 * 1e6);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
        vm.prank(bob);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));

        assertEq(usdc.balanceOf(alice), 750 * 1e6);
        assertEq(usdc.balanceOf(bob), 250 * 1e6);
    }

    function test_ClaimYield_Proportional_ThreeHolders() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 5000);
        supply.mint(TOKEN_1, bob, 3000);
        supply.mint(TOKEN_1, carol, 2000);
        vm.stopPrank();

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 500 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 300 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), carol), 200 * 1e6);
    }

    /*//////////////////////////////////////////////////////////////
                    CLAIM — NOTHING TO CLAIM
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ClaimYield_NothingToClaim() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__NothingToClaim()"));
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
    }

    function test_RevertWhen_ClaimYield_AlreadyClaimed() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__NothingToClaim()"));
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
    }

    /*//////////////////////////////////////////////////////////////
                    CLAIM ALL — MIXED REWARD TYPES
    //////////////////////////////////////////////////////////////*/

    function test_ClaimAllYield_MixedERC20AndERC1155() public {
        uint256 rewardId = 10;

        // Add WETH and external ERC-1155 as additional rewards
        vm.startPrank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));
        yieldModule.addRewardAsset(TOKEN_1, _erc1155Asset(address(extNft), rewardId));
        vm.stopPrank();

        // Mint staked tokens
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        // Deposit USDC
        uint256 usdcAmount = 500 * 1e6;
        usdc.mint(owner, usdcAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), usdcAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), usdcAmount);

        // Deposit WETH
        uint256 wethAmount = 2 * 1e18;
        weth.mint(owner, wethAmount);
        weth.approve(address(yieldModule), wethAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(weth)), wethAmount);

        // Deposit external ERC-1155
        uint256 nftAmount = 50;
        extNft.mint(owner, rewardId, nftAmount);
        extNft.setApprovalForAll(address(yieldModule), true);
        yieldModule.depositYield(TOKEN_1, _erc1155Asset(address(extNft), rewardId), nftAmount);
        vm.stopPrank();

        // Claim all
        vm.prank(alice);
        yieldModule.claimAllYield(TOKEN_1);

        assertEq(usdc.balanceOf(alice), usdcAmount);
        assertEq(weth.balanceOf(alice), wethAmount);
        assertEq(extNft.balanceOf(alice, rewardId), nftAmount);
    }

    /*//////////////////////////////////////////////////////////////
                    HOOKS — TRANSFER CHECKPOINT
    //////////////////////////////////////////////////////////////*/

    function test_OnTransfer_CrystallizesRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 5000, "");

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 0);

        uint256 depositAmount2 = 500 * 1e6;
        usdc.mint(owner, depositAmount2);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount2);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount2);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 1250 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 250 * 1e6);
    }

    function test_OnTransfer_NewReceiverNoRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 10_000, "");

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 0);
    }

    /*//////////////////////////////////////////////////////////////
                    HOOKS — MINT CHECKPOINT
    //////////////////////////////////////////////////////////////*/

    function test_OnMint_ExistingHolder_PreservesRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 5000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 5000);

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
    }

    function test_OnMint_NewHolder_NoRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(owner);
        supply.mint(TOKEN_1, bob, 10_000);

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 0);
    }

    /*//////////////////////////////////////////////////////////////
                    HOOKS — BURN CHECKPOINT
    //////////////////////////////////////////////////////////////*/

    function test_OnBurn_PreservesRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 200 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(owner);
        supply.burn(TOKEN_1, alice, 5000);

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
        assertEq(usdc.balanceOf(alice), depositAmount);
    }

    /*//////////////////////////////////////////////////////////////
                    HOOKS — FORCED TRANSFER
    //////////////////////////////////////////////////////////////*/

    function test_OnForcedTransfer_CrystallizesRewards() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(owner);
        supply.forcedTransfer(TOKEN_1, alice, bob, 5000, bytes32(0));

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 0);
    }

    /*//////////////////////////////////////////////////////////////
                    MULTIPLE DEPOSITS
    //////////////////////////////////////////////////////////////*/

    function test_MultipleDeposits_AccumulateCorrectly() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        usdc.mint(owner, 100 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 100 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 100 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 100 * 1e6);

        usdc.mint(owner, 200 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 200 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 200 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 300 * 1e6);

        usdc.mint(owner, 50 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 50 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 50 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 350 * 1e6);
    }

    /*//////////////////////////////////////////////////////////////
                    DECIMAL PRECISION
    //////////////////////////////////////////////////////////////*/

    function test_Precision_WETH_18Decimals() public {
        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 7000);
        supply.mint(TOKEN_1, bob, 3000);
        vm.stopPrank();

        uint256 depositAmount = 10 * 1e18;
        weth.mint(owner, depositAmount);
        vm.startPrank(owner);
        weth.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(weth)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(weth)), alice), 7 * 1e18);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(weth)), bob), 3 * 1e18);
    }

    function test_Precision_WBTC_8Decimals() public {
        vm.prank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(wbtc)));

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 5000);
        supply.mint(TOKEN_1, bob, 5000);
        vm.stopPrank();

        uint256 depositAmount = 1 * 1e8;
        wbtc.mint(owner, depositAmount);
        vm.startPrank(owner);
        wbtc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(wbtc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(wbtc)), alice), 5000_0000);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(wbtc)), bob), 5000_0000);
    }

    function test_Precision_SmallDeposit_USDC() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 1_000_000);

        uint256 depositAmount = 10_000;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
    }

    function test_Precision_VerySmallDeposit_USDC() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 1;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), depositAmount);
    }

    function test_Precision_SmallHolder_LargeSupply() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 1);
        supply.mint(TOKEN_1, bob, 999_999);
        vm.stopPrank();

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        uint256 aliceClaimable = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice);
        assertEq(aliceClaimable, 1000);

        uint256 bobClaimable = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob);
        assertEq(bobClaimable, 999_999_000);

        assertLe(depositAmount - (aliceClaimable + bobClaimable), 1);
    }

    /*//////////////////////////////////////////////////////////////
                    NON-STANDARD ERC-20 (USDT-style)
    //////////////////////////////////////////////////////////////*/

    function test_DepositAndClaim_NonStandardERC20() public {
        MockNonStandardERC20 usdt = new MockNonStandardERC20("Tether", "USDT", 6);

        vm.startPrank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(usdt)));
        supply.mint(TOKEN_1, alice, 10_000);
        vm.stopPrank();

        uint256 depositAmount = 100 * 1e6;
        usdt.mint(owner, depositAmount);

        vm.startPrank(owner);
        usdt.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdt)), depositAmount);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdt)), alice), depositAmount);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdt)));

        assertEq(usdt.balanceOf(alice), depositAmount);
    }

    /*//////////////////////////////////////////////////////////////
                    FAILING ERC-20
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_DepositYield_TransferFails() public {
        MockFailingERC20 failToken = new MockFailingERC20();

        vm.startPrank(owner);
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(failToken)));
        supply.mint(TOKEN_1, alice, 10_000);
        vm.stopPrank();

        failToken.mint(owner, 1000);
        vm.startPrank(owner);
        failToken.approve(address(yieldModule), 1000);
        failToken.setFail(true);

        vm.expectRevert();
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(failToken)), 1000);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                    COMPLEX SCENARIO — MULTI-DEPOSIT + TRANSFER
    //////////////////////////////////////////////////////////////*/

    function test_ComplexScenario_DepositTransferDeposit() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        usdc.mint(owner, 100 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 100 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 100 * 1e6);
        vm.stopPrank();

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 5000, "");

        usdc.mint(owner, 200 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 200 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 200 * 1e6);
        vm.stopPrank();

        vm.prank(bob);
        token.safeTransferFrom(bob, carol, TOKEN_1, 2500, "");

        usdc.mint(owner, 300 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 300 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 300 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 350 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 175 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), carol), 75 * 1e6);

        uint256 total = 350 * 1e6 + 175 * 1e6 + 75 * 1e6;
        assertEq(total, 600 * 1e6);
    }

    function test_ComplexScenario_MintDilution() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        usdc.mint(owner, 100 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 100 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 100 * 1e6);
        vm.stopPrank();

        vm.prank(owner);
        supply.mint(TOKEN_1, bob, 10_000);

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 100 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 0);

        usdc.mint(owner, 200 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 200 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 200 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 200 * 1e6);
        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 100 * 1e6);
    }

    function test_ComplexScenario_BurnThenClaim() public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 10_000);
        supply.mint(TOKEN_1, bob, 10_000);
        vm.stopPrank();

        usdc.mint(owner, 200 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 200 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 200 * 1e6);
        vm.stopPrank();

        vm.prank(owner);
        supply.burn(TOKEN_1, bob, 10_000);

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob), 100 * 1e6);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
        vm.prank(bob);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));

        assertEq(usdc.balanceOf(alice), 100 * 1e6);
        assertEq(usdc.balanceOf(bob), 100 * 1e6);
    }

    function test_ComplexScenario_ClaimBetweenDeposits() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        usdc.mint(owner, 100 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 100 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 100 * 1e6);
        vm.stopPrank();

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
        assertEq(usdc.balanceOf(alice), 100 * 1e6);

        usdc.mint(owner, 200 * 1e6);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), 200 * 1e6);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), 200 * 1e6);
        vm.stopPrank();

        assertEq(yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice), 200 * 1e6);

        vm.prank(alice);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
        assertEq(usdc.balanceOf(alice), 300 * 1e6);
    }

    /*//////////////////////////////////////////////////////////////
                    HOOK ACCESS CONTROL
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_OnAction_Transfer_NotDiamond() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__OnlyDiamond()"));
        yieldModule.onAction(
            IHookablePlugin.ActionParams({
                actionType: IHookablePlugin.ActionType.Transfer,
                tokenId: TOKEN_1,
                operator: alice,
                from: alice,
                to: bob,
                amount: 100
            })
        );
    }

    function test_RevertWhen_OnAction_Mint_NotDiamond() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__OnlyDiamond()"));
        yieldModule.onAction(
            IHookablePlugin.ActionParams({
                actionType: IHookablePlugin.ActionType.Mint,
                tokenId: TOKEN_1,
                operator: alice,
                from: address(0),
                to: alice,
                amount: 100
            })
        );
    }

    function test_RevertWhen_OnAction_Burn_NotDiamond() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("YieldDistributorModule__OnlyDiamond()"));
        yieldModule.onAction(
            IHookablePlugin.ActionParams({
                actionType: IHookablePlugin.ActionType.Burn,
                tokenId: TOKEN_1,
                operator: alice,
                from: alice,
                to: address(0),
                amount: 100
            })
        );
    }

    /*//////////////////////////////////////////////////////////////
                    NO REWARD ASSETS — HOOKS ARE NO-OPS
    //////////////////////////////////////////////////////////////*/

    function test_Hooks_NoOp_WhenNoRewardAssets() public {
        vm.prank(owner);
        yieldModule.removeRewardAsset(TOKEN_1, _erc20Asset(address(usdc)));

        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 5000, "");

        vm.prank(owner);
        supply.burn(TOKEN_1, bob, 1000);
    }

    /*//////////////////////////////////////////////////////////////
                    IERC1155Receiver
    //////////////////////////////////////////////////////////////*/

    function test_SupportsInterface_ERC1155Receiver() public view {
        // IERC1155Receiver interfaceId = 0x4e2312e0
        assertTrue(yieldModule.supportsInterface(0x4e2312e0));
        // IERC165 interfaceId = 0x01ffc9a7
        assertTrue(yieldModule.supportsInterface(0x01ffc9a7));
        // Random interface
        assertFalse(yieldModule.supportsInterface(0xdeadbeef));
    }

    /*//////////////////////////////////////////////////////////////
                    FUZZ TESTS
    //////////////////////////////////////////////////////////////*/

    function testFuzz_ProportionalDistribution(uint256 aliceBalance, uint256 bobBalance, uint256 depositAmount) public {
        aliceBalance = bound(aliceBalance, 1, 1e12);
        bobBalance = bound(bobBalance, 1, 1e12);
        depositAmount = bound(depositAmount, 1, 1e12);

        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, aliceBalance);
        supply.mint(TOKEN_1, bob, bobBalance);
        vm.stopPrank();

        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        uint256 totalSupply_ = aliceBalance + bobBalance;
        uint256 aliceExpected = (depositAmount * aliceBalance) / totalSupply_;
        uint256 bobExpected = (depositAmount * bobBalance) / totalSupply_;

        uint256 aliceClaimable = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice);
        uint256 bobClaimable = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob);

        assertApproxEqAbs(aliceClaimable, aliceExpected, 1);
        assertApproxEqAbs(bobClaimable, bobExpected, 1);
        assertLe(aliceClaimable + bobClaimable, depositAmount);
    }

    function testFuzz_ClaimNeverExceedsDeposit(uint256 holderCount_, uint256 depositAmount) public {
        holderCount_ = bound(holderCount_, 1, 10);
        depositAmount = bound(depositAmount, 1, 1e15);

        address[] memory holders = new address[](holderCount_);
        for (uint256 i; i < holderCount_; ++i) {
            holders[i] = makeAddr(string(abi.encodePacked("holder", i)));
            uint256 balance = 1000 + i * 500;
            vm.prank(owner);
            supply.mint(TOKEN_1, holders[i], balance);
        }

        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        uint256 totalClaimable;
        for (uint256 i; i < holderCount_; ++i) {
            totalClaimable += yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), holders[i]);
        }

        assertLe(totalClaimable, depositAmount);
        assertGe(totalClaimable + holderCount_, depositAmount);
    }

    function testFuzz_TransferPreservesTotalClaimable(uint256 transferAmount) public {
        vm.startPrank(owner);
        supply.mint(TOKEN_1, alice, 10_000);
        supply.mint(TOKEN_1, bob, 10_000);
        vm.stopPrank();

        uint256 depositAmount = 1000 * 1e6;
        usdc.mint(owner, depositAmount);
        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        uint256 beforeTotal = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice)
            + yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob);

        transferAmount = bound(transferAmount, 1, 10_000);
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, transferAmount, "");

        uint256 afterTotal = yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), alice)
            + yieldModule.claimableYield(TOKEN_1, _erc20Asset(address(usdc)), bob);

        assertEq(afterTotal, beforeTotal);
    }

    /*//////////////////////////////////////////////////////////////
                    EVENTS
    //////////////////////////////////////////////////////////////*/

    function test_EmitEvent_YieldDeposited() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(usdc)));

        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);

        vm.expectEmit(true, true, false, true);
        emit YieldDistributorModule.YieldDeposited(TOKEN_1, key, depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();
    }

    function test_EmitEvent_YieldClaimed() public {
        vm.prank(owner);
        supply.mint(TOKEN_1, alice, 10_000);

        uint256 depositAmount = 100 * 1e6;
        usdc.mint(owner, depositAmount);
        bytes32 key = yieldModule.rewardKey(_erc20Asset(address(usdc)));

        vm.startPrank(owner);
        usdc.approve(address(yieldModule), depositAmount);
        yieldModule.depositYield(TOKEN_1, _erc20Asset(address(usdc)), depositAmount);
        vm.stopPrank();

        vm.prank(alice);
        vm.expectEmit(true, true, true, true);
        emit YieldDistributorModule.YieldClaimed(TOKEN_1, key, alice, depositAmount);
        yieldModule.claimYield(TOKEN_1, _erc20Asset(address(usdc)));
    }

    function test_EmitEvent_RewardAssetAdded() public {
        vm.prank(owner);
        vm.expectEmit(true, true, false, true);
        emit YieldDistributorModule.RewardAssetAdded(
            TOKEN_1,
            address(weth),
            0,
            YieldDistributorModule.RewardType.ERC20
        );
        yieldModule.addRewardAsset(TOKEN_1, _erc20Asset(address(weth)));
    }

    function test_EmitEvent_RewardAssetRemoved() public {
        vm.prank(owner);
        vm.expectEmit(true, true, false, true);
        emit YieldDistributorModule.RewardAssetRemoved(
            TOKEN_1,
            address(usdc),
            0,
            YieldDistributorModule.RewardType.ERC20
        );
        yieldModule.removeRewardAsset(TOKEN_1, _erc20Asset(address(usdc)));
    }
}
