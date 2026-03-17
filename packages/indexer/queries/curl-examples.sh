#!/bin/bash
# ============================================================================
# Exemplos de chamadas curl para o GraphQL do indexer
# Uso: INDEXER_URL=http://localhost:8080 ./curl-examples.sh
# ============================================================================

URL="${INDEXER_URL:-http://localhost:8080}/graphql"

echo "=== Endpoint: $URL ==="
echo ""

# --------------------------------------------------------------------------
# DASHBOARD
# --------------------------------------------------------------------------

echo "--- Indexer Status ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ status { lastBlock tokenCount } }"
}' | jq .

echo ""
echo "--- Dashboard Overview ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ status { lastBlock tokenCount } assets { id name symbol paused totalSupply holderCount } tokens { id totalSupply holderCount } events(first: 5) { txHash block tokenId eventType } protocolEvents(first: 5) { txHash block eventType data } }"
}' | jq .

# --------------------------------------------------------------------------
# TOKENS
# --------------------------------------------------------------------------

echo ""
echo "--- All Tokens ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ tokens { id totalSupply holderCount } }"
}' | jq .

echo ""
echo "--- Token 1 Detail ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($id: String!) { token(id: $id) { id totalSupply holderCount holders(first: 10) { address balance } events(first: 5) { txHash block from to amount eventType } } }",
  "variables": { "id": "1" }
}' | jq .

echo ""
echo "--- Token 1 Top 5 Holders ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($id: String!, $first: Int!) { token(id: $id) { id holderCount holders(first: $first) { address balance } } }",
  "variables": { "id": "1", "first": 5 }
}' | jq .

echo ""
echo "--- Holder Balance (Token 1, Deployer) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($tokenId: String!, $address: String!) { holder(tokenId: $tokenId, address: $address) { address balance } }",
  "variables": { "tokenId": "1", "address": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- Recent Transfers (last 10) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ events(first: 10) { txHash block tokenId from to amount eventType } }"
}' | jq .

# --------------------------------------------------------------------------
# ASSETS
# --------------------------------------------------------------------------

echo ""
echo "--- All Assets ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ assets { id name symbol issuer profileId paused totalSupply holderCount registeredAt } }"
}' | jq .

echo ""
echo "--- Asset 1 Detail (with holders + events) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($id: String!) { asset(id: $id) { id name symbol issuer profileId uri paused totalSupply holderCount registeredAt holders(first: 10) { address balance } events(first: 5) { txHash block eventType data } } }",
  "variables": { "id": "1" }
}' | jq .

echo ""
echo "--- Asset 1 Events (only snapshots) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($id: String!, $eventType: String) { asset(id: $id) { id name events(first: 20, eventType: $eventType) { txHash block eventType data } } }",
  "variables": { "id": "1", "eventType": "snapshot_created" }
}' | jq .

echo ""
echo "--- Asset Status (paused check) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ assets { id name symbol paused } }"
}' | jq .

# --------------------------------------------------------------------------
# IDENTITY
# --------------------------------------------------------------------------

echo ""
echo "--- Identity Lookup ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($wallet: String!) { identity(wallet: $wallet) { wallet identity country boundAt } }",
  "variables": { "wallet": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- All Identities (first 20) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ identities(first: 20) { wallet identity country boundAt } }"
}' | jq .

echo ""
echo "--- Identities by Country (Brazil = 76) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($country: Int!, $first: Int!) { identities(country: $country, first: $first) { wallet identity country boundAt } }",
  "variables": { "country": 76, "first": 50 }
}' | jq .

echo ""
echo "--- Identities by Country (USA = 840) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($country: Int!, $first: Int!) { identities(country: $country, first: $first) { wallet identity country boundAt } }",
  "variables": { "country": 840, "first": 50 }
}' | jq .

# --------------------------------------------------------------------------
# FREEZE
# --------------------------------------------------------------------------

echo ""
echo "--- Freezes by Wallet ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($wallet: String!) { freezes(wallet: $wallet) { wallet tokenId frozen frozenAmount lockupExpiry } }",
  "variables": { "wallet": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- All Frozen Wallets ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ frozenWallets { wallet tokenId frozen frozenAmount lockupExpiry } }"
}' | jq .

echo ""
echo "--- Frozen Wallets for Token 1 ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($tokenId: String) { frozenWallets(tokenId: $tokenId) { wallet tokenId frozen frozenAmount lockupExpiry } }",
  "variables": { "tokenId": "1" }
}' | jq .

# --------------------------------------------------------------------------
# PROTOCOL EVENTS
# --------------------------------------------------------------------------

echo ""
echo "--- All Protocol Events (last 10) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "{ protocolEvents(first: 10) { txHash block logIndex eventType tokenId address data } }"
}' | jq .

echo ""
echo "--- Protocol Events: identity_bound ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType address data } }",
  "variables": { "eventType": "identity_bound" }
}' | jq .

echo ""
echo "--- Protocol Events: asset_registered ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType tokenId data } }",
  "variables": { "eventType": "asset_registered" }
}' | jq .

echo ""
echo "--- Protocol Events: role_granted ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType address data } }",
  "variables": { "eventType": "role_granted" }
}' | jq .

echo ""
echo "--- Protocol Events: snapshot_created ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType tokenId data } }",
  "variables": { "eventType": "snapshot_created" }
}' | jq .

echo ""
echo "--- Protocol Events: dividend_created ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType tokenId data } }",
  "variables": { "eventType": "dividend_created" }
}' | jq .

echo ""
echo "--- Protocol Events by Token 1 ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($tokenId: String!) { protocolEvents(first: 20, tokenId: $tokenId) { txHash block eventType tokenId data } }",
  "variables": { "tokenId": "1" }
}' | jq .

echo ""
echo "--- Protocol Events by Address ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($address: String!) { protocolEvents(first: 20, address: $address) { txHash block eventType tokenId data } }",
  "variables": { "address": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- Protocol Events: emergency_pause ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 10, eventType: $eventType) { txHash block eventType address data } }",
  "variables": { "eventType": "emergency_pause" }
}' | jq .

echo ""
echo "--- Protocol Events: wallet_frozen ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType address data } }",
  "variables": { "eventType": "wallet_frozen" }
}' | jq .

echo ""
echo "--- Protocol Events: compliance_module_added ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($eventType: String!) { protocolEvents(first: 20, eventType: $eventType) { txHash block eventType tokenId data } }",
  "variables": { "eventType": "compliance_module_added" }
}' | jq .

# --------------------------------------------------------------------------
# PORTFOLIO
# --------------------------------------------------------------------------

echo ""
echo "--- Portfolio ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($address: String!) { portfolio(address: $address) { tokenId balance } }",
  "variables": { "address": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- Wallet Overview (combined) ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($address: String!) { identity(wallet: $address) { wallet identity country boundAt } portfolio(address: $address) { tokenId balance } freezes(wallet: $address) { tokenId frozen frozenAmount lockupExpiry } protocolEvents(first: 10, address: $address) { txHash block eventType tokenId data } }",
  "variables": { "address": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "--- Portfolio with Asset Metadata ---"
curl -s "$URL" -H 'Content-Type: application/json' -d '{
  "query": "query($address: String!) { portfolio(address: $address) { tokenId balance } assets { id name symbol paused totalSupply holderCount } }",
  "variables": { "address": "0xB40061C7bf8394eb130Fcb5EA06868064593BFAa" }
}' | jq .

echo ""
echo "=== Done ==="
