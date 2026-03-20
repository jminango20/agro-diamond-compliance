// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IDiamondCompliance} from "./IDiamondCompliance.sol";

/// @title IAgroDiamondCompliance
/// @notice Extends IDiamondCompliance with identity registry queries needed
///         by agricultural compliance modules to enforce producer KYC.
///         Modules cast the Diamond address to this interface to verify
///         whether a wallet is registered as a KYC'd agricultural producer.
interface IAgroDiamondCompliance is IDiamondCompliance {
    /// @notice Returns true if `wallet` is registered in the identity registry.
    ///         Only wallets that have passed KYC and been bound to an ONCHAINID
    ///         by a TRANSFER_AGENT are registered (from IdentityRegistryFacet).
    function contains(address wallet) external view returns (bool);
}
