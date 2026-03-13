package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	RPCURL          string
	RPCWSURL        string
	DiamondAddress  common.Address
	HTTPListenAddr  string
	StartBlock      uint64
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	rpc := os.Getenv("RPC_URL")
	if rpc == "" {
		rpc = "http://localhost:8545"
	}

	ws := os.Getenv("RPC_WS_URL")
	if ws == "" {
		ws = "ws://localhost:8545"
	}

	addr := os.Getenv("DIAMOND_ADDRESS")
	if addr == "" {
		return nil, fmt.Errorf("DIAMOND_ADDRESS env var is required")
	}

	listen := os.Getenv("HTTP_LISTEN")
	if listen == "" {
		listen = ":8080"
	}

	var startBlock uint64
	if sb := os.Getenv("START_BLOCK"); sb != "" {
		if v, err := strconv.ParseUint(sb, 10, 64); err == nil {
			startBlock = v
		}
	}

	return &Config{
		RPCURL:         rpc,
		RPCWSURL:       ws,
		DiamondAddress: common.HexToAddress(addr),
		HTTPListenAddr: listen,
		StartBlock:     startBlock,
	}, nil
}
