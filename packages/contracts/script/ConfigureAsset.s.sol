// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Script, console2} from "forge-std/Script.sol";

import {AccessControlFacet} from "../src/facets/security/AccessControlFacet.sol";
import {AssetManagerFacet} from "../src/facets/token/AssetManagerFacet.sol";
import {IAssetManager} from "../src/interfaces/token/IAssetManager.sol";

/// @title ConfigureAsset
/// @author Renan Correa <renan.correa@hubweb3.com>
/// @notice Post-deploy script: grants roles and registers an example asset.
///         Customize the parameters below for your use case.
///
/// Usage:
///   DIAMOND=0x... forge script script/ConfigureAsset.s.sol \
///     --rpc-url http://localhost:8545 --broadcast -vvvv
///
/// Environment:
///   DIAMOND          — deployed Diamond address (required)
///   ISSUER_ADDRESS   — address to grant ISSUER_ROLE (defaults to msg.sender)
///   AGENT_ADDRESS    — address to grant TRANSFER_AGENT (defaults to msg.sender)
///   RECOVERY_ADDRESS — address to grant RECOVERY_AGENT (defaults to msg.sender)
///   TOKEN_ID         — tokenId to register (defaults to 1)
///   TOKEN_NAME       — asset name (defaults to "Security Token A")
///   TOKEN_SYMBOL     — asset symbol (defaults to "STA")
///   SUPPLY_CAP       — max supply, 0 = unlimited (defaults to 0)
contract ConfigureAsset is Script {
    bytes32 internal constant ISSUER_ROLE = keccak256("ISSUER_ROLE");
    bytes32 internal constant TRANSFER_AGENT = keccak256("TRANSFER_AGENT");
    bytes32 internal constant RECOVERY_AGENT = keccak256("RECOVERY_AGENT");
    bytes32 internal constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    function run() external {
        address diamond = vm.envAddress("DIAMOND");
        address issuerAddr = vm.envOr("ISSUER_ADDRESS", msg.sender);
        address agentAddr = vm.envOr("AGENT_ADDRESS", msg.sender);
        address recoveryAddr = vm.envOr("RECOVERY_ADDRESS", msg.sender);

        uint256 tokenId = vm.envOr("TOKEN_ID", uint256(1));
        string memory tokenName = vm.envOr("TOKEN_NAME", string("Security Token A"));
        string memory tokenSymbol = vm.envOr("TOKEN_SYMBOL", string("STA"));
        uint256 supplyCap = vm.envOr("SUPPLY_CAP", uint256(0));

        AccessControlFacet ac = AccessControlFacet(diamond);
        AssetManagerFacet am = AssetManagerFacet(diamond);

        vm.startBroadcast();

        // ── 1. Grant roles ──────────────────────────────────────────

        ac.grantRole(ISSUER_ROLE, issuerAddr);
        ac.grantRole(TRANSFER_AGENT, agentAddr);
        ac.grantRole(RECOVERY_AGENT, recoveryAddr);
        ac.grantRole(PAUSER_ROLE, msg.sender);

        console2.log("");
        console2.log("=== Roles Granted ===");
        console2.log("ISSUER_ROLE    :", issuerAddr);
        console2.log("TRANSFER_AGENT :", agentAddr);
        console2.log("RECOVERY_AGENT :", recoveryAddr);
        console2.log("PAUSER_ROLE    :", msg.sender);

        // ── 2. Register asset ───────────────────────────────────────

        uint16[] memory countries = new uint16[](0);
        address[] memory modules = new address[](0);

        uint256 registeredTokenId = am.registerAsset(
            IAssetManager.RegisterAssetParams({
                name: tokenName,
                symbol: tokenSymbol,
                uri: "",
                supplyCap: supplyCap,
                identityProfileId: 0,
                complianceModules: modules,
                issuer: issuerAddr,
                allowedCountries: countries
            })
        );

        vm.stopBroadcast();

        console2.log("");
        console2.log("=== Asset Registered ===");
        console2.log("Token ID       :", registeredTokenId);
        console2.log("Name           :", tokenName);
        console2.log("Symbol         :", tokenSymbol);
        console2.log("Supply Cap     :", supplyCap);
        console2.log("Issuer         :", issuerAddr);
        console2.log("Compliance     : none (address(0))");
        console2.log("");
    }
}
