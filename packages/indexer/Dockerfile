# ── Stage 1: Build with Ubuntu (has rocksdb PPA) ──
FROM ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

# Install Go 1.24 + RocksDB dev from Ubuntu repos
RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-go \
    librocksdb-dev \
    build-essential \
    ca-certificates \
    git \
    && rm -rf /var/lib/apt/lists/*

# Ubuntu 24.04 ships Go 1.22 — we need 1.24, install manually
RUN rm -rf /usr/local/go && \
    apt-get update && apt-get install -y wget && \
    wget -qO- https://go.dev/dl/go1.24.5.linux-amd64.tar.gz | tar -C /usr/local -xz && \
    rm -rf /var/lib/apt/lists/*
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH=/go
ENV PATH="${GOPATH}/bin:${PATH}"

WORKDIR /src

# Cache Go modules
COPY go.mod go.sum ./
ENV CGO_ENABLED=1
RUN go mod download

# Copy source and build (static link rocksdb)
COPY . .
RUN go build -ldflags="-s -w" -o /indexer ./cmd/indexer/

# ── Stage 2: Minimal runtime ──
FROM ubuntu:24.04

RUN apt-get update && apt-get install -y --no-install-recommends \
    librocksdb8.9 \
    ca-certificates \
    wget \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -r -s /bin/false indexer

WORKDIR /app
COPY --from=builder /indexer /app/indexer

RUN mkdir -p /app/data && chown indexer:indexer /app/data
VOLUME ["/app/data"]

USER indexer

ENV HTTP_LISTEN=:8080
ENV DATA_DIR=/app/data

EXPOSE 8080

HEALTHCHECK --interval=15s --timeout=3s --start-period=10s \
  CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/indexer"]
