// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// solhint-disable no-inline-assembly

/// @dev A single balance snapshot for one holder at one point in time.
struct HolderSnapshot {
    uint256 balance;    // total balance (free + locked + custody + pending) at snapshot time
    bool recorded;      // true once balance has been captured (allows zero-balance snapshots)
}

/// @dev A snapshot record for a tokenId.
struct Snapshot {
    uint256 id;
    uint256 tokenId;
    uint256 totalSupply;                         // totalSupply at snapshot time
    uint64 timestamp;                            // block.timestamp when created
    uint256 holderCount;                         // holder count at snapshot time
    mapping(address => HolderSnapshot) balances; // holder → balance at snapshot
    bool exists;
}

struct SnapshotStorage {
    /// Global snapshot counter (auto-incremented)
    uint256 nextSnapshotId;
    /// snapshotId → Snapshot
    mapping(uint256 => Snapshot) snapshots;
    /// tokenId → array of snapshot IDs (ordered chronologically)
    mapping(uint256 => uint256[]) tokenSnapshots;
}

/// @title LibSnapshotStorage
/// @notice Namespaced snapshot storage for the Diamond.
///         slot = keccak256("diamond.rwa.snapshot.storage") - 1
library LibSnapshotStorage {
    bytes32 internal constant POSITION =
        0x9e3a4c79d4a58c3e3e29e56e8b6c76d3a6f8d5c2b1a0e9f8d7c6b5a4938271a0;

    function layout() internal pure returns (SnapshotStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}
