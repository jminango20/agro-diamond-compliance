package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/config"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/graph"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/indexer"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/store"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// Open RocksDB
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}
	db, err := store.New(dataDir)
	if err != nil {
		log.Fatalf("rocksdb: %v", err)
	}
	defer db.Close()

	cursor, _ := db.GetCursor()
	log.Printf("[main] RocksDB opened at %s (last block: %d)", dataDir, cursor)

	// Build GraphQL schema
	schema, err := graph.NewSchema(db)
	if err != nil {
		log.Fatalf("graphql schema: %v", err)
	}

	// Setup indexer
	idx := indexer.New(cfg, db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Start indexer in background
	go func() {
		if err := idx.Run(ctx); err != nil {
			log.Printf("[main] indexer stopped: %v", err)
			cancel()
		}
	}()

	// Start HTTP server (GraphQL + Playground)
	mux := http.NewServeMux()
	mux.Handle("POST /graphql", graph.Handler(schema))
	mux.Handle("GET /", graph.PlaygroundHandler())
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cursor, _ := db.GetCursor()
		w.Write([]byte(`{"status":"ok","lastBlock":` + formatUint(cursor) + `}`))
	})

	httpSrv := &http.Server{
		Addr:    cfg.HTTPListenAddr,
		Handler: mux,
	}

	go func() {
		log.Printf("[main] GraphQL playground at http://localhost%s", cfg.HTTPListenAddr)
		log.Printf("[main] GraphQL endpoint at http://localhost%s/graphql", cfg.HTTPListenAddr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[main] http server error: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown
	select {
	case sig := <-sigCh:
		log.Printf("[main] received %s, shutting down", sig)
	case <-ctx.Done():
	}

	cancel()
	if err := httpSrv.Close(); err != nil {
		log.Printf("[main] http close error: %v", err)
	}
}

func formatUint(n uint64) string {
	return fmt.Sprintf("%d", n)
}
