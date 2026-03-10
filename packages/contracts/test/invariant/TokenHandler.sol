// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {SupplyFacet} from "../../src/facets/token/SupplyFacet.sol";
import {ERC1155Facet} from "../../src/facets/token/ERC1155Facet.sol";

/// @title TokenHandler
/// @notice Foundry invariant handler that exposes bounded supply and transfer
///         operations. The fuzzer calls these functions with random args;
///         the invariant test asserts global properties after each call sequence.
contract TokenHandler is Test {
    SupplyFacet internal supply;
    ERC1155Facet internal token;
    address internal diamond;
    address internal owner;
    uint256 internal tokenId;

    address[] internal actors;
    mapping(address => bool) internal isActor;

    uint256 public ghost_totalMinted;
    uint256 public ghost_totalBurned;
    uint256 public ghost_transferCount;

    constructor(address _diamond, address _owner, uint256 _tokenId, address[] memory _actors) {
        diamond = _diamond;
        supply = SupplyFacet(_diamond);
        token = ERC1155Facet(_diamond);
        owner = _owner;
        tokenId = _tokenId;

        for (uint256 i; i < _actors.length; ++i) {
            actors.push(_actors[i]);
            isActor[_actors[i]] = true;
        }
    }

    /*//////////////////////////////////////////////////////////////
                            BOUNDED ACTIONS
    //////////////////////////////////////////////////////////////*/

    function handler_mint(uint256 actorSeed, uint256 amount) external {
        address to = _pickActor(actorSeed);
        amount = bound(amount, 1, 1000);

        vm.prank(owner);
        try supply.mint(tokenId, to, amount) {
            ghost_totalMinted += amount;
        } catch {}
    }

    function handler_burn(uint256 actorSeed, uint256 amount) external {
        address from = _pickActor(actorSeed);
        uint256 bal = token.balanceOf(from, tokenId);
        if (bal == 0) return;
        amount = bound(amount, 1, bal);

        vm.prank(owner);
        try supply.burn(tokenId, from, amount) {
            ghost_totalBurned += amount;
        } catch {}
    }

    function handler_transfer(uint256 fromSeed, uint256 toSeed, uint256 amount) external {
        address from = _pickActor(fromSeed);
        address to = _pickActor(toSeed);
        if (from == to) return;

        uint256 bal = token.balanceOf(from, tokenId);
        if (bal == 0) return;
        amount = bound(amount, 1, bal);

        vm.prank(from);
        try token.safeTransferFrom(from, to, tokenId, amount, "") {
            ghost_transferCount += 1;
        } catch {}
    }

    function handler_forcedTransfer(uint256 fromSeed, uint256 toSeed, uint256 amount) external {
        address from = _pickActor(fromSeed);
        address to = _pickActor(toSeed);
        if (from == to) return;

        uint256 bal = token.balanceOf(from, tokenId);
        if (bal == 0) return;
        amount = bound(amount, 1, bal);

        vm.prank(owner);
        try supply.forcedTransfer(tokenId, from, to, amount, bytes32("INVARIANT_TEST")) {
            ghost_transferCount += 1;
        } catch {}
    }

    /*//////////////////////////////////////////////////////////////
                            VIEW HELPERS
    //////////////////////////////////////////////////////////////*/

    function getActors() external view returns (address[] memory) {
        return actors;
    }

    function actorCount() external view returns (uint256) {
        return actors.length;
    }

    /*//////////////////////////////////////////////////////////////
                            INTERNAL
    //////////////////////////////////////////////////////////////*/

    function _pickActor(uint256 seed) internal view returns (address) {
        return actors[seed % actors.length];
    }
}
