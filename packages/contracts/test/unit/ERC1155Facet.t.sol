// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {MockComplianceModule} from "../helpers/MockComplianceModule.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {FreezeFacet} from "../../src/facets/rwa/FreezeFacet.sol";
import {PauseFacet} from "../../src/facets/security/PauseFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";
import {LibERC1155Storage, ERC1155Storage} from "../../src/storage/LibERC1155Storage.sol";
import {LibReasonCodes} from "../../src/libraries/LibReasonCodes.sol";

contract ERC1155FacetTest is DiamondHelper {
    DeployedDiamond internal d;
    address internal owner = makeAddr("owner");
    address internal alice = makeAddr("alice");
    address internal bob = makeAddr("bob");
    address internal operator = makeAddr("operator");

    ERC1155Facet internal token;
    AssetManagerFacet internal am;
    FreezeFacet internal freeze;
    PauseFacet internal pause;
    AccessControlFacet internal ac;
    MockComplianceModule internal mockModule;

    uint256 internal TOKEN_1;
    uint256 internal TOKEN_2;

    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");

    address[] internal emptyModules;

    function setUp() public {
        d = deployDiamond(owner);
        token = ERC1155Facet(address(d.diamond));
        am = AssetManagerFacet(address(d.diamond));
        freeze = FreezeFacet(address(d.diamond));
        pause = PauseFacet(address(d.diamond));
        ac = AccessControlFacet(address(d.diamond));
        mockModule = new MockComplianceModule();

        // Register two assets
        uint16[] memory countries = new uint16[](0);
        vm.startPrank(owner);
        TOKEN_1 = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: "Bond A",
                symbol: "BNDA",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModules: emptyModules,
                issuer: owner,
                allowedCountries: countries
            })
        );
        TOKEN_2 = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: "Bond B",
                symbol: "BNDB",
                uri: "",
                supplyCap: 0,
                identityProfileId: 0,
                complianceModules: emptyModules,
                issuer: owner,
                allowedCountries: countries
            })
        );
        ac.grantRole(TRANSFER_AGENT, owner);
        vm.stopPrank();

        // Give alice some balance by writing directly to storage (no SupplyFacet yet)
        _setFreeBalance(TOKEN_1, alice, 1000);
        _setFreeBalance(TOKEN_2, alice, 500);
    }

    /*//////////////////////////////////////////////////////////////
                        SAFE TRANSFER FROM
    //////////////////////////////////////////////////////////////*/

    function test_SafeTransferFrom() public {
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
        assertEq(token.balanceOf(alice, TOKEN_1), 900);
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
    }

    function test_SafeTransferFrom_EmitsEvents() public {
        vm.prank(alice);
        vm.expectEmit(true, true, true, true);
        emit ERC1155Facet.TransferSingle(alice, alice, bob, TOKEN_1, 100);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_SafeTransferFrom_ByOperator() public {
        vm.prank(alice);
        token.setApprovalForAll(operator, true);

        vm.prank(operator);
        token.safeTransferFrom(alice, bob, TOKEN_1, 50, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 50);
    }

    function test_RevertWhen_SafeTransferFrom_NotApproved() public {
        vm.prank(operator);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__NotApprovedOrOwner()"));
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_RevertWhen_SafeTransferFrom_ToZero() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__TransferToZeroAddress()"));
        token.safeTransferFrom(alice, address(0), TOKEN_1, 100, "");
    }

    function test_RevertWhen_SafeTransferFrom_InsufficientBalance() public {
        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSignature(
                "ERC1155Facet__InsufficientFreeBalance(uint256,address,uint256,uint256)", TOKEN_1, alice, 1000, 2000
            )
        );
        token.safeTransferFrom(alice, bob, TOKEN_1, 2000, "");
    }

    function test_RevertWhen_SafeTransferFrom_AssetNotRegistered() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__AssetNotRegistered(uint256)", 999));
        token.safeTransferFrom(alice, bob, 999, 100, "");
    }

    /*//////////////////////////////////////////////////////////////
                    SAFE BATCH TRANSFER FROM
    //////////////////////////////////////////////////////////////*/

    function test_SafeBatchTransferFrom() public {
        uint256[] memory ids = new uint256[](2);
        ids[0] = TOKEN_1;
        ids[1] = TOKEN_2;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 100;
        amounts[1] = 50;

        vm.prank(alice);
        token.safeBatchTransferFrom(alice, bob, ids, amounts, "");

        assertEq(token.balanceOf(alice, TOKEN_1), 900);
        assertEq(token.balanceOf(alice, TOKEN_2), 450);
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
        assertEq(token.balanceOf(bob, TOKEN_2), 50);
    }

    function test_RevertWhen_BatchTransfer_ArrayLengthMismatch() public {
        uint256[] memory ids = new uint256[](2);
        uint256[] memory amounts = new uint256[](1);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__ArrayLengthMismatch()"));
        token.safeBatchTransferFrom(alice, bob, ids, amounts, "");
    }

    /*//////////////////////////////////////////////////////////////
                        APPROVAL
    //////////////////////////////////////////////////////////////*/

    function test_SetApprovalForAll() public {
        vm.prank(alice);
        token.setApprovalForAll(operator, true);
        assertTrue(token.isApprovedForAll(alice, operator));

        vm.prank(alice);
        token.setApprovalForAll(operator, false);
        assertFalse(token.isApprovedForAll(alice, operator));
    }

    function test_RevertWhen_SelfApproval() public {
        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__SelfApproval()"));
        token.setApprovalForAll(alice, true);
    }

    /*//////////////////////////////////////////////////////////////
                        BALANCE VIEWS
    //////////////////////////////////////////////////////////////*/

    function test_BalanceOfBatch() public {
        address[] memory accounts = new address[](2);
        accounts[0] = alice;
        accounts[1] = alice;
        uint256[] memory ids = new uint256[](2);
        ids[0] = TOKEN_1;
        ids[1] = TOKEN_2;

        uint256[] memory bals = token.balanceOfBatch(accounts, ids);
        assertEq(bals[0], 1000);
        assertEq(bals[1], 500);
    }

    function test_PartitionBalanceOf() public {
        (uint256 free, uint256 locked, uint256 custody, uint256 pending) =
            token.partitionBalanceOf(alice, TOKEN_1);
        assertEq(free, 1000);
        assertEq(locked, 0);
        assertEq(custody, 0);
        assertEq(pending, 0);
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — GLOBAL PAUSE
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_ProtocolPaused() public {
        vm.prank(owner);
        pause.pauseProtocol();

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__ProtocolPaused()"));
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — ASSET PAUSE
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_AssetPaused() public {
        vm.prank(owner);
        pause.pauseAsset(TOKEN_1);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__AssetPaused(uint256)", TOKEN_1));
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — FREEZE
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_SenderGloballyFrozen() public {
        vm.prank(owner);
        freeze.setWalletFrozen(alice, true);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__WalletFrozenGlobal(address)", alice));
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_RevertWhen_ReceiverGloballyFrozen() public {
        vm.prank(owner);
        freeze.setWalletFrozen(bob, true);

        vm.prank(alice);
        vm.expectRevert(abi.encodeWithSignature("ERC1155Facet__WalletFrozenGlobal(address)", bob));
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_RevertWhen_SenderAssetFrozen() public {
        vm.prank(owner);
        freeze.setAssetWalletFrozen(TOKEN_1, alice, true);

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSignature("ERC1155Facet__WalletFrozenAsset(uint256,address)", TOKEN_1, alice)
        );
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — PARTIAL FREEZE
    //////////////////////////////////////////////////////////////*/

    function test_PartialFreeze_ReducesAvailable() public {
        vm.prank(owner);
        freeze.setFrozenAmount(TOKEN_1, alice, 800);

        // Available = 1000 - 800 = 200
        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 200, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 200);
    }

    function test_RevertWhen_PartialFreeze_ExceedsAvailable() public {
        vm.prank(owner);
        freeze.setFrozenAmount(TOKEN_1, alice, 800);

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSignature(
                "ERC1155Facet__InsufficientFreeBalance(uint256,address,uint256,uint256)", TOKEN_1, alice, 200, 300
            )
        );
        token.safeTransferFrom(alice, bob, TOKEN_1, 300, "");
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — LOCKUP
    //////////////////////////////////////////////////////////////*/

    function test_RevertWhen_LockupActive() public {
        uint64 futureExpiry = uint64(block.timestamp + 1 days);
        vm.prank(owner);
        freeze.setLockupExpiry(TOKEN_1, alice, futureExpiry);

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSignature("ERC1155Facet__LockupActive(uint256,address,uint64)", TOKEN_1, alice, futureExpiry)
        );
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    function test_LockupExpired_AllowsTransfer() public {
        uint64 pastExpiry = uint64(block.timestamp - 1);
        vm.prank(owner);
        freeze.setLockupExpiry(TOKEN_1, alice, pastExpiry);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
    }

    /*//////////////////////////////////////////////////////////////
                    REGULATORY — COMPLIANCE MODULE
    //////////////////////////////////////////////////////////////*/

    function test_ComplianceModule_Allows() public {
        vm.prank(owner);
        am.addComplianceModule(TOKEN_1, address(mockModule));

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
        assertEq(token.balanceOf(bob, TOKEN_1), 100);
        assertEq(mockModule.transferredCount(), 1);
    }

    function test_RevertWhen_ComplianceModule_Rejects() public {
        vm.prank(owner);
        am.addComplianceModule(TOKEN_1, address(mockModule));
        mockModule.setResult(false, LibReasonCodes.REASON_COUNTRY_RESTRICTED);

        vm.prank(alice);
        vm.expectRevert(
            abi.encodeWithSignature(
                "ERC1155Facet__ComplianceRejected(uint256,bytes32)", TOKEN_1, LibReasonCodes.REASON_COUNTRY_RESTRICTED
            )
        );
        token.safeTransferFrom(alice, bob, TOKEN_1, 100, "");
    }

    /*//////////////////////////////////////////////////////////////
                            FUZZ
    //////////////////////////////////////////////////////////////*/

    function testFuzz_TransferAmount(uint256 amount) public {
        vm.assume(amount > 0 && amount <= 1000);

        vm.prank(alice);
        token.safeTransferFrom(alice, bob, TOKEN_1, amount, "");
        assertEq(token.balanceOf(alice, TOKEN_1), 1000 - amount);
        assertEq(token.balanceOf(bob, TOKEN_1), amount);
    }

    /*//////////////////////////////////////////////////////////////
                        HELPERS — STORAGE WRITE
    //////////////////////////////////////////////////////////////*/

    /// @dev Directly writes free balance to ERC1155Storage (no SupplyFacet yet)
    function _setFreeBalance(uint256 tokenId, address account, uint256 amount) internal {
        // Use vm.store to set the free balance in the partition storage
        // ERC1155Storage.partitions[tokenId][account].free
        // Storage slot: keccak256(abi.encode(account, keccak256(abi.encode(tokenId, POSITION))))
        bytes32 position = LibERC1155Storage.POSITION;
        bytes32 partitionsSlot = keccak256(abi.encode(account, keccak256(abi.encode(tokenId, position))));
        // .free is the first field (offset 0)
        vm.store(address(d.diamond), partitionsSlot, bytes32(amount));
    }
}
