// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibDiamond, DiamondStorage} from "../libraries/LibDiamond.sol";
import {LibAppStorage, AppStorage} from "../libraries/LibAppStorage.sol";
import {IDiamondCut, IDiamondLoupe} from "../interfaces/core/IDiamond.sol";
// solhint-disable-next-line import-path-check
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {IERC1155} from "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import {IERC1155MetadataURI} from "@openzeppelin/contracts/token/ERC1155/extensions/IERC1155MetadataURI.sol";

/// @title DiamondInit
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Called via delegatecall during the first diamondCut to initialize
///         Diamond storage that cannot be set in the constructor.
/// @dev Must be called with _init = address(this) and _calldata = abi.encodeCall(DiamondInit.init, (...))
///      during the diamondCut that registers DiamondLoupeFacet.
///      Add new initialization logic here as facets are added — never re-deploy
///      existing initializers, always add new ones to avoid re-initialization.
contract DiamondInit {
    /// @notice Registers ERC-165 interfaces and sets contract-level metadata.
    /// @param name_ Contract name shown on block explorers (e.g. "Diamond RWA")
    /// @param symbol_ Contract symbol shown on block explorers (e.g. "dRWA")
    function init(string calldata name_, string calldata symbol_) external {
        DiamondStorage storage ds = LibDiamond.diamondStorage();
        ds.supportedInterfaces[type(IERC165).interfaceId] = true;
        ds.supportedInterfaces[type(IDiamondCut).interfaceId] = true;
        ds.supportedInterfaces[type(IDiamondLoupe).interfaceId] = true;
        ds.supportedInterfaces[type(IERC1155).interfaceId] = true; // 0xd9b67a26
        ds.supportedInterfaces[type(IERC1155MetadataURI).interfaceId] = true; // 0x0e89341c

        AppStorage storage app = LibAppStorage.layout();
        app.contractName = name_;
        app.contractSymbol = symbol_;
    }
}
