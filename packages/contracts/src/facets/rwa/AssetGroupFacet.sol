// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibDiamond} from "../../libraries/LibDiamond.sol";
import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibAssetGroupStorage, AssetGroupStorage, AssetGroup} from "../../storage/LibAssetGroupStorage.sol";
import {LibERC1155Storage, ERC1155Storage} from "../../storage/LibERC1155Storage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";
import {LibAccessStorage, AccessStorage} from "../../storage/LibAccessStorage.sol";
import {LibAppStorage, AppStorage} from "../../libraries/LibAppStorage.sol";
import {LibFreezeStorage, FreezeStorage} from "../../storage/LibFreezeStorage.sol";
import {IAssetGroup} from "../../interfaces/token/IAssetGroup.sol";
import {IComplianceModule} from "../../interfaces/compliance/IComplianceModule.sol";

/*//////////////////////////////////////////////////////////////
                            ERRORS
//////////////////////////////////////////////////////////////*/

error AssetGroupFacet__Unauthorized();
error AssetGroupFacet__ParentNotRegistered(uint256 parentTokenId);
error AssetGroupFacet__GroupNotFound(uint256 groupId);
error AssetGroupFacet__MaxUnitsReached(uint256 groupId, uint256 maxUnits);
error AssetGroupFacet__ChildTokenIdCollision(uint256 childTokenId);
error AssetGroupFacet__EmptyName();
error AssetGroupFacet__ProtocolPaused();
error AssetGroupFacet__MintToZeroAddress();
error AssetGroupFacet__ReceiverFrozen(address account);
error AssetGroupFacet__EmptyBatch();

/*//////////////////////////////////////////////////////////////
                            CONTRACT
//////////////////////////////////////////////////////////////*/

