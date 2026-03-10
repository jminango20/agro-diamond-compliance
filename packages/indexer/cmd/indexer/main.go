package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/api"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/config"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/indexer"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	state := indexer.NewState()
	idx := indexer.New(cfg, state)
	srv := api.New(state)

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

	// Start HTTP API
	httpSrv := &http.Server{
		Addr:    cfg.HTTPListenAddr,
		Handler: srv.Handler(),
	}

	go func() {
		log.Printf("[main] HTTP API listening on %s", cfg.HTTPListenAddr)
		if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[main] http server error: %v", err)
			cancel()
		}
	}()

	// Wait for shutdown signal
	select {
	case sig := <-sigCh:
		log.Printf("[main] received %s, shutting down", sig)
	case <-ctx.Done():
	}

	cancel()
	if err := httpSrv.Close(); err != nil {
		log.Printf("[main] http server close error: %v", err)
	}
}
