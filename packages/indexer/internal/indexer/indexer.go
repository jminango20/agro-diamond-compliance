package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

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
	// Token events (existing)
	MintedSig         = crypto.Keccak256Hash([]byte("Minted(uint256,address,uint256)"))
	BurnedSig         = crypto.Keccak256Hash([]byte("Burned(uint256,address,uint256)"))
	ForcedTransferSig = crypto.Keccak256Hash([]byte("ForcedTransfer(uint256,address,address,uint256,bytes32)"))
	TransferSingleSig = crypto.Keccak256Hash([]byte("TransferSingle(address,address,address,uint256,uint256)"))

	// Asset management
	AssetRegisteredSig      = crypto.Keccak256Hash([]byte("AssetRegistered(uint256,address,uint32)"))
	AssetConfigUpdatedSig   = crypto.Keccak256Hash([]byte("AssetConfigUpdated(uint256)"))
	ComplianceModuleAddSig  = crypto.Keccak256Hash([]byte("ComplianceModuleAdded(uint256,address)"))
	ComplianceModuleRemSig  = crypto.Keccak256Hash([]byte("ComplianceModuleRemoved(uint256,address)"))
	URISig                  = crypto.Keccak256Hash([]byte("URI(string,uint256)"))

	// Identity
	IdentityBoundSig   = crypto.Keccak256Hash([]byte("IdentityBound(address,address,uint16)"))
	IdentityUnboundSig = crypto.Keccak256Hash([]byte("IdentityUnbound(address)"))

	// Freeze & lockup
	WalletFrozenSig  = crypto.Keccak256Hash([]byte("WalletFrozen(address,bool)"))
	AssetFrozenSig   = crypto.Keccak256Hash([]byte("AssetFrozen(uint256,address,bool)"))
	PartialFreezeSig = crypto.Keccak256Hash([]byte("PartialFreeze(uint256,address,uint256)"))
	LockupSetSig     = crypto.Keccak256Hash([]byte("LockupSet(uint256,address,uint64)"))

	// Access control
	RoleGrantedSig = crypto.Keccak256Hash([]byte("RoleGranted(bytes32,address,address)"))
	RoleRevokedSig = crypto.Keccak256Hash([]byte("RoleRevoked(bytes32,address,address)"))

	// Pause
	EmergencyPauseSig  = crypto.Keccak256Hash([]byte("EmergencyPause(address)"))
	ProtocolUnpausedSig = crypto.Keccak256Hash([]byte("ProtocolUnpaused(address)"))
	AssetPausedSig     = crypto.Keccak256Hash([]byte("AssetPaused(uint256,address)"))
	AssetUnpausedSig   = crypto.Keccak256Hash([]byte("AssetUnpaused(uint256,address)"))

	// Recovery
	WalletRecoveredSig = crypto.Keccak256Hash([]byte("WalletRecovered(address,address,address)"))

	// Snapshots & dividends
	SnapshotCreatedSig = crypto.Keccak256Hash([]byte("SnapshotCreated(uint256,uint256,uint256,uint64)"))
	DividendCreatedSig = crypto.Keccak256Hash([]byte("DividendCreated(uint256,uint256,uint256,uint256,address)"))
	DividendClaimedSig = crypto.Keccak256Hash([]byte("DividendClaimed(uint256,address,uint256)"))

	// Asset groups
	GroupCreatedSig = crypto.Keccak256Hash([]byte("GroupCreated(uint256,uint256,string,uint256)"))
	UnitMintedSig   = crypto.Keccak256Hash([]byte("UnitMinted(uint256,uint256,address,string,uint256)"))

	// Plugin system — GlobalPluginFacet
	GlobalPluginRegisteredSig    = crypto.Keccak256Hash([]byte("GlobalPluginRegistered(address,string)"))
	GlobalPluginRemovedSig       = crypto.Keccak256Hash([]byte("GlobalPluginRemoved(address)"))
	GlobalPluginStatusChangedSig = crypto.Keccak256Hash([]byte("GlobalPluginStatusChanged(address,bool)"))

	// Plugin system — AssetManagerFacet (per-asset hookable plugins)
	PluginModuleAddedSig   = crypto.Keccak256Hash([]byte("PluginModuleAdded(uint256,address)"))
	PluginModuleRemovedSig = crypto.Keccak256Hash([]byte("PluginModuleRemoved(uint256,address)"))
	PluginModulesSetSig    = crypto.Keccak256Hash([]byte("PluginModulesSet(uint256,address[])"))

	// YieldDistributorModule (external contract)
	RewardAssetAddedSig   = crypto.Keccak256Hash([]byte("RewardAssetAdded(uint256,address,uint256,uint8)"))
	RewardAssetRemovedSig = crypto.Keccak256Hash([]byte("RewardAssetRemoved(uint256,address,uint256,uint8)"))
	YieldDepositedSig     = crypto.Keccak256Hash([]byte("YieldDeposited(uint256,bytes32,uint256)"))
	YieldClaimedSig       = crypto.Keccak256Hash([]byte("YieldClaimed(uint256,bytes32,address,uint256)"))

	// Compliance modules (external contracts)
	CountryRestrictedSig   = crypto.Keccak256Hash([]byte("CountryRestricted(uint256,uint16)"))
	CountryUnrestrictedSig = crypto.Keccak256Hash([]byte("CountryUnrestricted(uint256,uint16)"))
	MaxBalanceSetSig       = crypto.Keccak256Hash([]byte("MaxBalanceSet(uint256,uint256)"))
	MaxHoldersSetSig       = crypto.Keccak256Hash([]byte("MaxHoldersSet(uint256,uint256)"))
)

