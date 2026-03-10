// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibDiamond} from "../../libraries/LibDiamond.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibAccessStorage} from "../../storage/LibAccessStorage.sol";
import {IAssetManager} from "../../interfaces/token/IAssetManager.sol";

/*//////////////////////////////////////////////////////////////
                            ERRORS
//////////////////////////////////////////////////////////////*/

error AssetManagerFacet__NotRegistered(uint256 tokenId);
error AssetManagerFacet__ZeroAddress();
error AssetManagerFacet__EmptyString();
error AssetManagerFacet__Unauthorized();
error AssetManagerFacet__TooManyModules(uint256 count, uint256 max);
error AssetManagerFacet__ModuleNotFound(address module);
error AssetManagerFacet__ModuleAlreadyAdded(address module);

/*//////////////////////////////////////////////////////////////
                            CONTRACT
//////////////////////////////////////////////////////////////*/

/// @title AssetManagerFacet
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Manages per-tokenId asset configuration (architecture §3 — Level 2).
///         Each tokenId is an independent asset class with its own regulatory config:
///         identity profile, compliance modules, issuer, supply cap, and jurisdictions.
///         TokenIds are auto-incremented — callers never choose the ID.
///         Caller: Diamond owner or COMPLIANCE_ADMIN for config changes.
/// @custom:security-contact renan.correa@hubweb3.com
contract AssetManagerFacet {
    /*//////////////////////////////////////////////////////////////
                                CONSTANTS
    //////////////////////////////////////////////////////////////*/

    /// @dev Maximum compliance modules per tokenId to prevent unbounded loops in transfers.
    uint256 internal constant MAX_COMPLIANCE_MODULES = 10;

    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event AssetRegistered(uint256 indexed tokenId, address indexed issuer, uint32 profileId);
    event AssetConfigUpdated(uint256 indexed tokenId);
    event ComplianceModulesSet(uint256 indexed tokenId, address[] modules);
    event ComplianceModuleAdded(uint256 indexed tokenId, address indexed module);
    event ComplianceModuleRemoved(uint256 indexed tokenId, address indexed module);

    /*//////////////////////////////////////////////////////////////
                        STATE-CHANGING FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Registers a new asset class with auto-incremented tokenId.
    /// @param p Asset registration parameters (name, symbol, modules, issuer, etc.)
    /// @return tokenId The auto-generated token identifier
    function registerAsset(IAssetManager.RegisterAssetParams calldata p)
        external
        returns (uint256 tokenId)
    {
        _enforceComplianceAdminOrOwner();
        if (p.issuer == address(0)) revert AssetManagerFacet__ZeroAddress();
        if (bytes(p.name).length == 0) revert AssetManagerFacet__EmptyString();
        if (bytes(p.symbol).length == 0) revert AssetManagerFacet__EmptyString();
        if (p.complianceModules.length > MAX_COMPLIANCE_MODULES) {
            revert AssetManagerFacet__TooManyModules(p.complianceModules.length, MAX_COMPLIANCE_MODULES);
        }

        AssetStorage storage s = LibAssetStorage.layout();
        tokenId = ++s.nextTokenId;

        AssetConfig storage cfg = s.configs[tokenId];
        cfg.name = p.name;
        cfg.symbol = p.symbol;
        cfg.uri = p.uri;
        cfg.supplyCap = p.supplyCap;
        cfg.identityProfileId = p.identityProfileId;
        cfg.complianceModules = p.complianceModules;
        cfg.issuer = p.issuer;
        cfg.exists = true;
        cfg.allowedCountries = p.allowedCountries;

        s.registeredTokenIds.push(tokenId);

        emit AssetRegistered(tokenId, p.issuer, p.identityProfileId);
        if (p.complianceModules.length > 0) {
            emit ComplianceModulesSet(tokenId, p.complianceModules);
        }
    }

    /// @notice Adds a compliance module to a tokenId's module list.
    /// @param tokenId The asset class to modify
    /// @param module The compliance module address to add
    function addComplianceModule(uint256 tokenId, address module) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        if (module == address(0)) revert AssetManagerFacet__ZeroAddress();

        address[] storage modules = LibAssetStorage.layout().configs[tokenId].complianceModules;

        uint256 len = modules.length;
        for (uint256 i; i < len;) {
            if (modules[i] == module) revert AssetManagerFacet__ModuleAlreadyAdded(module);
            unchecked { ++i; }
        }
        if (len >= MAX_COMPLIANCE_MODULES) {
            revert AssetManagerFacet__TooManyModules(len + 1, MAX_COMPLIANCE_MODULES);
        }

        modules.push(module);
        emit ComplianceModuleAdded(tokenId, module);
    }

    /// @notice Removes a compliance module from a tokenId's module list.
    ///         Uses swap-and-pop for gas efficiency (order not guaranteed).
    /// @param tokenId The asset class to modify
    /// @param module The compliance module address to remove
    function removeComplianceModule(uint256 tokenId, address module) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);

        address[] storage modules = LibAssetStorage.layout().configs[tokenId].complianceModules;
        uint256 len = modules.length;

        for (uint256 i; i < len;) {
            if (modules[i] == module) {
                modules[i] = modules[len - 1];
                modules.pop();
                emit ComplianceModuleRemoved(tokenId, module);
                return;
            }
            unchecked { ++i; }
        }
        revert AssetManagerFacet__ModuleNotFound(module);
    }

    /// @notice Replaces all compliance modules for a tokenId.
    ///         Pass empty array to remove all modules (unrestricted transfers).
    /// @param tokenId The asset class to modify
    /// @param modules The new compliance module addresses
    function setComplianceModules(uint256 tokenId, address[] calldata modules) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        if (modules.length > MAX_COMPLIANCE_MODULES) {
            revert AssetManagerFacet__TooManyModules(modules.length, MAX_COMPLIANCE_MODULES);
        }
        LibAssetStorage.layout().configs[tokenId].complianceModules = modules;
        emit ComplianceModulesSet(tokenId, modules);
    }

    /// @notice Updates the identity profile used for KYC verification on this asset.
    function setIdentityProfile(uint256 tokenId, uint32 profileId) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        LibAssetStorage.layout().configs[tokenId].identityProfileId = profileId;
        emit AssetConfigUpdated(tokenId);
    }

    /// @notice Updates the authorized minter (issuer) for a tokenId.
    ///         Only Diamond owner can change issuer — higher privilege than COMPLIANCE_ADMIN.
    function setIssuer(uint256 tokenId, address issuer) external {
        LibDiamond.enforceIsContractOwner();
        _requireRegistered(tokenId);
        if (issuer == address(0)) revert AssetManagerFacet__ZeroAddress();
        LibAssetStorage.layout().configs[tokenId].issuer = issuer;
        emit AssetConfigUpdated(tokenId);
    }

    /// @notice Updates the supply cap. Pass 0 for unlimited.
    function setSupplyCap(uint256 tokenId, uint256 cap) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        LibAssetStorage.layout().configs[tokenId].supplyCap = cap;
        emit AssetConfigUpdated(tokenId);
    }

    /// @notice Replaces the allowed jurisdictions for a tokenId.
    ///         Pass empty array to allow all countries.
    function setAllowedCountries(uint256 tokenId, uint16[] calldata countries) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        LibAssetStorage.layout().configs[tokenId].allowedCountries = countries;
        emit AssetConfigUpdated(tokenId);
    }

    /// @notice Updates the metadata URI for a tokenId.
    function setAssetUri(uint256 tokenId, string calldata uri) external {
        _enforceComplianceAdminOrOwner();
        _requireRegistered(tokenId);
        LibAssetStorage.layout().configs[tokenId].uri = uri;
        emit AssetConfigUpdated(tokenId);
    }

    /*//////////////////////////////////////////////////////////////
                            VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the full AssetConfig for a tokenId.
    function getAssetConfig(uint256 tokenId) external view returns (AssetConfig memory) {
        return LibAssetStorage.layout().configs[tokenId];
    }

    /// @notice Returns the compliance modules for a tokenId.
    function getComplianceModules(uint256 tokenId) external view returns (address[] memory) {
        return LibAssetStorage.layout().configs[tokenId].complianceModules;
    }

    /// @notice Returns all registered tokenIds in registration order.
    function getRegisteredTokenIds() external view returns (uint256[] memory) {
        return LibAssetStorage.layout().registeredTokenIds;
    }

    /// @notice Returns true if the tokenId has been registered.
    function assetExists(uint256 tokenId) external view returns (bool) {
        return LibAssetStorage.layout().configs[tokenId].exists;
    }

    /// @notice Returns the current auto-increment counter for tokenIds.
    function nextTokenId() external view returns (uint256) {
        return LibAssetStorage.layout().nextTokenId;
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant COMPLIANCE_ADMIN = keccak256("COMPLIANCE_ADMIN");

    function _enforceComplianceAdminOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isAdmin = LibAccessStorage.layout().roles[COMPLIANCE_ADMIN][msg.sender];
        if (!isOwner && !isAdmin) revert AssetManagerFacet__Unauthorized();
    }

    function _requireRegistered(uint256 tokenId) internal view {
        if (!LibAssetStorage.layout().configs[tokenId].exists) {
            revert AssetManagerFacet__NotRegistered(tokenId);
        }
    }
}
