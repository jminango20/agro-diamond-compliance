// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {AssetConfig} from "../../storage/LibAssetStorage.sol";

interface IAssetManager {
    struct RegisterAssetParams {
        string name;
        string symbol;
        string uri;
        uint256 supplyCap;
        uint32 identityProfileId;
        address[] complianceModules;
        address issuer;
        uint16[] allowedCountries;
    }

    function registerAsset(RegisterAssetParams calldata params) external returns (uint256 tokenId);
    function addComplianceModule(uint256 tokenId, address module) external;
    function removeComplianceModule(uint256 tokenId, address module) external;
    function setComplianceModules(uint256 tokenId, address[] calldata modules) external;
    function setIdentityProfile(uint256 tokenId, uint32 profileId) external;
    function setIssuer(uint256 tokenId, address issuer) external;
    function setSupplyCap(uint256 tokenId, uint256 cap) external;
    function setAllowedCountries(uint256 tokenId, uint16[] calldata countries) external;
    function setAssetUri(uint256 tokenId, string calldata uri) external;

    function getAssetConfig(uint256 tokenId) external view returns (AssetConfig memory);
    function getComplianceModules(uint256 tokenId) external view returns (address[] memory);
    function getRegisteredTokenIds() external view returns (uint256[] memory);
    function assetExists(uint256 tokenId) external view returns (bool);
    function nextTokenId() external view returns (uint256);
}
