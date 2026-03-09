// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {LibAssetStorage, AssetStorage, AssetConfig} from "../../storage/LibAssetStorage.sol";
import {LibSupplyStorage, SupplyStorage} from "../../storage/LibSupplyStorage.sol";

/**
 * @title MetadataFacet
 * @author Renan Correa <renan.correa@hubweb3.com>
 * @notice ERC-1155 metadata views and per-tokenId asset information.
 *         Exposes `uri(tokenId)` per the ERC-1155 standard, plus
 *         `name(tokenId)` and `symbol(tokenId)` for RWA asset classes.
 * @dev All data is read from AssetConfig set via AssetManagerFacet.
 * @custom:security-contact renan.correa@hubweb3.com
 */
contract MetadataFacet {
    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/

    /// @dev ERC-1155 standard metadata URI event
    event URI(string value, uint256 indexed id);

    /*//////////////////////////////////////////////////////////////
                            EXTERNAL VIEWS
    //////////////////////////////////////////////////////////////*/

    /// @notice Returns the metadata URI for a tokenId (ERC-1155 standard)
    /// @param id The token type ID
    /// @return The URI string
    function uri(uint256 id) external view returns (string memory) {
        return LibAssetStorage.layout().configs[id].uri;
    }

    /// @notice Returns the human-readable name for a tokenId
    /// @param tokenId The asset class
    /// @return The asset name
    function name(uint256 tokenId) external view returns (string memory) {
        return LibAssetStorage.layout().configs[tokenId].name;
    }

    /// @notice Returns the ticker symbol for a tokenId
    /// @param tokenId The asset class
    /// @return The asset symbol
    function symbol(uint256 tokenId) external view returns (string memory) {
        return LibAssetStorage.layout().configs[tokenId].symbol;
    }

    /// @notice Returns the supply cap for a tokenId (0 = unlimited)
    /// @param tokenId The asset class
    /// @return The supply cap
    function supplyCap(uint256 tokenId) external view returns (uint256) {
        return LibAssetStorage.layout().configs[tokenId].supplyCap;
    }

    /// @notice Returns the issuer address for a tokenId
    /// @param tokenId The asset class
    /// @return The issuer address
    function issuer(uint256 tokenId) external view returns (address) {
        return LibAssetStorage.layout().configs[tokenId].issuer;
    }

    /// @notice Returns the allowed countries for a tokenId
    /// @param tokenId The asset class
    /// @return Array of ISO 3166-1 numeric country codes (empty = all)
    function allowedCountries(uint256 tokenId) external view returns (uint16[] memory) {
        return LibAssetStorage.layout().configs[tokenId].allowedCountries;
    }

    /// @notice Returns a full summary of a tokenId's metadata and state
    /// @param tokenId The asset class
    /// @return name_ Asset name
    /// @return symbol_ Asset symbol
    /// @return uri_ Metadata URI
    /// @return totalSupply_ Current minted supply
    /// @return supplyCap_ Maximum supply (0 = unlimited)
    /// @return holderCount_ Number of unique holders
    /// @return issuer_ Authorized issuer address
    /// @return paused_ Whether the asset is paused
    function tokenInfo(uint256 tokenId)
        external
        view
        returns (
            string memory name_,
            string memory symbol_,
            string memory uri_,
            uint256 totalSupply_,
            uint256 supplyCap_,
            uint256 holderCount_,
            address issuer_,
            bool paused_
        )
    {
        AssetConfig storage config = LibAssetStorage.layout().configs[tokenId];
        name_ = config.name;
        symbol_ = config.symbol;
        uri_ = config.uri;
        totalSupply_ = config.totalSupply;
        supplyCap_ = config.supplyCap;
        holderCount_ = LibSupplyStorage.layout().holderCount[tokenId];
        issuer_ = config.issuer;
        paused_ = config.paused;
    }
}
