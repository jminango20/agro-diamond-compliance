// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

// NOTE: Excluded from solhint — assembly and ordering are required by Diamond pattern.
// See .solhintignore.

/*//////////////////////////////////////////////////////////////
    DESIGN: Namespaced Storage per Domain
    ────────────────────────────────────────────────────────────
    Rather than one giant AppStorage mixing all concerns, each
    domain has its own lib + deterministic slot:

    LibAppStorage      — global protocol state (owner, globalPaused)
    LibERC1155Storage  — balances, partitions, operator approvals
    LibAssetStorage    — per-tokenId config (name, profile, module)
    LibIdentityStorage — wallet→identity, profiles, verification cache
    LibComplianceStorage — tokenId→module, module registry
    LibFreezeStorage   — freeze flags, partial freeze, lockup expiry
    LibAccessStorage   — role-based access control
    LibSupplyStorage   — totalSupply, supplyCap per tokenId

    Slots computed as:
        bytes32(uint256(keccak256("<namespace>")) - 1)
    following EIP-1967 convention to avoid storage collisions.
//////////////////////////////////////////////////////////////*/

/*//////////////////////////////////////////////////////////////
                    GLOBAL APP STORAGE
    slot: keccak256("diamond.rwa.app.storage") - 1
//////////////////////////////////////////////////////////////*/

/// @dev Intentionally minimal — only truly global state lives here.
///      All domain-specific state lives in LibXxxStorage.
struct AppStorage {
    address contractOwner;
    address pendingOwner;
    bool globalPaused;
    uint16 protocolVersion;
}

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

/*//////////////////////////////////////////////////////////////
                    ERC-1155 STORAGE
    slot: keccak256("diamond.rwa.erc1155.storage") - 1
//////////////////////////////////////////////////////////////*/

/// @dev Sub-balance partitions per holder per tokenId.
///      Allows granular RWA-specific freeze/lock semantics beyond
///      a simple binary frozen flag.
struct PartitionBalance {
    uint256 free;              // transferable
    uint256 locked;            // frozen/restricted
    uint256 custody;           // held by custodian
    uint256 pendingSettlement; // DVP pending
}

struct ERC1155Storage {
    /// tokenId → holder → partition balances
    mapping(uint256 => mapping(address => PartitionBalance)) partitions;
    /// owner → operator → approved
    mapping(address => mapping(address => bool)) operatorApprovals;
}

