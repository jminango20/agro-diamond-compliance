# Diamond ERC-3643

ERC-3643 security tokens for Real World Assets, built on the EIP-2535 Diamond Proxy with ERC-1155 multi-token support.

A single Diamond contract manages multiple regulated asset classes — each `tokenId` carries its own compliance rules, identity requirements, and supply controls. Designed for debt securities, fractional real estate, tokenized commodities, and any RWA that requires on-chain KYC/AML enforcement.

## Architecture

```
                     ┌──────────────────────────────┐
                     │     Diamond Proxy (EIP-2535)  │
                     │     single contract address   │
                     └──────────────┬───────────────┘
                                    │ delegatecall
          ┌─────────────────────────┼─────────────────────────┐
          │            │            │            │             │
     ┌────┴────┐ ┌─────┴─────┐ ┌───┴────┐ ┌────┴────┐ ┌──────┴──────┐
     │  Core   │ │  Security │ │ Token  │ │Identity │ │     RWA     │
     │  (3)    │ │    (3)    │ │  (4)   │ │  (3)    │ │     (6)     │
     └─────────┘ └───────────┘ └────────┘ └─────────┘ └─────────────┘
          │            │            │            │             │
          └─────────────────────────┴─────────────────────────┘
                                    │
                     ┌──────────────┴───────────────┐
                     │   11 Isolated Storage Slots   │
                     │  keccak256("diamond.rwa.X") - 1│
                     └──────────────────────────────┘
```

### Standards

| Standard | Role |
|----------|------|
| **ERC-3643 (T-REX)** | On-chain compliance: KYC/AML identity, jurisdiction restrictions, freeze, recovery |
| **EIP-2535 Diamond** | Modular upgradeability via facets; bypasses the 24 KB contract size limit |
| **ERC-1155 Multi-Token** | Multiple asset classes (`tokenId` = asset class) in a single contract |

### Three-Level Regulatory Model

```
Global              ──  Diamond ownership, global pause, RBAC, identity registry
  └─ Per tokenId    ──  compliance module, identity profile, supply cap, allowed countries
       └─ Per holder ── balance partitions (free/locked/custody), freeze, lockup
```

## Facets (19 total)

### Core (3)

| Facet | Purpose |
|-------|---------|
| `DiamondCutFacet` | Add, replace, or remove facets |
| `DiamondLoupeFacet` | Introspection + ERC-165 support |
| `OwnershipFacet` | Ownable2Step ownership transfer |

### Security (3)

| Facet | Purpose |
|-------|---------|
| `AccessControlFacet` | Role-based access control (ISSUER, COMPLIANCE_ADMIN, TRANSFER_AGENT) |
| `PauseFacet` | Global pause + per-asset pause |
| `EmergencyFacet` | Circuit breaker for emergency shutdown |

### Token — ERC-1155 (4)

| Facet | Purpose |
|-------|---------|
| `AssetManagerFacet` | Register and configure asset classes per `tokenId` |
| `ERC1155Facet` | Compliant transfers and balance queries |
| `SupplyFacet` | Mint, burn, forced transfer |
| `MetadataFacet` | Name, symbol, URI per `tokenId` |

### Identity — KYC/AML (3)

| Facet | Purpose |
|-------|---------|
| `IdentityRegistryFacet` | Bind wallets to ONCHAINID + country code |
| `ClaimTopicsFacet` | Define required KYC claim topics per identity profile |
| `TrustedIssuerFacet` | Manage authorized claim issuers |

### Compliance (1)

| Facet | Purpose |
|-------|---------|
| `ComplianceRouterFacet` | Route `canTransfer()` to the compliance module assigned to each `tokenId` |

### RWA Operations (5)

| Facet | Purpose |
|-------|---------|
| `AssetGroupFacet` | Hierarchical asset groups with lazy minting (e.g., building → apartments) |
| `FreezeFacet` | Freeze wallets globally, per asset, or partial amounts; lockup with expiry |
| `RecoveryFacet` | Wallet recovery and balance migration |
| `SnapshotFacet` | Point-in-time balance snapshots |
| `DividendFacet` | Pro-rata dividend distribution linked to snapshots |

### Compliance Modules (pluggable)

| Module | Description |
|--------|-------------|
| `CountryRestrictModule` | ISO-3166 country-based transfer restrictions |
| `MaxBalanceModule` | Maximum token balance per holder |
| `MaxHoldersModule` | Cap on number of unique holders per asset |

## Transfer Flow

Every `safeTransferFrom` passes through 6 validation stages before execution:

