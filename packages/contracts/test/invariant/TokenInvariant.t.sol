// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DiamondHelper} from "../helpers/DiamondHelper.sol";
import {TokenHandler} from "./TokenHandler.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";
import {AssetManagerFacet} from "../../src/facets/token/AssetManagerFacet.sol";
import {AccessControlFacet} from "../../src/facets/security/AccessControlFacet.sol";
import {IAssetManager} from "../../src/interfaces/token/IAssetManager.sol";

/// @title TokenInvariant
/// @notice Invariant tests for supply conservation, holder tracking,
///         and balance consistency across mint/burn/transfer sequences.
contract TokenInvariant is DiamondHelper {
    DeployedDiamond internal d;
    TokenHandler internal handler;

    address internal owner = makeAddr("owner");
    uint256 internal constant TOKEN_1 = 1;
    uint256 internal constant SUPPLY_CAP = 1_000_000;
    uint256 internal constant ACTOR_COUNT = 5;

    SupplyFacet internal supply;
    ERC1155Facet internal token;

    address[] internal actors;

    function setUp() public {
        d = deployDiamond(owner);
        supply = SupplyFacet(address(d.diamond));
        token = ERC1155Facet(address(d.diamond));

        // Register asset
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

        // Grant TRANSFER_AGENT to owner for forcedTransfer
        bytes32 transferAgent = keccak256("TRANSFER_AGENT");
        AccessControlFacet(address(d.diamond)).grantRole(transferAgent, owner);
        vm.stopPrank();

        // Create actors
        for (uint256 i; i < ACTOR_COUNT; ++i) {
            actors.push(makeAddr(string(abi.encodePacked("actor", i))));
        }

        // Deploy handler
        handler = new TokenHandler(address(d.diamond), owner, TOKEN_1, actors);

        // Target only the handler
        targetContract(address(handler));
    }

    /*//////////////////////////////////////////////////////////////
                    INVARIANT: SUPPLY CONSERVATION
    //////////////////////////////////////////////////////////////*/

    /// @notice totalSupply == totalMinted - totalBurned
    function invariant_supplyConservation() public view {
        uint256 currentSupply = supply.totalSupply(TOKEN_1);
        uint256 expected = handler.ghost_totalMinted() - handler.ghost_totalBurned();
        assertEq(currentSupply, expected, "SUPPLY CONSERVATION VIOLATED");
    }

    /// @notice sum of all actor free balances == totalSupply
    function invariant_balanceSumEqualsTotalSupply() public view {
        uint256 sum;
        address[] memory a = handler.getActors();
        for (uint256 i; i < a.length; ++i) {
            sum += token.balanceOf(a[i], TOKEN_1);
        }
        assertEq(sum, supply.totalSupply(TOKEN_1), "BALANCE SUM != TOTAL SUPPLY");
    }

    /*//////////////////////////////////////////////////////////////
                    INVARIANT: SUPPLY CAP
    //////////////////////////////////////////////////////////////*/

    /// @notice totalSupply never exceeds supplyCap
    function invariant_supplyCapRespected() public view {
        assertLe(supply.totalSupply(TOKEN_1), SUPPLY_CAP, "SUPPLY CAP EXCEEDED");
    }

    /*//////////////////////////////////////////////////////////////
                    INVARIANT: HOLDER TRACKING
    //////////////////////////////////////////////////////////////*/

    /// @notice holderCount matches actual non-zero balance holders
    function invariant_holderCountAccuracy() public view {
        uint256 counted;
        address[] memory a = handler.getActors();
        for (uint256 i; i < a.length; ++i) {
            if (token.balanceOf(a[i], TOKEN_1) > 0) {
                counted += 1;
            }
        }
        assertEq(supply.holderCount(TOKEN_1), counted, "HOLDER COUNT MISMATCH");
    }

    /// @notice isHolder flag matches actual balance
    function invariant_isHolderConsistency() public view {
        address[] memory a = handler.getActors();
        for (uint256 i; i < a.length; ++i) {
            bool hasBalance = token.balanceOf(a[i], TOKEN_1) > 0;
            bool flagged = supply.isHolder(TOKEN_1, a[i]);
            assertEq(flagged, hasBalance, "IS_HOLDER FLAG MISMATCH");
        }
    }

    /*//////////////////////////////////////////////////////////////
                    INVARIANT: NO NEGATIVE BALANCES
    //////////////////////////////////////////////////////////////*/

    /// @notice no actor has more balance than totalSupply
    function invariant_noBalanceExceedsTotalSupply() public view {
        uint256 total = supply.totalSupply(TOKEN_1);
        address[] memory a = handler.getActors();
        for (uint256 i; i < a.length; ++i) {
            assertLe(token.balanceOf(a[i], TOKEN_1), total, "BALANCE > TOTAL SUPPLY");
        }
    }

    /*//////////////////////////////////////////////////////////////
                        CALL SUMMARY
    //////////////////////////////////////////////////////////////*/

    function invariant_callSummary() public {
        emit log_named_uint("totalMinted", handler.ghost_totalMinted());
        emit log_named_uint("totalBurned", handler.ghost_totalBurned());
        emit log_named_uint("transfers", handler.ghost_transferCount());
        emit log_named_uint("totalSupply", supply.totalSupply(TOKEN_1));
        emit log_named_uint("holderCount", supply.holderCount(TOKEN_1));
    }
}
