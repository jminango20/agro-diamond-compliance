// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {AssetGroup} from "../../storage/LibAssetGroupStorage.sol";

interface IAssetGroup {
    struct CreateGroupParams {
        string name;               // group display name
        uint256 parentTokenId;     // must be already registered via registerAsset
        uint256 maxUnits;          // max child tokens (0 = unlimited)
    }

    struct MintUnitParams {
        uint256 groupId;
        string name;               // child asset name (e.g., "Apt 101")
        string symbol;             // child asset symbol
        string uri;                // metadata URI
        uint256 supplyCap;         // fractions cap (0 = unlimited)
        address investor;          // who receives the initial mint
        uint256 amount;            // initial fractionalized supply to mint
    }

    function createGroup(CreateGroupParams calldata params) external returns (uint256 groupId);
    function mintUnit(MintUnitParams calldata params) external returns (uint256 childTokenId);
    function mintUnitBatch(MintUnitParams[] calldata params) external returns (uint256[] memory childTokenIds);

    function getGroup(uint256 groupId) external view returns (AssetGroup memory);
    function getGroupChildren(uint256 groupId) external view returns (uint256[] memory);
    function getChildGroup(uint256 childTokenId) external view returns (uint256 groupId);
    function getRegisteredGroupIds() external view returns (uint256[] memory);
    function groupExists(uint256 groupId) external view returns (bool);
}
