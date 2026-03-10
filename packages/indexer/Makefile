.PHONY: build run generate clean test

# RocksDB CGO flags (macOS Homebrew)
export CGO_CFLAGS  = -I/opt/homebrew/include
export CGO_LDFLAGS = -L/opt/homebrew/lib -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd
export CGO_ENABLED = 1

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

# Clean build artifacts and data
clean:
	rm -rf bin/ data/
