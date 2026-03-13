// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// solhint-disable no-inline-assembly

/// @dev Global protocol state — intentionally minimal.
///      All domain-specific state lives in src/storage/LibXxxStorage.sol.
struct AppStorage {
    address contractOwner;
    address pendingOwner;
    bool globalPaused;
    uint16 protocolVersion;
    string contractName;
    string contractSymbol;
}

/// @title LibAppStorage
/// @notice Global namespaced storage for the Diamond protocol.
///         slot = keccak256("diamond.rwa.app.storage") - 1
library LibAppStorage {
    bytes32 internal constant POSITION =
        0x16bcb91a8fc307d752c7c6050be37c62191cf5d127ed81a8a62d3854c599d3fb;

    function layout() internal pure returns (AppStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}