library LibERC1155Storage {
    bytes32 internal constant POSITION =
        0xfa7964d34d14df8c2bc86a420d87e49e27dbfab86f789b31f9dd2373beee6837;

    function layout() internal pure returns (ERC1155Storage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    ASSET STORAGE
    slot: keccak256("diamond.rwa.asset.storage") - 1
//////////////////////////////////////////////////////////////*/

/// @dev Per-tokenId regulatory configuration.
///      identityProfileId points to a reusable policy in LibIdentityStorage
///      so multiple tokenIds can share the same KYC/AML requirements.
struct AssetConfig {
    string name;
    string symbol;
    string uri;
    uint256 supplyCap;         // 0 = no cap
    uint32 identityProfileId;  // → LibIdentityStorage.profiles[id]
    address complianceModule;  // → IComplianceModule
    address issuer;            // authorized minter for this asset
    bool paused;
    bool exists;
    uint16[] allowedCountries; // ISO 3166-1 numeric; empty = all allowed
}

struct AssetStorage {
    mapping(uint256 => AssetConfig) configs;
    uint256[] registeredTokenIds;
}

library LibAssetStorage {
    bytes32 internal constant POSITION =
        0x715f14664c3dc6471d9524868696b43c245e0aba6f0c77ed9e2dea380113fffb;

    function layout() internal pure returns (AssetStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    IDENTITY STORAGE
    slot: keccak256("diamond.rwa.identity.storage") - 1
//////////////////////////////////////////////////////////////*/

/// @dev Regulatory profile: defines WHICH claims are required and
///      WHICH issuers are trusted to attest them.
///      Multiple tokenIds can share one profile (e.g. "accredited investor").
struct IdentityProfile {
    uint256[] requiredClaimTopics;
    mapping(address => bool) trustedIssuers;
    uint64 version; // incremented on change → invalidates verification cache
}

struct IdentityStorage {
    /// wallet → ONCHAINID contract address
    mapping(address => address) walletToIdentity;
    /// wallet → ISO 3166-1 country code
    mapping(address => uint16) walletCountry;
    /// wallet → identity version (incremented on identity update)
    mapping(address => uint64) identityVersion;
    /// profileId → regulatory profile
    mapping(uint32 => IdentityProfile) profiles;
    /// wallet → profileId → cached verification result
    mapping(address => mapping(uint32 => bool)) verifiedCache;
    /// wallet → profileId → version at which cache was computed
    ///   cache is valid if identityVersion[wallet] == cacheVersion &&
    ///                     profiles[profileId].version == profileVersion
    mapping(address => mapping(uint32 => uint64)) cacheIdentityVersion;
    mapping(address => mapping(uint32 => uint64)) cacheProfileVersion;
    uint32 profileCount;
}

library LibIdentityStorage {
    bytes32 internal constant POSITION =
        0xb301aa8172310cb299b51011e8e8be41f21916ff640c952e834ef2501a47d375;

    function layout() internal pure returns (IdentityStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    COMPLIANCE STORAGE
    slot: keccak256("diamond.rwa.compliance.storage") - 1
//////////////////////////////////////////////////////////////*/

struct ComplianceStorage {
    /// tokenId → compliance module address
    mapping(uint256 => address) tokenModule;
    /// registered module addresses
    mapping(address => bool) registeredModules;
}

library LibComplianceStorage {
    bytes32 internal constant POSITION =
        0x6b502ce75824dd1c89ed4f44589fce210e1cbc92ddffac6021b7a14f73382a47;

    function layout() internal pure returns (ComplianceStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    FREEZE STORAGE
    slot: keccak256("diamond.rwa.freeze.storage") - 1
//////////////////////////////////////////////////////////////*/

struct FreezeStorage {
    /// wallet → frozen across ALL assets
    mapping(address => bool) globalFreeze;
    /// tokenId → wallet → frozen for THIS asset
    mapping(uint256 => mapping(address => bool)) assetFreeze;
    /// tokenId → wallet → amount frozen within free partition
    mapping(uint256 => mapping(address => uint256)) frozenAmount;
    /// tokenId → wallet → lockup expiry timestamp (0 = no lockup)
    mapping(uint256 => mapping(address => uint64)) lockupExpiry;
}

library LibFreezeStorage {
    bytes32 internal constant POSITION =
        0x816ddc8ce7b7a63b00e3aba2be8d7cda4caf59ac1130e1bde05d22df1dd742a4;

    function layout() internal pure returns (FreezeStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    ACCESS CONTROL STORAGE
    slot: keccak256("diamond.rwa.access.storage") - 1
//////////////////////////////////////////////////////////////*/

struct AccessStorage {
    /// role → account → has role
    mapping(bytes32 => mapping(address => bool)) roles;
    /// role → admin role (who can grant/revoke)
    mapping(bytes32 => bytes32) roleAdmin;
}

library LibAccessStorage {
    bytes32 internal constant POSITION =
        0xd243a54eb63068fbcd3b75bed7e8048bd76bf590c2627f5dd4b1909b9ae4cb5a;

    function layout() internal pure returns (AccessStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}

/*//////////////////////////////////////////////////////////////
                    SUPPLY STORAGE
    slot: keccak256("diamond.rwa.supply.storage") - 1
//////////////////////////////////////////////////////////////*/

struct SupplyStorage {
    /// tokenId → total minted supply
    mapping(uint256 => uint256) totalSupply;
    /// tokenId → unique holder count
    mapping(uint256 => uint256) holderCount;
    /// tokenId → holder → has any balance (for holderCount tracking)
    mapping(uint256 => mapping(address => bool)) isHolder;
}

library LibSupplyStorage {
    bytes32 internal constant POSITION =
        0x46830504cc643be957af372bc50edfbfbf33b72cb2c5f2919048fc26e5d225bb;

    function layout() internal pure returns (SupplyStorage storage s) {
        bytes32 position = POSITION;
        assembly {
            s.slot := position
        }
    }
}
