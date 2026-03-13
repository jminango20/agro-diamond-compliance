# ── Stage 1: Build ──
FROM golang:1.24-bookworm AS builder

# Install RocksDB build deps
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    librocksdb-dev \
    libsnappy-dev \
    liblz4-dev \
    libzstd-dev \
    libbz2-dev \
    zlib1g-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /src

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .

ENV CGO_ENABLED=1
RUN go build -ldflags="-s -w" -o /indexer ./cmd/indexer/

# ── Stage 2: Runtime ──
FROM debian:bookworm-slim

# Install runtime libs for RocksDB (use wildcard to match available version)
RUN apt-get update && apt-get install -y --no-install-recommends \
    librocksdb-dev \
    libsnappy1v5 \
    liblz4-1 \
    libzstd1 \
    libbz2-1.0 \
    zlib1g \
    ca-certificates \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN useradd -r -s /bin/false indexer

WORKDIR /app

# Copy binary
COPY --from=builder /indexer /app/indexer

# Data volume
RUN mkdir -p /app/data && chown indexer:indexer /app/data
VOLUME ["/app/data"]

USER indexer

ENV HTTP_LISTEN=:8080
ENV DATA_DIR=/app/data

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=3s --start-period=10s \
  CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/indexer"]
