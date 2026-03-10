.PHONY: build run generate clean

# Build the indexer binary
build:
	go build -o bin/indexer ./cmd/indexer/

# Run the indexer (requires DIAMOND_ADDRESS env var)
run: build
	./bin/indexer

# Generate Go bindings from Forge artifacts
generate:
	./scripts/generate_bindings.sh

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
