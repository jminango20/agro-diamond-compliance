#!/usr/bin/env bash
# Extracts ABIs from Forge artifacts and generates Go bindings via abigen.
# Usage: ./scripts/generate_bindings.sh

set -euo pipefail

ABIGEN="${ABIGEN:-$HOME/go/bin/abigen}"
ARTIFACTS_DIR="../contracts/out"
BINDINGS_DIR="./internal/bindings"
ABI_DIR="./abi"

mkdir -p "$BINDINGS_DIR" "$ABI_DIR"

# Facets to generate bindings for
FACETS=(
    "ERC1155Facet"
    "SupplyFacet"
    "AssetManagerFacet"
    "MetadataFacet"
    "AccessControlFacet"
    "PauseFacet"
    "EmergencyFacet"
    "FreezeFacet"
    "RecoveryFacet"
    "SnapshotFacet"
    "DividendFacet"
    "IdentityRegistryFacet"
    "ClaimTopicsFacet"
    "TrustedIssuerFacet"
    "ComplianceRouterFacet"
    "OwnershipFacet"
)

for facet in "${FACETS[@]}"; do
    artifact="$ARTIFACTS_DIR/${facet}.sol/${facet}.json"

    if [ ! -f "$artifact" ]; then
        echo "SKIP: $artifact not found"
        continue
    fi

    # Extract ABI from forge artifact
    python3 -c "
import json, sys
with open('$artifact') as f:
    data = json.load(f)
print(json.dumps(data['abi']))
" > "$ABI_DIR/${facet}.abi.json"

    # Lowercase filename, single 'bindings' package
    filename=$(echo "$facet" | tr '[:upper:]' '[:lower:]')

    # Generate Go binding
    "$ABIGEN" \
        --abi "$ABI_DIR/${facet}.abi.json" \
        --pkg bindings \
        --type "$facet" \
        --out "$BINDINGS_DIR/${filename}.go"

    echo "OK: $facet -> $BINDINGS_DIR/${filename}.go"
done

echo ""
echo "Done. Generated $(ls -1 "$BINDINGS_DIR"/*.go 2>/dev/null | wc -l | tr -d ' ') binding files."