// Indexer subscribes to Diamond contract events and persists them to RocksDB.
type Indexer struct {
	cfg        *config.Config
	store      *store.Store
	httpClient *ethclient.Client // for on-chain enrichment calls
}

// New creates a new Indexer.
func New(cfg *config.Config, store *store.Store) *Indexer {
	return &Indexer{cfg: cfg, store: store}
}

// Run connects to the node, backfills historical events, then subscribes to new ones.
func (idx *Indexer) Run(ctx context.Context) error {
	// Open persistent HTTP client for on-chain enrichment (getAssetConfig, etc.)
	httpClient, err := ethclient.Dial(idx.cfg.RPCURL)
	if err != nil {
		return fmt.Errorf("failed to connect HTTP client: %w", err)
	}
	defer httpClient.Close()
	idx.httpClient = httpClient

	// Enrich existing assets that are missing name/symbol
	idx.enrichAssets()

	// Backfill historical events via HTTP RPC
	if err := idx.backfill(ctx); err != nil {
		log.Printf("[indexer] backfill warning: %v (continuing with live subscription)", err)
	}

	// Poll for new events via HTTP (no WebSocket dependency)
	log.Printf("[indexer] polling for new events every 5s on %d address(es)", len(idx.cfg.AllAddresses()))

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[indexer] shutting down")
			return nil

		case <-ticker.C:
			idx.pollNewBlocks(ctx)
		}
	}
}

// backfill fetches historical logs from StartBlock (or last cursor) to latest block.
func (idx *Indexer) backfill(ctx context.Context) error {
	httpClient := idx.httpClient

	// Determine start block: max(StartBlock, cursor+1)
	fromBlock := idx.cfg.StartBlock
	if cursor, err := idx.store.GetCursor(); err == nil && cursor > 0 {
		if cursor+1 > fromBlock {
			fromBlock = cursor + 1
		}
	}

	if fromBlock == 0 {
		log.Println("[indexer] backfill: no START_BLOCK configured and no cursor, skipping backfill")
		return nil
	}

	latest, err := httpClient.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("backfill: failed to get latest block: %w", err)
	}

	if fromBlock > latest {
		log.Printf("[indexer] backfill: already up to date (cursor=%d, latest=%d)", fromBlock-1, latest)
		return nil
	}

	log.Printf("[indexer] backfill: fetching logs from block %d to %d", fromBlock, latest)

	// Smaller chunks + rate limit delay to avoid 429 from providers like Infura
	const chunkSize uint64 = 1000
	const delayBetweenChunks = 500 * time.Millisecond
	const maxRetries = 5
	totalProcessed := 0

	for from := fromBlock; from <= latest; from += chunkSize + 1 {
		to := from + chunkSize
		if to > latest {
			to = latest
		}

		query := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(to),
			Addresses: idx.cfg.AllAddresses(),
		}

		var logs []types.Log
		for attempt := 0; attempt <= maxRetries; attempt++ {
			logs, err = httpClient.FilterLogs(ctx, query)
			if err == nil {
				break
			}

			if ctx.Err() != nil {
				return ctx.Err()
			}

			// Exponential backoff: 1s, 2s, 4s, 8s, 16s
			backoff := time.Duration(1<<uint(attempt)) * time.Second
			log.Printf("[indexer] backfill: 429/error on blocks %d-%d, retrying in %v (attempt %d/%d)", from, to, backoff, attempt+1, maxRetries)
			time.Sleep(backoff)
		}
		if err != nil {
			return fmt.Errorf("backfill: FilterLogs(%d-%d) failed after %d retries: %w", from, to, maxRetries, err)
		}

		for _, vLog := range logs {
			idx.handleLog(vLog)
			totalProcessed++
		}

		if len(logs) > 0 {
			log.Printf("[indexer] backfill: processed %d events in blocks %d-%d", len(logs), from, to)
		}

		// Rate limit: small delay between chunks
		time.Sleep(delayBetweenChunks)
	}

	log.Printf("[indexer] backfill: complete — %d events indexed", totalProcessed)
	return nil
}

// pollNewBlocks fetches logs from the last cursor to the latest block.
func (idx *Indexer) pollNewBlocks(ctx context.Context) {
	latest, err := idx.httpClient.BlockNumber(ctx)
	if err != nil {
		log.Printf("[indexer] poll: failed to get latest block: %v", err)
		return
	}

	var fromBlock uint64
	if cursor, err := idx.store.GetCursor(); err == nil && cursor > 0 {
		fromBlock = cursor + 1
	} else {
		return
	}

	if fromBlock > latest {
		return
	}

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(latest),
		Addresses: idx.cfg.AllAddresses(),
	}

	logs, err := idx.httpClient.FilterLogs(ctx, query)
	if err != nil {
		log.Printf("[indexer] poll: FilterLogs(%d-%d) error: %v", fromBlock, latest, err)
		return
	}

	for _, vLog := range logs {
		idx.handleLog(vLog)
	}

	// Update cursor even if no logs (to advance past empty blocks)
	if err := idx.store.SetCursor(latest); err != nil {
		log.Printf("[indexer] poll: failed to update cursor: %v", err)
	}

	if len(logs) > 0 {
		log.Printf("[indexer] poll: %d new events in blocks %d-%d", len(logs), fromBlock, latest)
	}
}