```
safeTransferFrom(from, to, tokenId, amount)
  │
  ├─ 1. Protocol paused?           → revert ProtocolPaused
  ├─ 2. Wallet frozen (global)?    → revert WalletFrozenGlobal
  ├─ 3. Asset registered & active? → revert AssetNotRegistered / AssetPaused
  ├─ 4. Wallet frozen (asset)?     → revert WalletFrozenAsset / LockupActive
  ├─ 5. Sufficient free balance?   → revert InsufficientFreeBalance
  ├─ 6. Compliance module check    → revert ComplianceCheckFailed
  │
  ├─ Execute: update balances + holder tracking
  └─ Post-hook: module.transferred() for state updates
```

## Asset Groups & Lazy Minting

The `AssetGroupFacet` enables hierarchical tokenization — a parent asset (e.g., a building) can have child assets (e.g., individual apartments) that are only minted on-chain when sold:

```
createGroup(parentTokenId: 1, name: "Aurora Apartments", maxUnits: 100)
  → groupId: 1                                              ~215k gas

mintUnit(groupId: 1, investor: Alice, fractions: 500)
  → childTokenId: (1 << 128) | 1  →  "Apt 101"            ~600k gas
  → inherits parent's compliance, identity profile, issuer, countries
```

Child tokens that haven't been sold don't exist on-chain — zero gas cost until minted.

## Storage

11 isolated storage namespaces prevent slot collisions during upgrades:

| Library | Domain |
|---------|--------|
| `LibAppStorage` | Owner, pending owner, global pause, protocol version |
| `LibAssetStorage` | Asset configs per `tokenId` (name, symbol, supply cap, compliance module) |
| `LibERC1155Storage` | Balance partitions (free/locked/custody/pending), operator approvals |
| `LibSupplyStorage` | Total supply, holder count, holder tracking per `tokenId` |
| `LibFreezeStorage` | Global freeze, asset freeze, frozen amounts, lockup expiry |
| `LibAccessStorage` | Role mappings and role admin configuration |
| `LibIdentityStorage` | Wallet-to-identity bindings, identity profiles, verification cache |
| `LibComplianceStorage` | Token-to-module mappings, registered modules |
| `LibSnapshotStorage` | Snapshot data with balance captures |
| `LibDividendStorage` | Dividend records with claim tracking |
| `LibAssetGroupStorage` | Group configs, parent-child relationships |

Each slot is derived from `keccak256("diamond.rwa.<domain>.storage") - 1`.

## Project Structure

```
diamond-erc3643/
├── packages/
│   ├── contracts/                    # Solidity (Foundry)
│   │   ├── src/
│   │   │   ├── Diamond.sol           # EIP-2535 proxy
│   │   │   ├── facets/               # 19 facets by domain
│   │   │   │   ├── core/
│   │   │   │   ├── token/
│   │   │   │   ├── identity/
│   │   │   │   ├── compliance/
│   │   │   │   ├── rwa/
│   │   │   │   └── security/
│   │   │   ├── compliance/modules/   # Pluggable compliance modules
│   │   │   ├── interfaces/           # Domain-organized interfaces
│   │   │   ├── libraries/            # LibDiamond, LibAppStorage, LibReasonCodes
│   │   │   ├── storage/              # 11 namespaced storage libraries
│   │   │   └── initializers/         # DiamondInit
│   │   ├── test/
│   │   │   ├── unit/                 # 26 test files (1:1 with facets + modules)
│   │   │   ├── fuzz/                 # Property-based tests
│   │   │   ├── invariant/            # FREI-PI pattern invariant tests
│   │   │   └── helpers/              # DiamondHelper, MockComplianceModule
│   │   ├── script/                   # Deploy.s.sol, ConfigureAsset.s.sol
│   │   ├── foundry.toml
│   │   ├── remappings.txt
│   │   └── Makefile
│   └── indexer/                      # Go blockchain event indexer
├── docs/
│   ├── architecture.md               # Detailed architecture specification
│   └── diagrams/                     # Mermaid diagrams (EN + PT)
│       ├── en/                       # 9 diagrams in English
│       ├── pt/                       # 9 diagrams in Portuguese
│       └── viewer.html               # Interactive diagram viewer
├── .github/workflows/ci.yml          # CI pipeline (7 parallel jobs)
├── turbo.json
├── pnpm-workspace.yaml
└── package.json
```

## Getting Started

### Prerequisites

