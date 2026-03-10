// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// solhint-disable no-inline-assembly

/// @dev Represents a group of related assets (e.g., a building with apartments).
///      The parent tokenId is registered first; child tokenIds are created
///      lazily via mintUnit() — only consuming gas when actually sold.
struct AssetGroup {
    uint256 parentTokenId;       // root asset (e.g., "Edifício Aurora")
    string name;                 // group name
    uint256 maxUnits;            // max child tokens allowed (0 = unlimited)
    uint256 unitCount;           // current number of minted child tokens
    uint256 nextUnitIndex;       // auto-increment for deterministic child tokenId
    bool exists;
}

struct AssetGroupStorage {
    mapping(uint256 => AssetGroup) groups;          // groupId → group
    mapping(uint256 => uint256[]) groupChildren;    // groupId → child tokenIds
    mapping(uint256 => uint256) childToGroup;       // childTokenId → groupId (0 = not a child)
    uint256[] registeredGroupIds;                   // all group IDs
    uint256 nextGroupId;                            // auto-increment
}

/// @title LibAssetGroupStorage
/// @notice Namespaced storage for asset groups (parent → children hierarchy).
///         slot = keccak256("diamond.rwa.assetgroup.storage") - 1
library LibAssetGroupStorage {
    bytes32 internal constant POSITION =
        0xec9a1d83e9fa81fab3c79de9268ae94d914c93201618209b8d8075722aa4e3dd;

    function layout() internal pure returns (AssetGroupStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}
