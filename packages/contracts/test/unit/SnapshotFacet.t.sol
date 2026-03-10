// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {SnapshotFacet} from "../../src/facets/rwa/SnapshotFacet.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

contract SnapshotFacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal charlie = makeAddr("charlie");
    address internal issuer = makeAddr("issuer");
    address internal attacker = makeAddr("attacker");

    SnapshotFacet internal snapshot;
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    AccessControlFacet internal ac;

    uint256 internal constant TOKEN_1 = 1;
    uint256 internal constant TOKEN_2 = 2;

    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    function setUp() public {
        d = deployDiamond(owner);
        snapshot = SnapshotFacet(address(d.diamond));
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));

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
        supply.mint(TOKEN_1, alice, 3000);
        supply.mint(TOKEN_1, bob, 7000);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                        CREATE SNAPSHOT
    //////////////////////////////////////////////////////////////*/

    function test_CreateSnapshot() public {
        vm.prank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);

        (uint256 tokenId, uint256 totalSupply, uint64 timestamp, uint256 holderCount) = snapshot.getSnapshot(id);
        assertEq(tokenId, TOKEN_1);
        assertEq(totalSupply, 10_000);
        assertGt(timestamp, 0);
        assertEq(holderCount, 2);
    }

    function test_CreateSnapshot_IssuerCanCreate() public {
        vm.prank(issuer);
        uint256 id = snapshot.createSnapshot(TOKEN_1);
        assertEq(id, 0);
    }

    function test_CreateSnapshot_IncrementsId() public {
        vm.startPrank(owner);
        uint256 id0 = snapshot.createSnapshot(TOKEN_1);
        uint256 id1 = snapshot.createSnapshot(TOKEN_1);
        vm.stopPrank();

        assertEq(id0, 0);
        assertEq(id1, 1);
    }

    function test_RevertWhen_CreateSnapshot_Unauthorized() public {
        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__Unauthorized()"));
        snapshot.createSnapshot(TOKEN_1);
    }

    function test_RevertWhen_CreateSnapshot_AssetNotRegistered() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__AssetNotRegistered(uint256)", 999));
        snapshot.createSnapshot(999);
    }

    /*//////////////////////////////////////////////////////////////
                        RECORD HOLDER
    //////////////////////////////////////////////////////////////*/

    function test_RecordHolder() public {
        vm.startPrank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);
        snapshot.recordHolder(id, alice);
        vm.stopPrank();

        (uint256 balance, bool recorded) = snapshot.getSnapshotBalance(id, alice);
        assertEq(balance, 3000);
        assertTrue(recorded);
    }

    function test_RecordHolder_ZeroBalanceHolder() public {
        vm.startPrank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);
        snapshot.recordHolder(id, charlie);
        vm.stopPrank();

        (uint256 balance, bool recorded) = snapshot.getSnapshotBalance(id, charlie);
        assertEq(balance, 0);
        assertTrue(recorded);
    }

    function test_RevertWhen_RecordHolder_SnapshotNotFound() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__SnapshotNotFound(uint256)", 999));
        snapshot.recordHolder(999, alice);
    }

    function test_RevertWhen_RecordHolder_AlreadyRecorded() public {
        vm.startPrank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);
        snapshot.recordHolder(id, alice);
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__HolderAlreadyRecorded(uint256,address)", id, alice));
        snapshot.recordHolder(id, alice);
        vm.stopPrank();
    }

    function test_RevertWhen_RecordHolder_Unauthorized() public {
        vm.prank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);

        vm.prank(attacker);
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__Unauthorized()"));
        snapshot.recordHolder(id, alice);
    }

    /*//////////////////////////////////////////////////////////////
                        RECORD HOLDERS BATCH
    //////////////////////////////////////////////////////////////*/

    function test_RecordHoldersBatch() public {
        vm.startPrank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);
        address[] memory holders = new address[](2);
        holders[0] = alice;
        holders[1] = bob;
        snapshot.recordHoldersBatch(id, holders);
        vm.stopPrank();

        (uint256 aliceBal, bool aliceRec) = snapshot.getSnapshotBalance(id, alice);
        (uint256 bobBal, bool bobRec) = snapshot.getSnapshotBalance(id, bob);
        assertEq(aliceBal, 3000);
        assertTrue(aliceRec);
        assertEq(bobBal, 7000);
        assertTrue(bobRec);
    }

    /*//////////////////////////////////////////////////////////////
                            VIEWS
    //////////////////////////////////////////////////////////////*/

    function test_GetTokenSnapshots() public {
        vm.startPrank(owner);
        snapshot.createSnapshot(TOKEN_1);
        snapshot.createSnapshot(TOKEN_1);
        vm.stopPrank();

        uint256[] memory ids = snapshot.getTokenSnapshots(TOKEN_1);
        assertEq(ids.length, 2);
        assertEq(ids[0], 0);
        assertEq(ids[1], 1);
    }

    function test_GetTokenSnapshots_EmptyForUnusedToken() public view {
        uint256[] memory ids = snapshot.getTokenSnapshots(TOKEN_2);
        assertEq(ids.length, 0);
    }

    function test_GetLatestSnapshotId() public {
        vm.startPrank(owner);
        snapshot.createSnapshot(TOKEN_1);
        snapshot.createSnapshot(TOKEN_1);
        vm.stopPrank();

        uint256 latest = snapshot.getLatestSnapshotId(TOKEN_1);
        assertEq(latest, 1);
    }

    function test_RevertWhen_GetLatestSnapshotId_NoSnapshots() public {
        vm.expectRevert(abi.encodeWithSignature("SnapshotFacet__SnapshotNotFound(uint256)", 0));
        snapshot.getLatestSnapshotId(TOKEN_2);
    }

    function test_NextSnapshotId() public {
        assertEq(snapshot.nextSnapshotId(), 0);

        vm.prank(owner);
        snapshot.createSnapshot(TOKEN_1);

        assertEq(snapshot.nextSnapshotId(), 1);
    }

    function test_GetSnapshotBalance_UnrecordedHolder() public {
        vm.prank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);

        (uint256 balance, bool recorded) = snapshot.getSnapshotBalance(id, alice);
        assertEq(balance, 0);
        assertFalse(recorded);
    }

    /*//////////////////////////////////////////////////////////////
                        SNAPSHOT CAPTURES POINT-IN-TIME
    //////////////////////////////////////////////////////////////*/

    function test_SnapshotCapturesBalanceAtRecordTime() public {
        vm.prank(owner);
        uint256 id = snapshot.createSnapshot(TOKEN_1);

        // Transfer after snapshot creation but before recording
        vm.prank(alice);
        token.safeTransferFrom(alice, charlie, TOKEN_1, 1000, "");

        // Record after transfer — should see post-transfer balance
        vm.prank(owner);
        snapshot.recordHolder(id, alice);

        (uint256 balance,) = snapshot.getSnapshotBalance(id, alice);
        assertEq(balance, 2000); // 3000 - 1000
    }
}
