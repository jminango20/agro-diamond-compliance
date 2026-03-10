// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {
    AssetGroupFacet,
    AssetGroupFacet__Unauthorized,
    AssetGroupFacet__ParentNotRegistered,
    AssetGroupFacet__GroupNotFound,
    AssetGroupFacet__MaxUnitsReached,
    AssetGroupFacet__EmptyName,
    AssetGroupFacet__MintToZeroAddress,
    AssetGroupFacet__ReceiverFrozen,
    AssetGroupFacet__EmptyBatch
} from "../../src/facets/rwa/AssetGroupFacet.sol";
import {IAssetGroup} from "../../src/interfaces/token/IAssetGroup.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {AssetGroup} from "../../src/storage/LibAssetGroupStorage.sol";
import {AssetConfig} from "../../src/storage/LibAssetStorage.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {FreezeFacet} from "../../src/facets/rwa/FreezeFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";

contract AssetGroupFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal investor1 = makeAddr("investor1");
    address internal investor2 = makeAddr("investor2");
    address internal stranger = makeAddr("stranger");

    uint256 internal constant PARENT_TOKEN_ID = 1;
    uint256 internal constant PARENT_SUPPLY_CAP = 1_000_000;

    AssetGroupFacet internal groupFacet;
    AssetManagerFacet internal assetManager;
    SupplyFacet internal supply;
    ERC1155Facet internal erc1155;

    function setUp() public {
        d = deployDiamond(owner);
        groupFacet = AssetGroupFacet(address(d.diamond));
        assetManager = AssetManagerFacet(address(d.diamond));
        supply = SupplyFacet(address(d.diamond));
        erc1155 = ERC1155Facet(address(d.diamond));

        // Register parent asset
        uint16[] memory countries = new uint16[](0);
        vm.prank(owner);
        assetManager.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: PARENT_TOKEN_ID,
                name: "Edificio Aurora",
                symbol: "AURORA",
                uri: "ipfs://aurora",
                supplyCap: PARENT_SUPPLY_CAP,
                identityProfileId: 0,
                complianceModule: address(0),
                issuer: owner,
                allowedCountries: countries
            })
        );
    }

    /*//////////////////////////////////////////////////////////////
                            CREATE GROUP
    //////////////////////////////////////////////////////////////*/

    function test_CreateGroup() public {
        vm.prank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Edificio Aurora Apts", parentTokenId: PARENT_TOKEN_ID, maxUnits: 50})
        );

        assertEq(groupId, 1);
        assertTrue(groupFacet.groupExists(groupId));

        AssetGroup memory g = groupFacet.getGroup(groupId);
        assertEq(g.parentTokenId, PARENT_TOKEN_ID);
        assertEq(g.maxUnits, 50);
        assertEq(g.unitCount, 0);
        assertEq(g.nextUnitIndex, 1);
    }

    function test_CreateGroup_AutoIncrementId() public {
        vm.startPrank(owner);
        uint256 g1 = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Group A", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );
        uint256 g2 = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Group B", parentTokenId: PARENT_TOKEN_ID, maxUnits: 20})
        );
        vm.stopPrank();

        assertEq(g1, 1);
        assertEq(g2, 2);

        uint256[] memory ids = groupFacet.getRegisteredGroupIds();
        assertEq(ids.length, 2);
        assertEq(ids[0], 1);
        assertEq(ids[1], 2);
    }

    function test_CreateGroup_UnlimitedUnits() public {
        vm.prank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Unlimited", parentTokenId: PARENT_TOKEN_ID, maxUnits: 0})
        );

        AssetGroup memory g = groupFacet.getGroup(groupId);
        assertEq(g.maxUnits, 0);
    }

    function test_RevertWhen_CreateGroup_ParentNotRegistered() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AssetGroupFacet__ParentNotRegistered.selector, 999));
        groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Bad Group", parentTokenId: 999, maxUnits: 10})
        );
    }

    function test_RevertWhen_CreateGroup_EmptyName() public {
        vm.prank(owner);
        vm.expectRevert(AssetGroupFacet__EmptyName.selector);
        groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );
    }

    function test_RevertWhen_CreateGroup_Unauthorized() public {
        vm.prank(stranger);
        vm.expectRevert(AssetGroupFacet__Unauthorized.selector);
        groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Fail", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );
    }

    /*//////////////////////////////////////////////////////////////
                            MINT UNIT (LAZY)
    //////////////////////////////////////////////////////////////*/

    function test_MintUnit() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 100})
        );

        uint256 childTokenId = groupFacet.mintUnit(
            IAssetGroup.MintUnitParams({
                groupId: groupId,
                name: "Apt 101",
                symbol: "APT101",
                uri: "ipfs://apt101",
                supplyCap: 1000,
                investor: investor1,
                amount: 500
            })
        );
        vm.stopPrank();

        // Verify deterministic tokenId: (groupId << 128) | unitIndex
        uint256 expectedTokenId = (groupId << 128) | 1;
        assertEq(childTokenId, expectedTokenId);

        // Verify child asset was registered
        AssetConfig memory cfg = assetManager.getAssetConfig(childTokenId);
        assertTrue(cfg.exists);
        assertEq(cfg.name, "Apt 101");
        assertEq(cfg.symbol, "APT101");
        assertEq(cfg.supplyCap, 1000);
        assertEq(cfg.totalSupply, 500);
        assertEq(cfg.issuer, owner); // inherited from parent

        // Verify balance
        assertEq(erc1155.balanceOf(investor1, childTokenId), 500);
        assertTrue(supply.isHolder(childTokenId, investor1));
        assertEq(supply.holderCount(childTokenId), 1);

        // Verify group state
        AssetGroup memory g = groupFacet.getGroup(groupId);
        assertEq(g.unitCount, 1);
        assertEq(g.nextUnitIndex, 2);

        // Verify child→group mapping
        assertEq(groupFacet.getChildGroup(childTokenId), groupId);

        // Verify group children
        uint256[] memory children = groupFacet.getGroupChildren(groupId);
        assertEq(children.length, 1);
        assertEq(children[0], childTokenId);
    }

    function test_MintUnit_MultipleChildren() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 100})
        );

        uint256 child1 = groupFacet.mintUnit(
            IAssetGroup.MintUnitParams({
                groupId: groupId,
                name: "Apt 101",
                symbol: "A101",
                uri: "ipfs://101",
                supplyCap: 0,
                investor: investor1,
                amount: 100
            })
        );

        uint256 child2 = groupFacet.mintUnit(
            IAssetGroup.MintUnitParams({
                groupId: groupId,
                name: "Apt 102",
                symbol: "A102",
                uri: "ipfs://102",
                supplyCap: 0,
                investor: investor2,
                amount: 200
            })
        );
        vm.stopPrank();

        assertEq(child1, (groupId << 128) | 1);
        assertEq(child2, (groupId << 128) | 2);

        uint256[] memory children = groupFacet.getGroupChildren(groupId);
        assertEq(children.length, 2);

        AssetGroup memory g = groupFacet.getGroup(groupId);
        assertEq(g.unitCount, 2);
    }

    function test_MintUnit_ZeroAmount() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );

        uint256 childTokenId = groupFacet.mintUnit(
            IAssetGroup.MintUnitParams({
                groupId: groupId,
                name: "Apt Reserved",
                symbol: "RSRV",
                uri: "ipfs://reserved",
                supplyCap: 1000,
                investor: investor1,
                amount: 0
            })
        );
        vm.stopPrank();

        // Asset registered but no tokens minted
        assertTrue(assetManager.assetExists(childTokenId));
        assertEq(erc1155.balanceOf(investor1, childTokenId), 0);
        assertEq(supply.totalSupply(childTokenId), 0);
    }

    function test_MintUnit_InheritsParentConfig() public {
        // Register parent with specific compliance settings
        uint16[] memory countries = new uint16[](2);
        countries[0] = 76;  // Brazil
        countries[1] = 840; // US

        uint256 parentId = 42;
        vm.prank(owner);
        assetManager.registerAsset(
            IAssetManager.RegisterAssetParams({
                tokenId: parentId,
                name: "Premium Building",
                symbol: "PREM",
                uri: "ipfs://premium",
                supplyCap: 0,
                identityProfileId: 5,
                complianceModule: address(0xBEEF),
                issuer: address(0xCAFE),
                allowedCountries: countries
            })
        );

        vm.prank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Premium Apts", parentTokenId: parentId, maxUnits: 0})
        );

        vm.prank(owner);
        // amount=0 to avoid compliance module call on mock address
        uint256 childId = groupFacet.mintUnit(
            IAssetGroup.MintUnitParams({
                groupId: groupId,
                name: "Suite 1",
                symbol: "S1",
                uri: "ipfs://suite1",
                supplyCap: 100,
                investor: investor1,
                amount: 0
            })
        );

        AssetConfig memory childCfg = assetManager.getAssetConfig(childId);
        assertEq(childCfg.identityProfileId, 5);
        assertEq(childCfg.complianceModule, address(0xBEEF));
        assertEq(childCfg.issuer, address(0xCAFE));
        assertEq(childCfg.allowedCountries.length, 2);
        assertEq(childCfg.allowedCountries[0], 76);
        assertEq(childCfg.allowedCountries[1], 840);
    }

    function test_RevertWhen_MintUnit_MaxUnitsReached() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Small Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 2})
        );

        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt 1", investor1, 100));
        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt 2", investor2, 100));

        vm.expectRevert(abi.encodeWithSelector(AssetGroupFacet__MaxUnitsReached.selector, groupId, 2));
        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt 3", investor1, 100));
        vm.stopPrank();
    }

    function test_RevertWhen_MintUnit_GroupNotFound() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(AssetGroupFacet__GroupNotFound.selector, 999));
        groupFacet.mintUnit(_defaultMintParams(999, "Apt", investor1, 100));
    }

    function test_RevertWhen_MintUnit_ZeroAddressInvestor() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );

        vm.expectRevert(AssetGroupFacet__MintToZeroAddress.selector);
        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt", address(0), 100));
        vm.stopPrank();
    }

    function test_RevertWhen_MintUnit_ReceiverFrozen() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );

        // Freeze investor globally
        FreezeFacet(address(d.diamond)).setWalletFrozen(investor1, true);

        vm.expectRevert(abi.encodeWithSelector(AssetGroupFacet__ReceiverFrozen.selector, investor1));
        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt", investor1, 100));
        vm.stopPrank();
    }

    function test_RevertWhen_MintUnit_Unauthorized() public {
        vm.prank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );

        vm.prank(stranger);
        vm.expectRevert(AssetGroupFacet__Unauthorized.selector);
        groupFacet.mintUnit(_defaultMintParams(groupId, "Apt", investor1, 100));
    }

    /*//////////////////////////////////////////////////////////////
                            MINT UNIT BATCH
    //////////////////////////////////////////////////////////////*/

    function test_MintUnitBatch() public {
        vm.startPrank(owner);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 100})
        );

        IAssetGroup.MintUnitParams[] memory batch = new IAssetGroup.MintUnitParams[](3);
        batch[0] = IAssetGroup.MintUnitParams({
            groupId: groupId,
            name: "Apt 101",
            symbol: "A101",
            uri: "ipfs://101",
            supplyCap: 0,
            investor: investor1,
            amount: 100
        });
        batch[1] = IAssetGroup.MintUnitParams({
            groupId: groupId,
            name: "Apt 102",
            symbol: "A102",
            uri: "ipfs://102",
            supplyCap: 0,
            investor: investor2,
            amount: 200
        });
        batch[2] = IAssetGroup.MintUnitParams({
            groupId: groupId,
            name: "Apt 103",
            symbol: "A103",
            uri: "ipfs://103",
            supplyCap: 0,
            investor: investor1,
            amount: 300
        });

        uint256[] memory childIds = groupFacet.mintUnitBatch(batch);
        vm.stopPrank();

        assertEq(childIds.length, 3);
        assertEq(erc1155.balanceOf(investor1, childIds[0]), 100);
        assertEq(erc1155.balanceOf(investor2, childIds[1]), 200);
        assertEq(erc1155.balanceOf(investor1, childIds[2]), 300);

        AssetGroup memory g = groupFacet.getGroup(groupId);
        assertEq(g.unitCount, 3);
    }

    function test_RevertWhen_MintUnitBatch_Empty() public {
        vm.prank(owner);
        IAssetGroup.MintUnitParams[] memory batch = new IAssetGroup.MintUnitParams[](0);
        vm.expectRevert(AssetGroupFacet__EmptyBatch.selector);
        groupFacet.mintUnitBatch(batch);
    }

    /*//////////////////////////////////////////////////////////////
                        COMPLIANCE_ADMIN ACCESS
    //////////////////////////////////////////////////////////////*/

    function test_ComplianceAdmin_CanCreateGroup() public {
        address admin = makeAddr("complianceAdmin");
        vm.prank(owner);
        AccessControlFacet(address(d.diamond)).grantRole(keccak256("COMPLIANCE_ADMIN"), admin);

        vm.prank(admin);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Admin Group", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );
        assertTrue(groupFacet.groupExists(groupId));
    }

    function test_IssuerRole_CanMintUnit() public {
        address issuerAddr = makeAddr("issuer");
        vm.startPrank(owner);
        AccessControlFacet(address(d.diamond)).grantRole(keccak256("ISSUER_ROLE"), issuerAddr);
        uint256 groupId = groupFacet.createGroup(
            IAssetGroup.CreateGroupParams({name: "Building", parentTokenId: PARENT_TOKEN_ID, maxUnits: 10})
        );
        vm.stopPrank();

        vm.prank(issuerAddr);
        uint256 childId = groupFacet.mintUnit(_defaultMintParams(groupId, "Apt", investor1, 100));

        assertEq(erc1155.balanceOf(investor1, childId), 100);
    }

    /*//////////////////////////////////////////////////////////////
                            VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    function test_GetChildGroup_ReturnsZero_ForNonChild() public view {
        assertEq(groupFacet.getChildGroup(999), 0);
    }

    function test_GroupExists_ReturnsFalse_ForNonExistent() public view {
        assertFalse(groupFacet.groupExists(999));
    }

    function test_GetRegisteredGroupIds_EmptyInitially() public view {
        uint256[] memory ids = groupFacet.getRegisteredGroupIds();
        assertEq(ids.length, 0);
    }

    /*//////////////////////////////////////////////////////////////
                            HELPERS
    //////////////////////////////////////////////////////////////*/

    function _defaultMintParams(uint256 groupId, string memory name, address inv, uint256 amt)
        internal
        pure
        returns (IAssetGroup.MintUnitParams memory)
    {
        return IAssetGroup.MintUnitParams({
            groupId: groupId,
            name: name,
            symbol: "UNIT",
            uri: "ipfs://unit",
            supplyCap: 0,
            investor: inv,
            amount: amt
        });
    }
}
