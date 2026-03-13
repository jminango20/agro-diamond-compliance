#!/bin/bash
# Exemplos de queries via curl contra o indexer GraphQL
# Endpoint: http://localhost:8080/graphql
# Playground: http://localhost:8080/

ENDPOINT="http://localhost:8080/graphql"

echo "=== Indexer Status ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ status { lastBlock tokenCount } }"}' | jq .

echo ""
echo "=== All Tokens ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ tokens { id totalSupply holderCount holders(first:10) { address balance } } }"}' | jq .

echo ""
echo "=== Token 1 Detail ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ token(id:\"1\") { id totalSupply holderCount holders(first:50) { address balance } events(first:20) { txHash block from to amount eventType } } }"}' | jq .

echo ""
echo "=== Holder Balance (Owner) ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ holder(tokenId:\"1\", address:\"0xB40061C7bf8394eb130Fcb5EA06868064593BFAa\") { address balance } }"}' | jq .

echo ""
echo "=== Recent Events (last 50) ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ events(first:50) { txHash block tokenId from to amount eventType } }"}' | jq .

echo ""
echo "=== Dashboard Overview ==="
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{"query":"{ status { lastBlock tokenCount } tokens { id totalSupply holderCount } events(first:10) { txHash tokenId from to amount eventType } }"}' | jq .
