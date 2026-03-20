// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

/// @title AgroClaimTopics
/// @notice ERC-735 claim topic IDs mapped to Brazilian agricultural regulations.
///         Each constant is a unique uint256 identifier used to register and verify
///         on-chain claims issued by trusted Brazilian government authorities.
///
///         Claim issuers (government oracles, authorized registrars) publish claims
///         against producer wallets. The Diamond's identity registry verifies that a
///         holder possesses all topics required by the asset's compliance profile
///         before allowing transfers.
///
/// @dev Topic IDs 100-199 are reserved for Brazilian agri-regulatory claims.
///      Topic IDs 1-99 are reserved for generic ERC-735 / T-REX standard topics.
library AgroClaimTopics {
    // ─── Brazilian Tax Identity ────────────────────────────────────────────────

    /// @notice CPF (individuals) or CNPJ (legal entities) verification.
    ///         Required for all producers and investors operating in Brazil.
    ///         Claim issuer: Receita Federal do Brasil oracle.
    uint256 internal constant CLAIM_CPF_CNPJ = 101;

    // ─── Ministry of Agriculture (MAPA) ───────────────────────────────────────

    /// @notice MAPA producer registration number (Cadastro de Produtores Rurais).
    ///         Required for any wallet that mints or receives tokenized agricultural
    ///         commodities on this platform. Verifies the holder is a registered
    ///         rural producer under MAPA supervision.
    uint256 internal constant CLAIM_MAPA_REGISTRATION = 102;

    // ─── Rural Credit System (SNCR) ───────────────────────────────────────────

    /// @notice Enrollment in the Sistema Nacional de Crédito Rural (SNCR).
    ///         Required for tokenIds that represent assets eligible as collateral
    ///         in rural credit operations (e.g., CPR-backed coffee lots).
    uint256 internal constant CLAIM_SNCR_ENROLLMENT = 103;

    // ─── Environmental Registry (CAR) ─────────────────────────────────────────

    /// @notice Cadastro Ambiental Rural (CAR) registration.
    ///         Required for coffee and honey assets — environmental compliance
    ///         is a prerequisite for legal commercialization of these commodities.
    ///         Verifies the rural property is registered with SICAR.
    uint256 internal constant CLAIM_CAR_REGISTRATION = 104;

    // ─── Sanitary / Export Authorization ──────────────────────────────────────

    /// @notice ANVISA or MAPA export authorization for the commodity class.
    ///         Applied via CountryRestrictModule + this claim for cross-border
    ///         transfers of honey and coffee to restricted jurisdictions.
    uint256 internal constant CLAIM_EXPORT_AUTHORIZATION = 105;

    // ─── CPR Holder ───────────────────────────────────────────────────────────

    /// @notice Cédula de Produto Rural (CPR) holder claim.
    ///         Certifies that the wallet has an active CPR instrument registered
    ///         with a Brazilian financial institution. Used in conjunction with
    ///         CprLockupModule to enforce settlement lock-up periods.
    uint256 internal constant CLAIM_CPR_HOLDER = 106;

    // ─── Livestock Traceability (SISBOV) ──────────────────────────────────────

    /// @notice SISBOV (Sistema Brasileiro de Identificação e Certificação de
    ///         Bovinos e Bubalinos) registration.
    ///         Required for cattle tokenIds — each NFT (amount = 1) represents
    ///         an individual or management-group entry in SISBOV.
    uint256 internal constant CLAIM_SISBOV_REGISTRATION = 107;
}
