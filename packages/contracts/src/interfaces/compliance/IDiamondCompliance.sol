// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title IDiamondCompliance
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Minimal interface for compliance modules to query the Diamond.
///         Modules call back into the Diamond to read identity, balance,
///         and supply data needed for transfer validation.
/// @custom:security-contact renan.correa@hubweb3.com
interface IDiamondCompliance {
    /// @notice Returns the country code for a wallet (from IdentityRegistryFacet)
    function getCountry(address wallet) external view returns (uint16);

    /// @notice Returns the free balance of a holder for a tokenId (from ERC1155Facet)
    function balanceOf(address account, uint256 id) external view returns (uint256);

    /// @notice Returns unique holder count for a tokenId (from SupplyFacet)
    function holderCount(uint256 tokenId) external view returns (uint256);

    /// @notice Returns the full partition balance (from ERC1155Facet)
    function partitionBalanceOf(address account, uint256 id)
        external
        view
        returns (uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement);
}