func (idx *Indexer) handleLog(vLog types.Log) {
	if len(vLog.Topics) == 0 {
		return
	}

	sig := vLog.Topics[0]

	switch sig {
	// Token events (state-mutating)
	case MintedSig:
		idx.handleMinted(vLog)
	case BurnedSig:
		idx.handleBurned(vLog)
	case ForcedTransferSig:
		idx.handleForcedTransfer(vLog)
	case TransferSingleSig:
		idx.handleTransferSingle(vLog)

	// Asset management
	case AssetRegisteredSig:
		idx.handleAssetRegistered(vLog)
	case AssetConfigUpdatedSig:
		idx.handleGenericEvent(vLog, "asset_config_updated")
	case ComplianceModuleAddSig:
		idx.handleGenericEventWithTokenAndAddr(vLog, "compliance_module_added")
	case ComplianceModuleRemSig:
		idx.handleGenericEventWithTokenAndAddr(vLog, "compliance_module_removed")
	case URISig:
		idx.handleURI(vLog)

	// Identity
	case IdentityBoundSig:
		idx.handleIdentityBound(vLog)
	case IdentityUnboundSig:
		idx.handleIdentityUnbound(vLog)

	// Freeze & lockup
	case WalletFrozenSig:
		idx.handleWalletFrozen(vLog)
	case AssetFrozenSig:
		idx.handleAssetFrozen(vLog)
	case PartialFreezeSig:
		idx.handlePartialFreeze(vLog)
	case LockupSetSig:
		idx.handleLockupSet(vLog)

	// Access control
	case RoleGrantedSig:
		idx.handleRoleEvent(vLog, "role_granted")
	case RoleRevokedSig:
		idx.handleRoleEvent(vLog, "role_revoked")

	// Pause
	case EmergencyPauseSig:
		idx.handleGenericEventAddrOnly(vLog, "emergency_pause")
	case ProtocolUnpausedSig:
		idx.handleGenericEventAddrOnly(vLog, "protocol_unpaused")
	case AssetPausedSig:
		idx.handleAssetPaused(vLog, true)
	case AssetUnpausedSig:
		idx.handleAssetPaused(vLog, false)

	// Recovery
	case WalletRecoveredSig:
		idx.handleWalletRecovered(vLog)

	// Snapshots & dividends
	case SnapshotCreatedSig:
		idx.handleSnapshotCreated(vLog)
	case DividendCreatedSig:
		idx.handleDividendCreated(vLog)
	case DividendClaimedSig:
		idx.handleDividendClaimed(vLog)

	// Asset groups
	case GroupCreatedSig:
		idx.handleGenericEvent(vLog, "group_created")
	case UnitMintedSig:
		idx.handleGenericEvent(vLog, "unit_minted")

	// Plugin system — GlobalPluginFacet
	case GlobalPluginRegisteredSig:
		idx.handleGlobalPluginRegistered(vLog)
	case GlobalPluginRemovedSig:
		idx.handleGenericEventAddrOnly(vLog, "global_plugin_removed")
	case GlobalPluginStatusChangedSig:
		idx.handleGlobalPluginStatusChanged(vLog)

	// Plugin system — AssetManagerFacet (per-asset hookable)
	case PluginModuleAddedSig:
		idx.handleGenericEventWithTokenAndAddr(vLog, "plugin_module_added")
	case PluginModuleRemovedSig:
		idx.handleGenericEventWithTokenAndAddr(vLog, "plugin_module_removed")
	case PluginModulesSetSig:
		idx.handleGenericEvent(vLog, "plugin_modules_set")

	// YieldDistributorModule (external contract)
	case RewardAssetAddedSig:
		idx.handleRewardAssetEvent(vLog, "reward_asset_added")
	case RewardAssetRemovedSig:
		idx.handleRewardAssetEvent(vLog, "reward_asset_removed")
	case YieldDepositedSig:
		idx.handleYieldDeposited(vLog)
	case YieldClaimedSig:
		idx.handleYieldClaimed(vLog)

	// Compliance modules (external contracts)
	case CountryRestrictedSig:
		idx.handleCountryRestriction(vLog, "country_restricted")
	case CountryUnrestrictedSig:
		idx.handleCountryRestriction(vLog, "country_unrestricted")
	case MaxBalanceSetSig:
		idx.handleMaxSetting(vLog, "max_balance_set", "maxBalance")
	case MaxHoldersSetSig:
		idx.handleMaxSetting(vLog, "max_holders_set", "maxHolders")
	}

	// Update cursor
	if err := idx.store.SetCursor(vLog.BlockNumber); err != nil {
		log.Printf("[indexer] failed to update cursor: %v", err)
	}
}

