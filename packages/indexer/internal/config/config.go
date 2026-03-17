package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Config holds runtime configuration loaded from environment variables.
type Config struct {
	RPCURL         string
	RPCWSURL       string
	DiamondAddress common.Address
	ExtraAddresses []common.Address
	HTTPListenAddr string
	StartBlock     uint64
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

	var extra []common.Address
	if ea := os.Getenv("EXTRA_ADDRESSES"); ea != "" {
		for _, a := range strings.Split(ea, ",") {
			a = strings.TrimSpace(a)
			if a != "" {
				extra = append(extra, common.HexToAddress(a))
			}
		}
	}

	return &Config{
		RPCURL:         rpc,
		RPCWSURL:       ws,
		DiamondAddress: common.HexToAddress(addr),
		ExtraAddresses: extra,
		HTTPListenAddr: listen,
		StartBlock:     startBlock,
	}, nil
}

// AllAddresses returns Diamond + ExtraAddresses for FilterQuery.
func (c *Config) AllAddresses() []common.Address {
	addrs := make([]common.Address, 0, 1+len(c.ExtraAddresses))
	addrs = append(addrs, c.DiamondAddress)
	addrs = append(addrs, c.ExtraAddresses...)
	return addrs
}