- [Foundry](https://book.getfoundry.sh/getting-started/installation) (forge, cast, anvil)
- [Node.js](https://nodejs.org/) >= 20
- [pnpm](https://pnpm.io/) >= 10

### Install

```bash
git clone git@github.com:renancorreadev/diamond-erc3643.git
cd diamond-erc3643
pnpm install
cd packages/contracts
forge install
```

### Build

```bash
make build           # forge build --sizes
```

### Test

```bash
make test            # all tests
make test-unit       # unit tests only
make test-fuzz       # fuzz tests (10,000 runs)
make test-invariant  # invariant tests (10,000 runs, depth 500)
make test-contract CONTRACT=ERC1155Facet   # single contract
```

### Coverage & Analysis

```bash
make coverage        # forge coverage → lcov.info
make slither         # static analysis
make lint            # solhint
```

### Deploy (Local)

```bash
# Terminal 1
anvil

# Terminal 2
make deploy-local
```

### Deploy (Testnet)

```bash
# Import deployer key (interactive, never in plaintext)
cast wallet import deployer --interactive

# Deploy
forge script script/Deploy.s.sol \
  --rpc-url $AMOY_RPC_URL \
  --account deployer \
  --broadcast \
  --verify \
  --etherscan-api-key $POLYGONSCAN_API_KEY
```

## CI/CD

The CI pipeline runs 7 parallel jobs on every PR:

| Job | What it does |
|-----|-------------|
| **Lint** | `solhint src/**/*.sol` |
| **Build** | `forge build --sizes` (checks 24 KB limit) |
| **Test** | `forge test -vvv` |
| **Fuzz** | 10,000 fuzz runs on `test/fuzz/**` |
| **Invariant** | 10,000 invariant runs on `test/invariant/**` |
| **Slither** | Static analysis with Slither |
| **Coverage** | `forge coverage` → Codecov |

## Roles & Permissions

| Role | Capabilities |
|------|-------------|
| **Owner** | `diamondCut`, `transferOwnership`, `emergencyPause`, `setRoleAdmin` |
| **COMPLIANCE_ADMIN** | `registerAsset`, `createGroup`, `setComplianceModule`, `setIdentityProfile`, `registerIdentity` |
| **ISSUER_ROLE** | `mint`, `burn`, `mintUnit`, `mintUnitBatch`, `createDividend` |
| **TRANSFER_AGENT** | `forcedTransfer`, `recoverWallet` |
| **Token Holder** | `safeTransferFrom` (if compliance passes), `claimDividend`, `setApprovalForAll` |

## Documentation

- **[Architecture Spec](docs/architecture.md)** — design principles, regulatory model, facet map, storage layout, transfer validation pipeline, compliance modules, reason codes, roles, events
- **[Diagrams](docs/diagrams/)** — 9 Mermaid diagrams covering the full system, available in [English](docs/diagrams/en/) and [Portuguese](docs/diagrams/pt/); open `docs/diagrams/viewer.html` for an interactive viewer

### Diagram Index

| # | Diagram |
|---|---------|
| 01 | High-Level Architecture — Diamond proxy + 19 facets + 11 storage slots |
| 02 | Regulated Transfer Flow — 6 validation stages with revert paths |
| 03 | Asset Groups & Lazy Minting — `createGroup` → `mintUnit` with gas costs |
| 04 | Storage Layout — 11 namespaced slots with all fields |
| 05 | Roles & Permissions — RBAC tree with all operations |
| 06 | Compliance Pipeline — Sequence diagram of full transfer validation |
| 07 | Full Technology Stack — Frontend → Indexer → Blockchain |
| 08 | Token Lifecycle — State machine: register → mint → freeze → snapshot → dividend |
| 09 | Real Estate Example — End-to-end building tokenization with rent distribution |

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Smart Contracts | Solidity 0.8.28, Foundry, OpenZeppelin |
| Proxy Pattern | EIP-2535 Diamond (Nick Mudge reference) |
| Token Standard | ERC-1155 with ERC-3643 compliance hooks |
| Identity | ONCHAINID (ERC-734 keys + ERC-735 claims) |
| Testing | Forge (unit + fuzz + invariant), Slither |
| Indexer | Go + RocksDB + GraphQL |
| CI/CD | GitHub Actions, Codecov, Changesets |
| Monorepo | pnpm workspaces + Turborepo |

## References

- [ERC-3643 (T-REX)](https://github.com/ERC-3643/ERC-3643) — Security Token Standard
- [EIP-2535 Diamond Standard](https://eips.ethereum.org/EIPS/eip-2535) — Multi-Facet Proxy
- [Nick Mudge Diamond Reference](https://github.com/mudgen/diamond-3-hardhat)
- [Trail of Bits: Building Secure Contracts](https://github.com/crytic/building-secure-contracts)
- [ONCHAINID](https://github.com/onchain-id/solidity) — On-Chain Identity

## License

MIT