/**
 * @title AssetGroupFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice Manages hierarchical asset groups with lazy minting for gas-efficient
 *         multi-token RWA creation (e.g., building → apartments).
 *
 *         Design:
 *         - createGroup(): registers a parent→children relationship (1 cheap tx)
 *         - mintUnit(): registers child asset + mints initial supply in 1 tx (lazy)
 *         - Child tokenId = keccak256(groupId, unitIndex) % 2^128 + groupId << 128
 *           This deterministic scheme prevents collisions between groups.
 *
 *         Gas optimization: child assets are only created on-chain when actually
 *         sold/minted, not upfront. A building with 100 apartments pays gas for
 *         1 createGroup tx + N mintUnit txs (only for apartments that are sold).
 *
 * @dev Caller: Diamond owner or COMPLIANCE_ADMIN.
 *      Child assets inherit the parent's identityProfileId and complianceModule
 *      by default, but can be overridden after creation via AssetManagerFacet.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract AssetGroupFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    event GroupCreated(uint256 indexed groupId, uint256 indexed parentTokenId, string name, uint256 maxUnits);
    event UnitMinted(
        uint256 indexed groupId, uint256 indexed childTokenId, address indexed investor, string name, uint256 amount
    );

    /*//////////////////////////////////////////////////////////////
                            ROLE CONSTANTS
    //////////////////////////////////////////////////////////////*/

    bytes32 internal constant COMPLIANCE_ADMIN = keccak256("COMPLIANCE_ADMIN");
    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");

    /*//////////////////////////////////////////////////////////////
                    EXTERNAL STATE-CHANGING FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Creates a new asset group linking a parent tokenId to future children.
    ///         The parent tokenId must already be registered via AssetManagerFacet.
    /// @param params Group creation parameters
    /// @return groupId The auto-generated group identifier
    function createGroup(IAssetGroup.CreateGroupParams calldata params) external returns (uint256 groupId) {
        _enforceComplianceAdminOrOwner();
        if (bytes(params.name).length == 0) revert AssetGroupFacet__EmptyName();

        AssetStorage storage as_ = LibAssetStorage.layout();
        if (!as_.configs[params.parentTokenId].exists) {
            revert AssetGroupFacet__ParentNotRegistered(params.parentTokenId);
        }

        AssetGroupStorage storage gs = LibAssetGroupStorage.layout();
        groupId = ++gs.nextGroupId;

        AssetGroup storage g = gs.groups[groupId];
        g.parentTokenId = params.parentTokenId;
        g.name = params.name;
        g.maxUnits = params.maxUnits;
        g.unitCount = 0;
        g.nextUnitIndex = 1;
        g.exists = true;

        gs.registeredGroupIds.push(groupId);

        emit GroupCreated(groupId, params.parentTokenId, params.name, params.maxUnits);
    }

    /// @notice Lazily creates a child asset and mints initial supply in a single tx.
    ///         The child inherits the parent's identity profile and compliance module.
    /// @param params Mint unit parameters (group, name, symbol, uri, cap, investor, amount)
    /// @return childTokenId The deterministic tokenId assigned to the new unit
    function mintUnit(IAssetGroup.MintUnitParams calldata params) external returns (uint256 childTokenId) {
        _enforceIssuerOrComplianceAdminOrOwner();
        childTokenId = _mintUnit(params);
    }

    /// @notice Batch version of mintUnit for creating multiple units in one tx.
    /// @param params Array of mint unit parameters
    /// @return childTokenIds Array of created child tokenIds
    function mintUnitBatch(IAssetGroup.MintUnitParams[] calldata params)
        external
        returns (uint256[] memory childTokenIds)
    {
        _enforceIssuerOrComplianceAdminOrOwner();
        if (params.length == 0) revert AssetGroupFacet__EmptyBatch();

        childTokenIds = new uint256[](params.length);
        for (uint256 i; i < params.length; ++i) {
            childTokenIds[i] = _mintUnit(params[i]);
        }
    }

    /*//////////////////////////////////////////////////////////////
                            VIEW FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the group metadata.
    function getGroup(uint256 groupId) external view returns (AssetGroup memory) {
        return LibAssetGroupStorage.layout().groups[groupId];
    }

    /// @notice Returns all child tokenIds for a group.
    function getGroupChildren(uint256 groupId) external view returns (uint256[] memory) {
        return LibAssetGroupStorage.layout().groupChildren[groupId];
    }

    /// @notice Returns the groupId for a child tokenId (0 = not a child).
    function getChildGroup(uint256 childTokenId) external view returns (uint256) {
        return LibAssetGroupStorage.layout().childToGroup[childTokenId];
    }

    /// @notice Returns all registered group IDs.
    function getRegisteredGroupIds() external view returns (uint256[] memory) {
        return LibAssetGroupStorage.layout().registeredGroupIds;
    }

    /// @notice Returns true if the group exists.
    function groupExists(uint256 groupId) external view returns (bool) {
        return LibAssetGroupStorage.layout().groups[groupId].exists;
    }

    /*//////////////////////////////////////////////////////////////
                        INTERNAL FUNCTIONS
    //////////////////////////////////////////////////////////////*/

    /// @dev Core logic for lazy minting a child unit.
    ///      1. Validates group exists, max units not exceeded
    ///      2. Generates deterministic childTokenId
    ///      3. Registers child as a new asset (inherits parent config)
    ///      4. Mints initial supply to investor
    function _mintUnit(IAssetGroup.MintUnitParams calldata params) internal returns (uint256 childTokenId) {
        _enforceNotPaused();

        AssetGroupStorage storage gs = LibAssetGroupStorage.layout();
        AssetGroup storage g = gs.groups[params.groupId];
        if (!g.exists) revert AssetGroupFacet__GroupNotFound(params.groupId);
        if (g.maxUnits != 0 && g.unitCount >= g.maxUnits) {
            revert AssetGroupFacet__MaxUnitsReached(params.groupId, g.maxUnits);
        }
        if (params.investor == address(0)) revert AssetGroupFacet__MintToZeroAddress();

        // Receiver freeze check
        FreezeStorage storage fs = LibFreezeStorage.layout();
        if (fs.globalFreeze[params.investor]) revert AssetGroupFacet__ReceiverFrozen(params.investor);

        // Generate deterministic childTokenId: high 128 bits = groupId, low 128 bits = unitIndex
        uint256 unitIndex = g.nextUnitIndex;
        childTokenId = (params.groupId << 128) | unitIndex;

        // Ensure no collision with existing asset
        AssetStorage storage as_ = LibAssetStorage.layout();
        if (as_.configs[childTokenId].exists) revert AssetGroupFacet__ChildTokenIdCollision(childTokenId);

        // Inherit parent config
        AssetConfig storage parentCfg = as_.configs[g.parentTokenId];

        // Register child asset
        AssetConfig storage childCfg = as_.configs[childTokenId];
        childCfg.name = params.name;
        childCfg.symbol = params.symbol;
        childCfg.uri = params.uri;
        childCfg.supplyCap = params.supplyCap;
        childCfg.totalSupply = 0;
        childCfg.identityProfileId = parentCfg.identityProfileId;
        childCfg.complianceModules = parentCfg.complianceModules;
        childCfg.issuer = parentCfg.issuer;
        childCfg.paused = false;
        childCfg.exists = true;
        childCfg.allowedCountries = parentCfg.allowedCountries;

        as_.registeredTokenIds.push(childTokenId);

        // Update group state
        g.unitCount += 1;
        g.nextUnitIndex = unitIndex + 1;
        gs.groupChildren[params.groupId].push(childTokenId);
        gs.childToGroup[childTokenId] = params.groupId;

        // Mint initial supply if amount > 0
        if (params.amount > 0) {
            _executeMint(childTokenId, childCfg, params.investor, params.amount);
        }

        emit UnitMinted(params.groupId, childTokenId, params.investor, params.name, params.amount);
    }

    /// @dev Mints tokens to investor — mirrors SupplyFacet._executeMint logic
    function _executeMint(uint256 tokenId, AssetConfig storage config, address to, uint256 amount) internal {
        // Supply cap check
        if (config.supplyCap != 0 && config.totalSupply + amount > config.supplyCap) {
            revert AssetGroupFacet__MaxUnitsReached(0, config.supplyCap);
        }

        // Update balances
        ERC1155Storage storage es = LibERC1155Storage.layout();
        es.partitions[tokenId][to].free += amount;

        // Update supply
        config.totalSupply += amount;

        // Holder tracking
        SupplyStorage storage ss = LibSupplyStorage.layout();
        if (!ss.isHolder[tokenId][to]) {
            ss.isHolder[tokenId][to] = true;
            ss.holderCount[tokenId] += 1;
        }

        // Compliance post-hooks
        address[] storage modules = config.complianceModules;
        uint256 mLen = modules.length;
        for (uint256 i; i < mLen;) {
            IComplianceModule(modules[i]).minted(tokenId, to, amount);
            unchecked { ++i; }
        }
    }

    function _enforceNotPaused() internal view {
        AppStorage storage app = LibAppStorage.layout();
        if (app.globalPaused) revert AssetGroupFacet__ProtocolPaused();
    }

    function _enforceComplianceAdminOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        bool isAdmin = LibAccessStorage.layout().roles[COMPLIANCE_ADMIN][msg.sender];
        if (!isOwner && !isAdmin) revert AssetGroupFacet__Unauthorized();
    }

    function _enforceIssuerOrComplianceAdminOrOwner() internal view {
        bool isOwner = msg.sender == LibDiamond.contractOwner();
        AccessStorage storage acl = LibAccessStorage.layout();
        bool isAdmin = acl.roles[COMPLIANCE_ADMIN][msg.sender];
        bool isIssuer = acl.roles[ISSUER_ROLE][msg.sender];
        if (!isOwner && !isAdmin && !isIssuer) revert AssetGroupFacet__Unauthorized();
    }
}
