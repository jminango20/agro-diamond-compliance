# GraphQL Queries — Diamond ERC-3643 Indexer

Queries organizadas por dominio do protocolo. Cada arquivo `.graphql` contem a query + exemplos de variaveis nos comentarios.

## Estrutura

```
queries/
├── tokens/                          # ERC-1155 token data (supply, holders, transfers)
│   ├── all-tokens.graphql           # Lista todos os tokens
│   ├── all-tokens-with-holders.graphql  # Tokens + top holders
│   ├── token-by-id.graphql          # Token por ID (holders + eventos)
│   ├── token-events.graphql         # Eventos de transferencia de um token
│   ├── token-top-holders.graphql    # Top N holders de um token
│   ├── holder-balance.graphql       # Balance de um holder especifico
│   └── recent-transfers.graphql     # Transferencias recentes (global)
│
├── assets/                          # Asset metadata (name, symbol, issuer, pause)
│   ├── all-assets.graphql           # Lista todos os assets com metadata
│   ├── asset-by-id.graphql          # Asset completo (holders + eventos)
│   ├── asset-holders.graphql        # Holders de um asset
│   ├── asset-events.graphql         # Eventos de um asset (filtro por tipo)
│   └── asset-status.graphql         # Status de pause de todos os assets
│
├── identity/                        # KYC/Identity Registry (ONCHAINID)
│   ├── identity-by-wallet.graphql   # Identidade de uma wallet
│   ├── all-identities.graphql       # Todas as identidades
│   ├── identities-by-country.graphql    # Filtro por pais (ISO 3166-1)
│   ├── identity-events.graphql      # Eventos de identity_bound
│   └── identity-events-all-types.graphql  # Bind + unbind + recovery
│
├── freeze/                          # Freeze, partial freeze, lockup
│   ├── freezes-by-wallet.graphql    # Freezes de uma wallet (global + per-asset)
│   ├── frozen-wallets.graphql       # Wallets congeladas (filtro por token)
│   ├── freeze-events.graphql        # Eventos de wallet_frozen
│   ├── partial-freeze-events.graphql    # Eventos de partial freeze
│   └── lockup-events.graphql        # Eventos de lockup
│
├── plugins/                         # Plugin system (global + per-asset hookable)
│   ├── global-plugin-events.graphql     # Plugins globais (register, remove, status)
│   ├── asset-plugin-events.graphql      # Plugins per-asset (add, remove, set)
│   └── all-plugin-events.graphql        # Todos os tipos de eventos de plugins
│
├── yield/                           # YieldDistributorModule (Synthetix/MasterChef)
│   ├── yield-deposited.graphql      # Depositos de yield
│   ├── yield-claimed.graphql        # Claims de yield por holders
│   ├── reward-asset-events.graphql  # Reward assets (add/remove)
│   └── yield-overview.graphql       # Overview completo (deposits + claims + rewards)
│
├── compliance-modules/              # Modulos de compliance externos
│   ├── country-restrict-events.graphql  # Paises restritos/liberados
│   ├── max-balance-events.graphql       # Limites de balance por holder
│   ├── max-holders-events.graphql       # Limites de holders unicos
│   └── compliance-overview.graphql      # Overview (modules + restricoes + limites)
│
├── protocol-events/                 # Eventos genericos do protocolo (33+ tipos)
│   ├── all-events.graphql           # Todos os eventos recentes
│   ├── events-by-type.graphql       # Filtro por eventType (todos os tipos)
│   ├── events-by-token.graphql      # Filtro por tokenId
│   ├── events-by-address.graphql    # Filtro por endereco
│   ├── events-by-type-and-token.graphql  # Filtro combinado
│   ├── snapshot-events.graphql      # Snapshots criados
│   ├── dividend-events.graphql      # Dividendos (criacao + claim)
│   ├── role-events.graphql          # Roles (granted + revoked)
│   ├── pause-events.graphql         # Pause/unpause (protocol + asset)
│   ├── compliance-events.graphql    # Compliance modules (add + remove)
│   └── asset-group-events.graphql   # Asset groups (create + mint)
│
├── portfolio/                       # Wallet portfolio e overview
│   ├── portfolio.graphql            # Holdings de uma wallet
│   ├── wallet-overview.graphql      # Overview completo (identity + portfolio + freezes + eventos)
│   └── wallet-portfolio-with-assets.graphql  # Portfolio + metadata dos assets
│
├── dashboard/                       # Dashboard e status
│   ├── dashboard-overview.graphql   # Overview completo (status + assets + tokens + eventos)
│   ├── indexer-status.graphql       # Status de sincronizacao
│   └── protocol-summary.graphql     # Resumo do protocolo
│
└── curl-examples.sh                 # 30+ exemplos curl com jq
```

## Tipos de Evento (eventType)

| Categoria | eventType | Descricao |
|-----------|-----------|-----------|
| **Asset** | `asset_registered` | Novo asset registrado |
| | `asset_config_updated` | Config de asset atualizada |
| | `uri_updated` | URI de metadata atualizada |
| **Identity** | `identity_bound` | Wallet vinculada a ONCHAINID |
| | `identity_unbound` | Wallet desvinculada |
| **Freeze** | `wallet_frozen` | Wallet congelada/descongelada (global) |
| | `asset_frozen` | Wallet congelada para asset especifico |
| | `partial_freeze` | Quantidade parcialmente congelada |
| | `lockup_set` | Lockup com expiracao temporal |
| **Access** | `role_granted` | Role concedido |
| | `role_revoked` | Role revogado |
| **Pause** | `emergency_pause` | Emergency pause global |
| | `protocol_unpaused` | Protocolo despausado |
| | `asset_paused` | Asset pausado |
| | `asset_unpaused` | Asset despausado |
| **Recovery** | `wallet_recovered` | Wallet recuperada |
| **Snapshot** | `snapshot_created` | Snapshot criado |
| **Dividend** | `dividend_created` | Dividendo criado |
| | `dividend_claimed` | Dividendo reivindicado |
| **Compliance** | `compliance_module_added` | Modulo adicionado |
| | `compliance_module_removed` | Modulo removido |
| **Groups** | `group_created` | Grupo de assets criado |
| | `unit_minted` | Unidade mintada em grupo |
| **Plugins (Global)** | `global_plugin_registered` | Plugin global registrado |
| | `global_plugin_removed` | Plugin global removido |
| | `global_plugin_status_changed` | Plugin ativado/desativado |
| **Plugins (Asset)** | `plugin_module_added` | Plugin hookable adicionado a asset |
| | `plugin_module_removed` | Plugin hookable removido |
| | `plugin_modules_set` | Plugins definidos em batch |
| **Yield** | `reward_asset_added` | Reward asset registrado |
| | `reward_asset_removed` | Reward asset removido |
| | `yield_deposited` | Yield depositado para distribuicao |
| | `yield_claimed` | Holder reivindicou yield |
| **Country Restrict** | `country_restricted` | Pais adicionado a blocklist |
| | `country_unrestricted` | Pais removido da blocklist |
| **Max Balance** | `max_balance_set` | Limite de balance definido |
| **Max Holders** | `max_holders_set` | Limite de holders definido |

## Uso

### GraphQL Playground
Acesse `http://localhost:8080` no navegador.

### curl
```bash
# Rodar todos os exemplos
INDEXER_URL=http://localhost:8080 ./queries/curl-examples.sh

# Exemplo individual
curl -s http://localhost:8080/graphql \
  -H 'Content-Type: application/json' \
  -d '{"query": "{ assets { id name symbol } }"}' | jq .
```
