package indexer

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/config"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/store"
)

// Event signatures (keccak256 of the event signature string)
var (
	MintedSig         = crypto.Keccak256Hash([]byte("Minted(uint256,address,uint256)"))
	BurnedSig         = crypto.Keccak256Hash([]byte("Burned(uint256,address,uint256)"))
	ForcedTransferSig = crypto.Keccak256Hash([]byte("ForcedTransfer(uint256,address,address,uint256,bytes32)"))
	TransferSingleSig = crypto.Keccak256Hash([]byte("TransferSingle(address,address,address,uint256,uint256)"))

	AssetRegisteredSig    = crypto.Keccak256Hash([]byte("AssetRegistered(uint256,string,string,address)"))
	IdentityRegisteredSig = crypto.Keccak256Hash([]byte("IdentityRegistered(address,address,uint16)"))
	SnapshotCreatedSig    = crypto.Keccak256Hash([]byte("SnapshotCreated(uint256,uint256,uint256,uint64)"))
	DividendCreatedSig    = crypto.Keccak256Hash([]byte("DividendCreated(uint256,uint256,uint256,uint256,address)"))
	DividendClaimedSig    = crypto.Keccak256Hash([]byte("DividendClaimed(uint256,address,uint256)"))
)

// Indexer subscribes to Diamond contract events and persists them to RocksDB.
type Indexer struct {
	cfg   *config.Config
	store *store.Store
}

// New creates a new Indexer.
func New(cfg *config.Config, store *store.Store) *Indexer {
	return &Indexer{cfg: cfg, store: store}
}

// Run connects to the node via WebSocket and starts indexing events.
func (idx *Indexer) Run(ctx context.Context) error {
	client, err := ethclient.Dial(idx.cfg.RPCWSURL)
	if err != nil {
		return fmt.Errorf("failed to connect to node: %w", err)
	}
	defer client.Close()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{idx.cfg.DiamondAddress},
	}

	logsCh := make(chan types.Log, 256)
	sub, err := client.SubscribeFilterLogs(ctx, query, logsCh)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %w", err)
	}
	defer sub.Unsubscribe()

	log.Printf("[indexer] subscribed to events on %s", idx.cfg.DiamondAddress.Hex())

	for {
		select {
		case <-ctx.Done():
			log.Println("[indexer] shutting down")
			return nil

		case err := <-sub.Err():
			return fmt.Errorf("subscription error: %w", err)

		case vLog := <-logsCh:
			idx.handleLog(vLog)
		}
	}
}

func (idx *Indexer) handleLog(vLog types.Log) {
	if len(vLog.Topics) == 0 {
		return
	}

	sig := vLog.Topics[0]

	switch sig {
	case MintedSig:
		idx.handleMinted(vLog)
	case BurnedSig:
		idx.handleBurned(vLog)
	case ForcedTransferSig:
		idx.handleForcedTransfer(vLog)
	case TransferSingleSig:
		idx.handleTransferSingle(vLog)
	default:
		switch sig {
		case AssetRegisteredSig:
			log.Printf("[indexer] AssetRegistered block=%d", vLog.BlockNumber)
		case IdentityRegisteredSig:
			log.Printf("[indexer] IdentityRegistered block=%d", vLog.BlockNumber)
		case SnapshotCreatedSig:
			log.Printf("[indexer] SnapshotCreated block=%d", vLog.BlockNumber)
		case DividendCreatedSig:
			log.Printf("[indexer] DividendCreated block=%d", vLog.BlockNumber)
		case DividendClaimedSig:
			log.Printf("[indexer] DividendClaimed block=%d", vLog.BlockNumber)
		}
	}

	// Update cursor
	if err := idx.store.SetCursor(vLog.BlockNumber); err != nil {
		log.Printf("[indexer] failed to update cursor: %v", err)
	}
}

func (idx *Indexer) handleMinted(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	to := common.BytesToAddress(vLog.Topics[2].Bytes())
	amount := new(big.Int).SetBytes(vLog.Data)

	if err := idx.store.RecordMint(tokenID.String(), to, amount, vLog.TxHash, vLog.BlockNumber, vLog.Index); err != nil {
		log.Printf("[indexer] RecordMint error: %v", err)
		return
	}
	log.Printf("[indexer] Minted tokenId=%s to=%s amount=%s block=%d",
		tokenID, to.Hex(), amount, vLog.BlockNumber)
}

func (idx *Indexer) handleBurned(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	from := common.BytesToAddress(vLog.Topics[2].Bytes())
	amount := new(big.Int).SetBytes(vLog.Data)

	if err := idx.store.RecordBurn(tokenID.String(), from, amount, vLog.TxHash, vLog.BlockNumber, vLog.Index); err != nil {
		log.Printf("[indexer] RecordBurn error: %v", err)
		return
	}
	log.Printf("[indexer] Burned tokenId=%s from=%s amount=%s block=%d",
		tokenID, from.Hex(), amount, vLog.BlockNumber)
}

func (idx *Indexer) handleForcedTransfer(vLog types.Log) {
	if len(vLog.Topics) < 4 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	from := common.BytesToAddress(vLog.Topics[2].Bytes())
	to := common.BytesToAddress(vLog.Topics[3].Bytes())
	amount := new(big.Int).SetBytes(vLog.Data[:32])

	if err := idx.store.RecordTransfer(tokenID.String(), from, to, amount, "forced_transfer", vLog.TxHash, vLog.BlockNumber, vLog.Index); err != nil {
		log.Printf("[indexer] RecordTransfer error: %v", err)
		return
	}
	log.Printf("[indexer] ForcedTransfer tokenId=%s from=%s to=%s amount=%s block=%d",
		tokenID, from.Hex(), to.Hex(), amount, vLog.BlockNumber)
}

func (idx *Indexer) handleTransferSingle(vLog types.Log) {
	if len(vLog.Topics) < 4 || len(vLog.Data) < 64 {
		return
	}
	from := common.BytesToAddress(vLog.Topics[2].Bytes())
	to := common.BytesToAddress(vLog.Topics[3].Bytes())
	tokenID := new(big.Int).SetBytes(vLog.Data[:32])
	amount := new(big.Int).SetBytes(vLog.Data[32:64])

	zeroAddr := common.Address{}
	if from == zeroAddr || to == zeroAddr {
		return // skip mint/burn — handled by Minted/Burned
	}

	if err := idx.store.RecordTransfer(tokenID.String(), from, to, amount, "transfer", vLog.TxHash, vLog.BlockNumber, vLog.Index); err != nil {
		log.Printf("[indexer] RecordTransfer error: %v", err)
		return
	}
	log.Printf("[indexer] Transfer tokenId=%s from=%s to=%s amount=%s block=%d",
		tokenID, from.Hex(), to.Hex(), amount, vLog.BlockNumber)
}