/*//////////////////////////////////////////////////////////////
                    EXISTING HANDLERS
//////////////////////////////////////////////////////////////*/

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

	// Only record the event — balance updates are handled by the paired TransferSingle event
	if err := idx.store.RecordEventOnly(tokenID.String(), from, to, amount, "forced_transfer", vLog.TxHash, vLog.BlockNumber, vLog.Index); err != nil {
		log.Printf("[indexer] RecordEventOnly error: %v", err)
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

/*//////////////////////////////////////////////////////////////
                    ASSET MANAGEMENT HANDLERS
//////////////////////////////////////////////////////////////*/

// AssetRegistered(uint256 indexed tokenId, address indexed issuer, uint32 profileId)
func (idx *Indexer) handleAssetRegistered(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	issuer := common.BytesToAddress(vLog.Topics[2].Bytes())

	var profileID uint32
	if len(vLog.Data) >= 32 {
		profileID = uint32(new(big.Int).SetBytes(vLog.Data[:32]).Uint64())
	}

	asset := &store.AssetConfig{
		TokenID:      tokenID.String(),
		Issuer:       issuer.Hex(),
		ProfileID:    profileID,
		RegisteredAt: vLog.BlockNumber,
	}

	// Enrich with on-chain name/symbol via getAssetConfig()
	if name, symbol, err := idx.fetchAssetNameSymbol(tokenID); err == nil {
		asset.Name = name
		asset.Symbol = symbol
	}

	if err := idx.store.RecordAsset(asset); err != nil {
		log.Printf("[indexer] RecordAsset error: %v", err)
		return
	}

	idx.recordGenericEvent(vLog, "asset_registered", tokenID.String(), issuer.Hex(), map[string]interface{}{
		"profileId": profileID,
		"name":      asset.Name,
		"symbol":    asset.Symbol,
	})

	log.Printf("[indexer] AssetRegistered tokenId=%s name=%s symbol=%s issuer=%s block=%d",
		tokenID, asset.Name, asset.Symbol, issuer.Hex(), vLog.BlockNumber)
}

// fetchAssetNameSymbol calls name(uint256) and symbol(uint256) on the Diamond.
func (idx *Indexer) fetchAssetNameSymbol(tokenID *big.Int) (string, string, error) {
	if idx.httpClient == nil {
		return "", "", fmt.Errorf("no HTTP client")
	}

	name, err := idx.callStringView("name(uint256)", tokenID)
	if err != nil {
		return "", "", fmt.Errorf("name: %w", err)
	}

	symbol, err := idx.callStringView("symbol(uint256)", tokenID)
	if err != nil {
		return name, "", fmt.Errorf("symbol: %w", err)
	}

	return name, symbol, nil
}

// callStringView calls a view function that returns a single string.
func (idx *Indexer) callStringView(sig string, tokenID *big.Int) (string, error) {
	selector := crypto.Keccak256([]byte(sig))[:4]
	data := make([]byte, 36)
	copy(data[:4], selector)
	tokenID.FillBytes(data[4:36])

	msg := ethereum.CallMsg{
		To:   &idx.cfg.DiamondAddress,
		Data: data,
	}

	result, err := idx.httpClient.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	// ABI-encoded string: [offset (32 bytes)] [length (32 bytes)] [string data]
	if len(result) < 64 {
		return "", fmt.Errorf("result too short: %d bytes", len(result))
	}

	offset := new(big.Int).SetBytes(result[:32]).Uint64()
	if offset+32 > uint64(len(result)) {
		return "", fmt.Errorf("offset out of bounds")
	}

	length := new(big.Int).SetBytes(result[offset : offset+32]).Uint64()
	if offset+32+length > uint64(len(result)) {
		return "", fmt.Errorf("string data out of bounds")
	}

	return string(result[offset+32 : offset+32+length]), nil
}

// enrichAssets fills in missing name/symbol for existing AssetConfig entries
// by calling getAssetConfig() on the Diamond contract.
func (idx *Indexer) enrichAssets() {
	assets, err := idx.store.GetAllAssets()
	if err != nil || len(assets) == 0 {
		return
	}

	enriched := 0
	for _, asset := range assets {
		if asset.Name != "" && asset.Symbol != "" {
			continue
		}
		tokenID := new(big.Int)
		tokenID.SetString(asset.TokenID, 10)

		name, symbol, err := idx.fetchAssetNameSymbol(tokenID)
		if err != nil {
			log.Printf("[indexer] enrich: failed to fetch name/symbol for tokenId=%s: %v", asset.TokenID, err)
			continue
		}

		asset.Name = name
		asset.Symbol = symbol
		if err := idx.store.RecordAsset(asset); err != nil {
			log.Printf("[indexer] enrich: failed to update asset %s: %v", asset.TokenID, err)
			continue
		}
		enriched++
	}

	if enriched > 0 {
		log.Printf("[indexer] enriched %d assets with name/symbol from on-chain", enriched)
	}
}

// URI(string value, uint256 indexed id)
func (idx *Indexer) handleURI(vLog types.Log) {
	if len(vLog.Topics) < 2 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	// URI string is ABI-encoded in data (offset + length + bytes)
	var uri string
	if len(vLog.Data) >= 64 {
		offset := new(big.Int).SetBytes(vLog.Data[:32]).Uint64()
		if offset+32 <= uint64(len(vLog.Data)) {
			length := new(big.Int).SetBytes(vLog.Data[offset : offset+32]).Uint64()
			if offset+32+length <= uint64(len(vLog.Data)) {
				uri = string(vLog.Data[offset+32 : offset+32+length])
			}
		}
	}

	if err := idx.store.UpdateAssetURI(tokenID.String(), uri); err != nil {
		log.Printf("[indexer] UpdateAssetURI error: %v", err)
	}

	idx.recordGenericEvent(vLog, "uri_updated", tokenID.String(), "", map[string]interface{}{
		"uri": uri,
	})

	log.Printf("[indexer] URI tokenId=%s uri=%s block=%d", tokenID, uri, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    IDENTITY HANDLERS
//////////////////////////////////////////////////////////////*/

// IdentityBound(address indexed wallet, address indexed identity, uint16 country)
func (idx *Indexer) handleIdentityBound(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	wallet := common.BytesToAddress(vLog.Topics[1].Bytes())
	identity := common.BytesToAddress(vLog.Topics[2].Bytes())

	var country uint16
	if len(vLog.Data) >= 32 {
		country = uint16(new(big.Int).SetBytes(vLog.Data[:32]).Uint64())
	}

	id := &store.Identity{
		Wallet:   wallet.Hex(),
		Identity: identity.Hex(),
		Country:  country,
		BoundAt:  vLog.BlockNumber,
	}

	if err := idx.store.RecordIdentity(id); err != nil {
		log.Printf("[indexer] RecordIdentity error: %v", err)
		return
	}

	idx.recordGenericEvent(vLog, "identity_bound", "", wallet.Hex(), map[string]interface{}{
		"identity": identity.Hex(),
		"country":  country,
	})

	log.Printf("[indexer] IdentityBound wallet=%s identity=%s country=%d block=%d",
		wallet.Hex(), identity.Hex(), country, vLog.BlockNumber)
}

// IdentityUnbound(address indexed wallet)
func (idx *Indexer) handleIdentityUnbound(vLog types.Log) {
	if len(vLog.Topics) < 2 {
		return
	}
	wallet := common.BytesToAddress(vLog.Topics[1].Bytes())

	if err := idx.store.DeleteIdentity(wallet.Hex()); err != nil {
		log.Printf("[indexer] DeleteIdentity error: %v", err)
	}

	idx.recordGenericEvent(vLog, "identity_unbound", "", wallet.Hex(), nil)

	log.Printf("[indexer] IdentityUnbound wallet=%s block=%d", wallet.Hex(), vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    FREEZE & LOCKUP HANDLERS
//////////////////////////////////////////////////////////////*/

// WalletFrozen(address indexed wallet, bool frozen)
func (idx *Indexer) handleWalletFrozen(vLog types.Log) {
	if len(vLog.Topics) < 2 {
		return
	}
	wallet := common.BytesToAddress(vLog.Topics[1].Bytes())

	var frozen bool
	if len(vLog.Data) >= 32 {
		frozen = new(big.Int).SetBytes(vLog.Data[:32]).Sign() != 0
	}

	state := &store.FreezeState{
		Wallet: wallet.Hex(),
		Frozen: frozen,
	}

	if err := idx.store.RecordFreeze(state); err != nil {
		log.Printf("[indexer] RecordFreeze error: %v", err)
	}

	idx.recordGenericEvent(vLog, "wallet_frozen", "", wallet.Hex(), map[string]interface{}{
		"frozen": frozen,
	})

	log.Printf("[indexer] WalletFrozen wallet=%s frozen=%t block=%d", wallet.Hex(), frozen, vLog.BlockNumber)
}

// AssetFrozen(uint256 indexed tokenId, address indexed wallet, bool frozen)
func (idx *Indexer) handleAssetFrozen(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	wallet := common.BytesToAddress(vLog.Topics[2].Bytes())

	var frozen bool
	if len(vLog.Data) >= 32 {
		frozen = new(big.Int).SetBytes(vLog.Data[:32]).Sign() != 0
	}

	state := &store.FreezeState{
		Wallet:  wallet.Hex(),
		TokenID: tokenID.String(),
		Frozen:  frozen,
	}

	if err := idx.store.RecordFreeze(state); err != nil {
		log.Printf("[indexer] RecordFreeze error: %v", err)
	}

	idx.recordGenericEvent(vLog, "asset_frozen", tokenID.String(), wallet.Hex(), map[string]interface{}{
		"frozen": frozen,
	})

	log.Printf("[indexer] AssetFrozen tokenId=%s wallet=%s frozen=%t block=%d",
		tokenID, wallet.Hex(), frozen, vLog.BlockNumber)
}

// PartialFreeze(uint256 indexed tokenId, address indexed wallet, uint256 amount)
func (idx *Indexer) handlePartialFreeze(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	wallet := common.BytesToAddress(vLog.Topics[2].Bytes())

	var amount string
	if len(vLog.Data) >= 32 {
		amount = new(big.Int).SetBytes(vLog.Data[:32]).String()
	}

	if err := idx.store.UpdateFreezeAmount(wallet.Hex(), tokenID.String(), amount); err != nil {
		log.Printf("[indexer] UpdateFreezeAmount error: %v", err)
	}

	idx.recordGenericEvent(vLog, "partial_freeze", tokenID.String(), wallet.Hex(), map[string]interface{}{
		"amount": amount,
	})

	log.Printf("[indexer] PartialFreeze tokenId=%s wallet=%s amount=%s block=%d",
		tokenID, wallet.Hex(), amount, vLog.BlockNumber)
}

// LockupSet(uint256 indexed tokenId, address indexed wallet, uint64 expiry)
func (idx *Indexer) handleLockupSet(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	wallet := common.BytesToAddress(vLog.Topics[2].Bytes())

	var expiry uint64
	if len(vLog.Data) >= 32 {
		expiry = new(big.Int).SetBytes(vLog.Data[:32]).Uint64()
	}

	if err := idx.store.UpdateLockupExpiry(wallet.Hex(), tokenID.String(), expiry); err != nil {
		log.Printf("[indexer] UpdateLockupExpiry error: %v", err)
	}

	idx.recordGenericEvent(vLog, "lockup_set", tokenID.String(), wallet.Hex(), map[string]interface{}{
		"expiry": expiry,
	})

	log.Printf("[indexer] LockupSet tokenId=%s wallet=%s expiry=%d block=%d",
		tokenID, wallet.Hex(), expiry, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    PAUSE HANDLERS
//////////////////////////////////////////////////////////////*/

// AssetPaused(uint256 indexed tokenId, address indexed by) / AssetUnpaused
func (idx *Indexer) handleAssetPaused(vLog types.Log, paused bool) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	by := common.BytesToAddress(vLog.Topics[2].Bytes())

	if err := idx.store.UpdateAssetPaused(tokenID.String(), paused); err != nil {
		log.Printf("[indexer] UpdateAssetPaused error: %v", err)
	}

	eventType := "asset_paused"
	if !paused {
		eventType = "asset_unpaused"
	}

	idx.recordGenericEvent(vLog, eventType, tokenID.String(), by.Hex(), nil)

	log.Printf("[indexer] %s tokenId=%s by=%s block=%d",
		eventType, tokenID, by.Hex(), vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    ACCESS CONTROL HANDLERS
//////////////////////////////////////////////////////////////*/

// RoleGranted/RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (idx *Indexer) handleRoleEvent(vLog types.Log, eventType string) {
	if len(vLog.Topics) < 4 {
		return
	}
	role := vLog.Topics[1]
	account := common.BytesToAddress(vLog.Topics[2].Bytes())
	sender := common.BytesToAddress(vLog.Topics[3].Bytes())

	idx.recordGenericEvent(vLog, eventType, "", account.Hex(), map[string]interface{}{
		"role":   role.Hex(),
		"sender": sender.Hex(),
	})

	log.Printf("[indexer] %s role=%s account=%s sender=%s block=%d",
		eventType, role.Hex()[:10], account.Hex(), sender.Hex(), vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    RECOVERY HANDLER
//////////////////////////////////////////////////////////////*/

// WalletRecovered(address indexed lostWallet, address indexed newWallet, address indexed agent)
func (idx *Indexer) handleWalletRecovered(vLog types.Log) {
	if len(vLog.Topics) < 4 {
		return
	}
	lostWallet := common.BytesToAddress(vLog.Topics[1].Bytes())
	newWallet := common.BytesToAddress(vLog.Topics[2].Bytes())
	agent := common.BytesToAddress(vLog.Topics[3].Bytes())

	// Update identity: rebind from lost to new wallet
	existing, err := idx.store.GetIdentity(lostWallet.Hex())
	if err == nil && existing != nil {
		newId := &store.Identity{
			Wallet:   newWallet.Hex(),
			Identity: existing.Identity,
			Country:  existing.Country,
			BoundAt:  vLog.BlockNumber,
		}
		_ = idx.store.RecordIdentity(newId)
		_ = idx.store.DeleteIdentity(lostWallet.Hex())
	}

	idx.recordGenericEvent(vLog, "wallet_recovered", "", newWallet.Hex(), map[string]interface{}{
		"lostWallet": lostWallet.Hex(),
		"agent":      agent.Hex(),
	})

	log.Printf("[indexer] WalletRecovered lost=%s new=%s agent=%s block=%d",
		lostWallet.Hex(), newWallet.Hex(), agent.Hex(), vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    SNAPSHOT & DIVIDEND HANDLERS
//////////////////////////////////////////////////////////////*/

// SnapshotCreated(uint256 indexed snapshotId, uint256 indexed tokenId, uint256 totalSupply, uint64 timestamp)
func (idx *Indexer) handleSnapshotCreated(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	snapshotID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(vLog.Topics[2].Bytes())

	var totalSupply string
	var timestamp uint64
	if len(vLog.Data) >= 64 {
		totalSupply = new(big.Int).SetBytes(vLog.Data[:32]).String()
		timestamp = new(big.Int).SetBytes(vLog.Data[32:64]).Uint64()
	}

	idx.recordGenericEvent(vLog, "snapshot_created", tokenID.String(), "", map[string]interface{}{
		"snapshotId":  snapshotID.String(),
		"totalSupply": totalSupply,
		"timestamp":   timestamp,
	})

	log.Printf("[indexer] SnapshotCreated snapshotId=%s tokenId=%s block=%d",
		snapshotID, tokenID, vLog.BlockNumber)
}

// DividendCreated(uint256 indexed dividendId, uint256 indexed tokenId, uint256 indexed snapshotId, uint256 totalAmount, address paymentToken)
func (idx *Indexer) handleDividendCreated(vLog types.Log) {
	if len(vLog.Topics) < 4 {
		return
	}
	dividendID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	tokenID := new(big.Int).SetBytes(vLog.Topics[2].Bytes())
	snapshotID := new(big.Int).SetBytes(vLog.Topics[3].Bytes())

	var totalAmount string
	var paymentToken string
	if len(vLog.Data) >= 64 {
		totalAmount = new(big.Int).SetBytes(vLog.Data[:32]).String()
		paymentToken = common.BytesToAddress(vLog.Data[32:64]).Hex()
	}

	idx.recordGenericEvent(vLog, "dividend_created", tokenID.String(), "", map[string]interface{}{
		"dividendId":   dividendID.String(),
		"snapshotId":   snapshotID.String(),
		"totalAmount":  totalAmount,
		"paymentToken": paymentToken,
	})

	log.Printf("[indexer] DividendCreated dividendId=%s tokenId=%s block=%d",
		dividendID, tokenID, vLog.BlockNumber)
}

// DividendClaimed(uint256 indexed dividendId, address indexed holder, uint256 amount)
func (idx *Indexer) handleDividendClaimed(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	dividendID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	holder := common.BytesToAddress(vLog.Topics[2].Bytes())

	var amount string
	if len(vLog.Data) >= 32 {
		amount = new(big.Int).SetBytes(vLog.Data[:32]).String()
	}

	idx.recordGenericEvent(vLog, "dividend_claimed", "", holder.Hex(), map[string]interface{}{
		"dividendId": dividendID.String(),
		"amount":     amount,
	})

	log.Printf("[indexer] DividendClaimed dividendId=%s holder=%s amount=%s block=%d",
		dividendID, holder.Hex(), amount, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    GENERIC EVENT HELPERS
//////////////////////////////////////////////////////////////*/

func (idx *Indexer) recordGenericEvent(vLog types.Log, eventType, tokenID, address string, extra map[string]interface{}) {
	dataJSON := "{}"
	if extra != nil {
		if b, err := json.Marshal(extra); err == nil {
			dataJSON = string(b)
		}
	}

	evt := &store.GenericEvent{
		TxHash:    vLog.TxHash.Hex(),
		Block:     vLog.BlockNumber,
		LogIndex:  vLog.Index,
		EventType: eventType,
		TokenID:   tokenID,
		Address:   address,
		Data:      dataJSON,
	}

	if err := idx.store.RecordGenericEvent(evt); err != nil {
		log.Printf("[indexer] RecordGenericEvent(%s) error: %v", eventType, err)
	}
}

// handleGenericEvent records an event with just the tokenId from topic[1].
func (idx *Indexer) handleGenericEvent(vLog types.Log, eventType string) {
	var tokenID string
	if len(vLog.Topics) >= 2 {
		tokenID = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).String()
	}
	idx.recordGenericEvent(vLog, eventType, tokenID, "", nil)
	log.Printf("[indexer] %s block=%d", eventType, vLog.BlockNumber)
}

// handleGenericEventWithTokenAndAddr records with tokenId from topic[1] and address from topic[2].
func (idx *Indexer) handleGenericEventWithTokenAndAddr(vLog types.Log, eventType string) {
	var tokenID, addr string
	if len(vLog.Topics) >= 2 {
		tokenID = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).String()
	}
	if len(vLog.Topics) >= 3 {
		addr = common.BytesToAddress(vLog.Topics[2].Bytes()).Hex()
	}
	idx.recordGenericEvent(vLog, eventType, tokenID, addr, map[string]interface{}{
		"module": addr,
	})
	log.Printf("[indexer] %s tokenId=%s module=%s block=%d", eventType, tokenID, addr, vLog.BlockNumber)
}

// handleGenericEventAddrOnly records with address from topic[1].
func (idx *Indexer) handleGenericEventAddrOnly(vLog types.Log, eventType string) {
	var addr string
	if len(vLog.Topics) >= 2 {
		addr = common.BytesToAddress(vLog.Topics[1].Bytes()).Hex()
	}
	idx.recordGenericEvent(vLog, eventType, "", addr, nil)
	log.Printf("[indexer] %s by=%s block=%d", eventType, addr, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    PLUGIN SYSTEM HANDLERS
//////////////////////////////////////////////////////////////*/

// GlobalPluginRegistered(address indexed plugin, string name)
func (idx *Indexer) handleGlobalPluginRegistered(vLog types.Log) {
	if len(vLog.Topics) < 2 {
		return
	}
	plugin := common.BytesToAddress(vLog.Topics[1].Bytes())

	// name is ABI-encoded string in data
	var pluginName string
	if len(vLog.Data) >= 64 {
		offset := new(big.Int).SetBytes(vLog.Data[:32]).Uint64()
		if offset+32 <= uint64(len(vLog.Data)) {
			length := new(big.Int).SetBytes(vLog.Data[offset : offset+32]).Uint64()
			if offset+32+length <= uint64(len(vLog.Data)) {
				pluginName = string(vLog.Data[offset+32 : offset+32+length])
			}
		}
	}

	idx.recordGenericEvent(vLog, "global_plugin_registered", "", plugin.Hex(), map[string]interface{}{
		"plugin": plugin.Hex(),
		"name":   pluginName,
	})

	log.Printf("[indexer] GlobalPluginRegistered plugin=%s name=%s block=%d",
		plugin.Hex(), pluginName, vLog.BlockNumber)
}

// GlobalPluginStatusChanged(address indexed plugin, bool active)
func (idx *Indexer) handleGlobalPluginStatusChanged(vLog types.Log) {
	if len(vLog.Topics) < 2 {
		return
	}
	plugin := common.BytesToAddress(vLog.Topics[1].Bytes())

	var active bool
	if len(vLog.Data) >= 32 {
		active = new(big.Int).SetBytes(vLog.Data[:32]).Sign() != 0
	}

	idx.recordGenericEvent(vLog, "global_plugin_status_changed", "", plugin.Hex(), map[string]interface{}{
		"plugin": plugin.Hex(),
		"active": active,
	})

	log.Printf("[indexer] GlobalPluginStatusChanged plugin=%s active=%t block=%d",
		plugin.Hex(), active, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    YIELD DISTRIBUTOR HANDLERS
//////////////////////////////////////////////////////////////*/

// RewardAssetAdded/Removed(uint256 indexed tokenId, address indexed token, uint256 id, RewardType assetType)
func (idx *Indexer) handleRewardAssetEvent(vLog types.Log, eventType string) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	rewardToken := common.BytesToAddress(vLog.Topics[2].Bytes())

	var rewardID string
	var assetType string
	if len(vLog.Data) >= 64 {
		rewardID = new(big.Int).SetBytes(vLog.Data[:32]).String()
		typeVal := new(big.Int).SetBytes(vLog.Data[32:64]).Uint64()
		if typeVal == 0 {
			assetType = "ERC20"
		} else {
			assetType = "ERC1155"
		}
	}

	idx.recordGenericEvent(vLog, eventType, tokenID.String(), rewardToken.Hex(), map[string]interface{}{
		"token":     rewardToken.Hex(),
		"id":        rewardID,
		"assetType": assetType,
	})

	log.Printf("[indexer] %s tokenId=%s rewardToken=%s type=%s block=%d",
		eventType, tokenID, rewardToken.Hex(), assetType, vLog.BlockNumber)
}

// YieldDeposited(uint256 indexed tokenId, bytes32 indexed rewardKey, uint256 amount)
func (idx *Indexer) handleYieldDeposited(vLog types.Log) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	rewardKey := vLog.Topics[2]

	var amount string
	if len(vLog.Data) >= 32 {
		amount = new(big.Int).SetBytes(vLog.Data[:32]).String()
	}

	idx.recordGenericEvent(vLog, "yield_deposited", tokenID.String(), "", map[string]interface{}{
		"rewardKey": rewardKey.Hex(),
		"amount":    amount,
	})

	log.Printf("[indexer] YieldDeposited tokenId=%s rewardKey=%s amount=%s block=%d",
		tokenID, rewardKey.Hex()[:10], amount, vLog.BlockNumber)
}

// YieldClaimed(uint256 indexed tokenId, bytes32 indexed rewardKey, address indexed holder, uint256 amount)
func (idx *Indexer) handleYieldClaimed(vLog types.Log) {
	if len(vLog.Topics) < 4 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	rewardKey := vLog.Topics[2]
	holder := common.BytesToAddress(vLog.Topics[3].Bytes())

	var amount string
	if len(vLog.Data) >= 32 {
		amount = new(big.Int).SetBytes(vLog.Data[:32]).String()
	}

	idx.recordGenericEvent(vLog, "yield_claimed", tokenID.String(), holder.Hex(), map[string]interface{}{
		"rewardKey": rewardKey.Hex(),
		"holder":    holder.Hex(),
		"amount":    amount,
	})

	log.Printf("[indexer] YieldClaimed tokenId=%s holder=%s amount=%s block=%d",
		tokenID, holder.Hex(), amount, vLog.BlockNumber)
}

/*//////////////////////////////////////////////////////////////
                    COMPLIANCE MODULE HANDLERS
//////////////////////////////////////////////////////////////*/

// CountryRestricted/CountryUnrestricted(uint256 indexed tokenId, uint16 indexed country)
func (idx *Indexer) handleCountryRestriction(vLog types.Log, eventType string) {
	if len(vLog.Topics) < 3 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())
	country := uint16(new(big.Int).SetBytes(vLog.Topics[2].Bytes()).Uint64())

	idx.recordGenericEvent(vLog, eventType, tokenID.String(), "", map[string]interface{}{
		"country": country,
	})

	log.Printf("[indexer] %s tokenId=%s country=%d block=%d",
		eventType, tokenID, country, vLog.BlockNumber)
}

// MaxBalanceSet/MaxHoldersSet(uint256 indexed tokenId, uint256 value)
func (idx *Indexer) handleMaxSetting(vLog types.Log, eventType, fieldName string) {
	if len(vLog.Topics) < 2 {
		return
	}
	tokenID := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	var value string
	if len(vLog.Data) >= 32 {
		value = new(big.Int).SetBytes(vLog.Data[:32]).String()
	}

	idx.recordGenericEvent(vLog, eventType, tokenID.String(), "", map[string]interface{}{
		fieldName: value,
	})

	log.Printf("[indexer] %s tokenId=%s %s=%s block=%d",
		eventType, tokenID, fieldName, value, vLog.BlockNumber)
}
