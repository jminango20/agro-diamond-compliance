package store

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/linxGnu/grocksdb"
)

// Key prefixes for RocksDB namespace isolation.
const (
	prefixTokenMeta    = "t:meta:"    // t:meta:{tokenId} → TokenMeta JSON
	prefixHolder       = "t:holder:"  // t:holder:{tokenId}:{address} → balance string
	prefixEvent        = "e:"         // e:{block}:{logIndex} → TransferEvent JSON
	prefixCursor       = "cursor"     // cursor → last indexed block (uint64 string)
	prefixGenericEvent = "ge:"        // ge:{block}:{logIndex} → GenericEvent JSON
	prefixAssetMeta    = "a:meta:"    // a:meta:{tokenId} → AssetConfig JSON
	prefixIdentity     = "i:"         // i:{wallet} → Identity JSON
	prefixIdentCountry = "i:country:" // i:country:{country}:{wallet} → wallet string
	prefixFreeze       = "f:"         // f:{wallet}:{tokenId|global} → FreezeState JSON
	prefixHolderAddr   = "h:addr:"    // h:addr:{address}:{tokenId} → balance string
	prefixGEToken      = "ge:token:"  // ge:token:{tokenId}:{block}:{logIndex} → key ref
	prefixGEAddr       = "ge:addr:"   // ge:addr:{address}:{block}:{logIndex} → key ref
)

// TokenMeta holds aggregate metadata for a tokenId.
type TokenMeta struct {
	TokenID     string `json:"tokenId"`
	TotalSupply string `json:"totalSupply"`
	HolderCount uint64 `json:"holderCount"`
}

// TransferEvent is a persisted event record.
type TransferEvent struct {
	TxHash    string `json:"txHash"`
	Block     uint64 `json:"block"`
	LogIndex  uint   `json:"logIndex"`
	From      string `json:"from"`
	To        string `json:"to"`
	TokenID   string `json:"tokenId"`
	Amount    string `json:"amount"`
	EventType string `json:"eventType"`
}

// Holder represents a holder with balance.
type Holder struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// GenericEvent captures any protocol event in a uniform shape.
type GenericEvent struct {
	TxHash    string `json:"txHash"`
	Block     uint64 `json:"block"`
	LogIndex  uint   `json:"logIndex"`
	EventType string `json:"eventType"`
	TokenID   string `json:"tokenId,omitempty"`
	Address   string `json:"address,omitempty"`
	Data      string `json:"data"`
}

// AssetConfig holds registered asset metadata.
type AssetConfig struct {
	TokenID      string `json:"tokenId"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Issuer       string `json:"issuer"`
	ProfileID    uint32 `json:"profileId"`
	URI          string `json:"uri,omitempty"`
	Paused       bool   `json:"paused"`
	RegisteredAt uint64 `json:"registeredAt"`
}

// Identity maps a wallet to its ONCHAINID and country.
type Identity struct {
	Wallet   string `json:"wallet"`
	Identity string `json:"identity"`
	Country  uint16 `json:"country"`
	BoundAt  uint64 `json:"boundAt"`
}

// FreezeState tracks freeze status for a wallet/asset pair.
type FreezeState struct {
	Wallet       string `json:"wallet"`
	TokenID      string `json:"tokenId,omitempty"`
	Frozen       bool   `json:"frozen"`
	FrozenAmount string `json:"frozenAmount,omitempty"`
	LockupExpiry uint64 `json:"lockupExpiry,omitempty"`
}

// Store wraps RocksDB for indexed state persistence.
type Store struct {
	db *grocksdb.DB
	ro *grocksdb.ReadOptions
	wo *grocksdb.WriteOptions
}

// New opens or creates a RocksDB database at the given path.
func New(path string) (*Store, error) {
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	opts.SetCompression(grocksdb.LZ4Compression)

	db, err := grocksdb.OpenDb(opts, path)
	if err != nil {
		return nil, fmt.Errorf("open rocksdb: %w", err)
	}

	return &Store{
		db: db,
		ro: grocksdb.NewDefaultReadOptions(),
		wo: grocksdb.NewDefaultWriteOptions(),
	}, nil
}

// Close releases RocksDB resources.
func (s *Store) Close() {
	s.ro.Destroy()
	s.wo.Destroy()
	s.db.Close()
}

/*//////////////////////////////////////////////////////////////
                        TOKEN META
//////////////////////////////////////////////////////////////*/

func tokenMetaKey(tokenID string) []byte {
	return []byte(prefixTokenMeta + tokenID)
}

func (s *Store) getTokenMeta(tokenID string) (*TokenMeta, error) {
	data, err := s.db.Get(s.ro, tokenMetaKey(tokenID))
	if err != nil {
		return nil, err
	}
	defer data.Free()

	if !data.Exists() {
		return &TokenMeta{TokenID: tokenID, TotalSupply: "0", HolderCount: 0}, nil
	}

	var meta TokenMeta
	if err := json.Unmarshal(data.Data(), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func (s *Store) putTokenMeta(meta *TokenMeta) error {
	data, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return s.db.Put(s.wo, tokenMetaKey(meta.TokenID), data)
}

// GetTokenMeta returns metadata for a tokenId.
func (s *Store) GetTokenMeta(tokenID string) (*TokenMeta, error) {
	return s.getTokenMeta(tokenID)
}

// GetAllTokens returns all token metadata using prefix scan.
func (s *Store) GetAllTokens() ([]*TokenMeta, error) {
	prefix := []byte(prefixTokenMeta)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var tokens []*TokenMeta
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		var meta TokenMeta
		if err := json.Unmarshal(it.Value().Data(), &meta); err != nil {
			continue
		}
		tokens = append(tokens, &meta)
	}
	return tokens, nil
}

/*//////////////////////////////////////////////////////////////
                        HOLDERS
//////////////////////////////////////////////////////////////*/

func holderKey(tokenID, address string) []byte {
	return []byte(prefixHolder + tokenID + ":" + address)
}

func holderPrefix(tokenID string) []byte {
	return []byte(prefixHolder + tokenID + ":")
}

func holderAddrKey(address, tokenID string) []byte {
	return []byte(prefixHolderAddr + address + ":" + tokenID)
}

func holderAddrPrefix(address string) []byte {
	return []byte(prefixHolderAddr + address + ":")
}

func (s *Store) getHolderBalance(tokenID, address string) (*big.Int, error) {
	data, err := s.db.Get(s.ro, holderKey(tokenID, address))
	if err != nil {
		return nil, err
	}
	defer data.Free()

	if !data.Exists() {
		return big.NewInt(0), nil
	}

	bal, ok := new(big.Int).SetString(string(data.Data()), 10)
	if !ok {
		return big.NewInt(0), nil
	}
	return bal, nil
}

func (s *Store) putHolderBalance(tokenID, address string, balance *big.Int) error {
	if err := s.db.Put(s.wo, holderKey(tokenID, address), []byte(balance.String())); err != nil {
		return err
	}
	return s.db.Put(s.wo, holderAddrKey(address, tokenID), []byte(balance.String()))
}

func (s *Store) deleteHolder(tokenID, address string) error {
	if err := s.db.Delete(s.wo, holderKey(tokenID, address)); err != nil {
		return err
	}
	return s.db.Delete(s.wo, holderAddrKey(address, tokenID))
}

// GetHolders returns all holders for a tokenId with their balances.
func (s *Store) GetHolders(tokenID string) ([]Holder, error) {
	prefix := holderPrefix(tokenID)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var holders []Holder
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		key := string(it.Key().Data())
		// Extract address from key: t:holder:{tokenId}:{address}
		addr := key[len(string(prefix)):]
		holders = append(holders, Holder{
			Address: addr,
			Balance: string(it.Value().Data()),
		})
	}
	return holders, nil
}

// GetHolderBalance returns the balance of a specific holder.
func (s *Store) GetHolderBalance(tokenID, address string) (string, error) {
	bal, err := s.getHolderBalance(tokenID, address)
	if err != nil {
		return "0", err
	}
	return bal.String(), nil
}

// GetPortfolio returns all token holdings for a given address.
func (s *Store) GetPortfolio(address string) ([]Holder, error) {
	prefix := holderAddrPrefix(address)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var holdings []Holder
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		key := string(it.Key().Data())
		tokenID := key[len(string(prefix)):]
		holdings = append(holdings, Holder{
			Address: tokenID,
			Balance: string(it.Value().Data()),
		})
	}
	return holdings, nil
}

/*//////////////////////////////////////////////////////////////
                        EVENTS
//////////////////////////////////////////////////////////////*/

func eventKey(block uint64, logIndex uint) []byte {
	return []byte(fmt.Sprintf("%s%012d:%06d", prefixEvent, block, logIndex))
}

func (s *Store) putEvent(evt *TransferEvent) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return s.db.Put(s.wo, eventKey(evt.Block, evt.LogIndex), data)
}

// GetRecentEvents returns the last N events (reverse iteration).
func (s *Store) GetRecentEvents(limit int) ([]TransferEvent, error) {
	prefix := []byte(prefixEvent)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	// Seek to end of prefix range
	endPrefix := []byte("e;") // 'e;' > 'e:' but < 'f'
	it.Seek(endPrefix)

	// Move back to last event
	if it.Valid() {
		it.Prev()
	} else {
		it.SeekToLast()
	}

	var events []TransferEvent
	for ; it.Valid() && len(events) < limit; it.Prev() {
		key := it.Key().Data()
		if len(key) < len(prefix) || string(key[:len(prefix)]) != string(prefix) {
			break
		}
		var evt TransferEvent
		if err := json.Unmarshal(it.Value().Data(), &evt); err != nil {
			continue
		}
		events = append(events, evt)
	}

	// Reverse to chronological order
	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events, nil
}

// GetTokenEvents returns events for a specific tokenId (scans all events).
func (s *Store) GetTokenEvents(tokenID string, limit int) ([]TransferEvent, error) {
	events, err := s.GetRecentEvents(limit * 5) // overscan then filter
	if err != nil {
		return nil, err
	}

	var filtered []TransferEvent
	for _, evt := range events {
		if evt.TokenID == tokenID {
			filtered = append(filtered, evt)
			if len(filtered) >= limit {
				break
			}
		}
	}
	return filtered, nil
}

/*//////////////////////////////////////////////////////////////
                    GENERIC EVENTS
//////////////////////////////////////////////////////////////*/

func genericEventKey(block uint64, logIndex uint) []byte {
	return []byte(fmt.Sprintf("%s%012d:%06d", prefixGenericEvent, block, logIndex))
}

func geTokenKey(tokenID string, block uint64, logIndex uint) []byte {
	return []byte(fmt.Sprintf("%s%s:%012d:%06d", prefixGEToken, tokenID, block, logIndex))
}

func geAddrKey(address string, block uint64, logIndex uint) []byte {
	return []byte(fmt.Sprintf("%s%s:%012d:%06d", prefixGEAddr, address, block, logIndex))
}

func (s *Store) putGenericEvent(evt *GenericEvent) error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	key := genericEventKey(evt.Block, evt.LogIndex)
	if err := s.db.Put(s.wo, key, data); err != nil {
		return err
	}
	// Secondary indexes
	if evt.TokenID != "" {
		if err := s.db.Put(s.wo, geTokenKey(evt.TokenID, evt.Block, evt.LogIndex), key); err != nil {
			return err
		}
	}
	if evt.Address != "" {
		if err := s.db.Put(s.wo, geAddrKey(evt.Address, evt.Block, evt.LogIndex), key); err != nil {
			return err
		}
	}
	return nil
}

// RecordGenericEvent persists a generic protocol event with optional secondary indexes.
func (s *Store) RecordGenericEvent(evt *GenericEvent) error {
	return s.putGenericEvent(evt)
}

// GetProtocolEvents returns generic events with optional filters.
func (s *Store) GetProtocolEvents(limit int, eventType, tokenID, address string) ([]GenericEvent, error) {
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	// Choose the best index based on filters
	if tokenID != "" {
		return s.getProtocolEventsByIndex(it, []byte(prefixGEToken+tokenID+":"), limit, eventType)
	}
	if address != "" {
		return s.getProtocolEventsByIndex(it, []byte(prefixGEAddr+address+":"), limit, eventType)
	}

	// Full scan of ge: prefix (reverse for most recent)
	prefix := []byte(prefixGenericEvent)
	endPrefix := []byte("ge;")
	it.Seek(endPrefix)
	if it.Valid() {
		it.Prev()
	} else {
		it.SeekToLast()
	}

	var events []GenericEvent
	for ; it.Valid() && len(events) < limit; it.Prev() {
		key := it.Key().Data()
		if len(key) < len(prefix) || string(key[:len(prefix)]) != string(prefix) {
			break
		}
		var evt GenericEvent
		if err := json.Unmarshal(it.Value().Data(), &evt); err != nil {
			continue
		}
		if eventType != "" && evt.EventType != eventType {
			continue
		}
		events = append(events, evt)
	}

	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events, nil
}

func (s *Store) getProtocolEventsByIndex(it *grocksdb.Iterator, prefix []byte, limit int, eventType string) ([]GenericEvent, error) {
	// Find the end of the prefix range for reverse iteration
	endPrefix := make([]byte, len(prefix))
	copy(endPrefix, prefix)
	endPrefix[len(endPrefix)-1]++ // increment last byte
	it.Seek(endPrefix)
	if it.Valid() {
		it.Prev()
	} else {
		it.SeekToLast()
	}

	var events []GenericEvent
	for ; it.Valid() && len(events) < limit; it.Prev() {
		key := it.Key().Data()
		if len(key) < len(prefix) || string(key[:len(prefix)]) != string(prefix) {
			break
		}
		// Value is the primary key — fetch actual event
		primaryKey := it.Value().Data()
		data, err := s.db.Get(s.ro, primaryKey)
		if err != nil {
			continue
		}
		if !data.Exists() {
			data.Free()
			continue
		}
		var evt GenericEvent
		if err := json.Unmarshal(data.Data(), &evt); err != nil {
			data.Free()
			continue
		}
		data.Free()
		if eventType != "" && evt.EventType != eventType {
			continue
		}
		events = append(events, evt)
	}

	for i, j := 0, len(events)-1; i < j; i, j = i+1, j-1 {
		events[i], events[j] = events[j], events[i]
	}
	return events, nil
}

/*//////////////////////////////////////////////////////////////
                    ASSET CONFIG
//////////////////////////////////////////////////////////////*/

func assetMetaKey(tokenID string) []byte {
	return []byte(prefixAssetMeta + tokenID)
}

// RecordAsset creates or updates an AssetConfig.
func (s *Store) RecordAsset(asset *AssetConfig) error {
	data, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return s.db.Put(s.wo, assetMetaKey(asset.TokenID), data)
}

// GetAsset returns the AssetConfig for a tokenId.
func (s *Store) GetAsset(tokenID string) (*AssetConfig, error) {
	data, err := s.db.Get(s.ro, assetMetaKey(tokenID))
	if err != nil {
		return nil, err
	}
	defer data.Free()
	if !data.Exists() {
		return nil, nil
	}
	var asset AssetConfig
	if err := json.Unmarshal(data.Data(), &asset); err != nil {
		return nil, err
	}
	return &asset, nil
}

// GetAllAssets returns all registered asset configs.
func (s *Store) GetAllAssets() ([]*AssetConfig, error) {
	prefix := []byte(prefixAssetMeta)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var assets []*AssetConfig
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		var asset AssetConfig
		if err := json.Unmarshal(it.Value().Data(), &asset); err != nil {
			continue
		}
		assets = append(assets, &asset)
	}
	return assets, nil
}

// UpdateAssetPaused sets the paused flag for a given tokenId.
func (s *Store) UpdateAssetPaused(tokenID string, paused bool) error {
	asset, err := s.GetAsset(tokenID)
	if err != nil {
		return err
	}
	if asset == nil {
		asset = &AssetConfig{TokenID: tokenID}
	}
	asset.Paused = paused
	return s.RecordAsset(asset)
}

// UpdateAssetURI sets the URI for a given tokenId.
func (s *Store) UpdateAssetURI(tokenID, uri string) error {
	asset, err := s.GetAsset(tokenID)
	if err != nil {
		return err
	}
	if asset == nil {
		asset = &AssetConfig{TokenID: tokenID}
	}
	asset.URI = uri
	return s.RecordAsset(asset)
}

/*//////////////////////////////////////////////////////////////
                    IDENTITY
//////////////////////////////////////////////////////////////*/

func identityKey(wallet string) []byte {
	return []byte(prefixIdentity + wallet)
}

func identityCountryKey(country uint16, wallet string) []byte {
	return []byte(fmt.Sprintf("%s%d:%s", prefixIdentCountry, country, wallet))
}

// RecordIdentity creates or updates an identity binding.
func (s *Store) RecordIdentity(id *Identity) error {
	data, err := json.Marshal(id)
	if err != nil {
		return err
	}
	if err := s.db.Put(s.wo, identityKey(id.Wallet), data); err != nil {
		return err
	}
	return s.db.Put(s.wo, identityCountryKey(id.Country, id.Wallet), []byte(id.Wallet))
}

// DeleteIdentity removes an identity binding.
func (s *Store) DeleteIdentity(wallet string) error {
	// Read existing to clean up country index
	existing, err := s.GetIdentity(wallet)
	if err != nil {
		return err
	}
	if existing != nil {
		_ = s.db.Delete(s.wo, identityCountryKey(existing.Country, wallet))
	}
	return s.db.Delete(s.wo, identityKey(wallet))
}

// GetIdentity returns the identity for a wallet.
func (s *Store) GetIdentity(wallet string) (*Identity, error) {
	data, err := s.db.Get(s.ro, identityKey(wallet))
	if err != nil {
		return nil, err
	}
	defer data.Free()
	if !data.Exists() {
		return nil, nil
	}
	var id Identity
	if err := json.Unmarshal(data.Data(), &id); err != nil {
		return nil, err
	}
	return &id, nil
}

// GetIdentitiesByCountry returns all identities for a given country code.
func (s *Store) GetIdentitiesByCountry(country uint16, limit int) ([]*Identity, error) {
	prefix := []byte(fmt.Sprintf("%s%d:", prefixIdentCountry, country))
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var identities []*Identity
	for it.Seek(prefix); it.ValidForPrefix(prefix) && len(identities) < limit; it.Next() {
		wallet := string(it.Value().Data())
		id, err := s.GetIdentity(wallet)
		if err != nil || id == nil {
			continue
		}
		identities = append(identities, id)
	}
	return identities, nil
}

// GetAllIdentities returns all identities up to a limit.
func (s *Store) GetAllIdentities(limit int) ([]*Identity, error) {
	prefix := []byte(prefixIdentity)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var identities []*Identity
	for it.Seek(prefix); it.ValidForPrefix(prefix) && len(identities) < limit; it.Next() {
		key := string(it.Key().Data())
		// Skip country index keys (i:country:...)
		if strings.HasPrefix(key, prefixIdentCountry) {
			continue
		}
		var id Identity
		if err := json.Unmarshal(it.Value().Data(), &id); err != nil {
			continue
		}
		identities = append(identities, &id)
	}
	return identities, nil
}

/*//////////////////////////////////////////////////////////////
                    FREEZE STATE
//////////////////////////////////////////////////////////////*/

func freezeKey(wallet, tokenID string) []byte {
	suffix := "global"
	if tokenID != "" {
		suffix = tokenID
	}
	return []byte(prefixFreeze + wallet + ":" + suffix)
}

func freezeWalletPrefix(wallet string) []byte {
	return []byte(prefixFreeze + wallet + ":")
}

// RecordFreeze creates or updates a freeze state.
func (s *Store) RecordFreeze(state *FreezeState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return s.db.Put(s.wo, freezeKey(state.Wallet, state.TokenID), data)
}

// GetFreeze returns the freeze state for a specific wallet/tokenId pair.
func (s *Store) GetFreeze(wallet, tokenID string) (*FreezeState, error) {
	data, err := s.db.Get(s.ro, freezeKey(wallet, tokenID))
	if err != nil {
		return nil, err
	}
	defer data.Free()
	if !data.Exists() {
		return nil, nil
	}
	var state FreezeState
	if err := json.Unmarshal(data.Data(), &state); err != nil {
		return nil, err
	}
	return &state, nil
}

// GetFreezes returns all freeze states for a wallet.
func (s *Store) GetFreezes(wallet string) ([]*FreezeState, error) {
	prefix := freezeWalletPrefix(wallet)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var states []*FreezeState
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		var state FreezeState
		if err := json.Unmarshal(it.Value().Data(), &state); err != nil {
			continue
		}
		states = append(states, &state)
	}
	return states, nil
}

// GetFrozenWallets scans all freeze states, optionally filtering by tokenId.
func (s *Store) GetFrozenWallets(tokenID string) ([]*FreezeState, error) {
	prefix := []byte(prefixFreeze)
	it := s.db.NewIterator(s.ro)
	defer it.Close()

	var states []*FreezeState
	for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
		var state FreezeState
		if err := json.Unmarshal(it.Value().Data(), &state); err != nil {
			continue
		}
		if !state.Frozen {
			continue
		}
		if tokenID != "" && state.TokenID != tokenID {
			continue
		}
		states = append(states, &state)
	}
	return states, nil
}

// UpdateFreezeAmount updates the frozen amount for a wallet/tokenId pair.
func (s *Store) UpdateFreezeAmount(wallet, tokenID, amount string) error {
	state, err := s.GetFreeze(wallet, tokenID)
	if err != nil {
		return err
	}
	if state == nil {
		state = &FreezeState{Wallet: wallet, TokenID: tokenID}
	}
	state.FrozenAmount = amount
	return s.RecordFreeze(state)
}

// UpdateLockupExpiry updates the lockup expiry for a wallet/tokenId pair.
func (s *Store) UpdateLockupExpiry(wallet, tokenID string, expiry uint64) error {
	state, err := s.GetFreeze(wallet, tokenID)
	if err != nil {
		return err
	}
	if state == nil {
		state = &FreezeState{Wallet: wallet, TokenID: tokenID}
	}
	state.LockupExpiry = expiry
	return s.RecordFreeze(state)
}

/*//////////////////////////////////////////////////////////////
                        CURSOR
//////////////////////////////////////////////////////////////*/

// GetCursor returns the last indexed block number.
func (s *Store) GetCursor() (uint64, error) {
	data, err := s.db.Get(s.ro, []byte(prefixCursor))
	if err != nil {
		return 0, err
	}
	defer data.Free()

	if !data.Exists() {
		return 0, nil
	}

	var block uint64
	_, err = fmt.Sscanf(string(data.Data()), "%d", &block)
	return block, err
}

// SetCursor saves the last indexed block number.
func (s *Store) SetCursor(block uint64) error {
	return s.db.Put(s.wo, []byte(prefixCursor), []byte(fmt.Sprintf("%d", block)))
}

/*//////////////////////////////////////////////////////////////
                    STATE MUTATION (called by indexer)
//////////////////////////////////////////////////////////////*/

// RecordMint updates token supply and holder balance for a mint event.
func (s *Store) RecordMint(tokenID string, to common.Address, amount *big.Int, txHash common.Hash, block uint64, logIndex uint) error {
	meta, err := s.getTokenMeta(tokenID)
	if err != nil {
		return err
	}

	supply, _ := new(big.Int).SetString(meta.TotalSupply, 10)
	supply.Add(supply, amount)
	meta.TotalSupply = supply.String()

	addr := to.Hex()
	bal, err := s.getHolderBalance(tokenID, addr)
	if err != nil {
		return err
	}

	wasZero := bal.Sign() == 0
	bal.Add(bal, amount)

	if err := s.putHolderBalance(tokenID, addr, bal); err != nil {
		return err
	}

	if wasZero {
		meta.HolderCount++
	}

	if err := s.putTokenMeta(meta); err != nil {
		return err
	}

	return s.putEvent(&TransferEvent{
		TxHash:    txHash.Hex(),
		Block:     block,
		LogIndex:  logIndex,
		From:      common.Address{}.Hex(),
		To:        addr,
		TokenID:   tokenID,
		Amount:    amount.String(),
		EventType: "mint",
	})
}

// RecordBurn updates token supply and holder balance for a burn event.
func (s *Store) RecordBurn(tokenID string, from common.Address, amount *big.Int, txHash common.Hash, block uint64, logIndex uint) error {
	meta, err := s.getTokenMeta(tokenID)
	if err != nil {
		return err
	}

	supply, _ := new(big.Int).SetString(meta.TotalSupply, 10)
	supply.Sub(supply, amount)
	meta.TotalSupply = supply.String()

	addr := from.Hex()
	bal, err := s.getHolderBalance(tokenID, addr)
	if err != nil {
		return err
	}

	bal.Sub(bal, amount)

	if bal.Sign() <= 0 {
		if err := s.deleteHolder(tokenID, addr); err != nil {
			return err
		}
		meta.HolderCount--
	} else {
		if err := s.putHolderBalance(tokenID, addr, bal); err != nil {
			return err
		}
	}

	if err := s.putTokenMeta(meta); err != nil {
		return err
	}

	return s.putEvent(&TransferEvent{
		TxHash:    txHash.Hex(),
		Block:     block,
		LogIndex:  logIndex,
		From:      addr,
		To:        common.Address{}.Hex(),
		TokenID:   tokenID,
		Amount:    amount.String(),
		EventType: "burn",
	})
}

// RecordTransfer updates holder balances for a transfer event.
func (s *Store) RecordTransfer(tokenID string, from, to common.Address, amount *big.Int, eventType string, txHash common.Hash, block uint64, logIndex uint) error {
	meta, err := s.getTokenMeta(tokenID)
	if err != nil {
		return err
	}

	fromAddr := from.Hex()
	toAddr := to.Hex()

	// Debit from
	fromBal, err := s.getHolderBalance(tokenID, fromAddr)
	if err != nil {
		return err
	}
	fromBal.Sub(fromBal, amount)

	if fromBal.Sign() <= 0 {
		if err := s.deleteHolder(tokenID, fromAddr); err != nil {
			return err
		}
		meta.HolderCount--
	} else {
		if err := s.putHolderBalance(tokenID, fromAddr, fromBal); err != nil {
			return err
		}
	}

	// Credit to
	toBal, err := s.getHolderBalance(tokenID, toAddr)
	if err != nil {
		return err
	}
	wasZero := toBal.Sign() == 0
	toBal.Add(toBal, amount)

	if err := s.putHolderBalance(tokenID, toAddr, toBal); err != nil {
		return err
	}
	if wasZero {
		meta.HolderCount++
	}

	if err := s.putTokenMeta(meta); err != nil {
		return err
	}

	return s.putEvent(&TransferEvent{
		TxHash:    txHash.Hex(),
		Block:     block,
		LogIndex:  logIndex,
		From:      fromAddr,
		To:        toAddr,
		TokenID:   tokenID,
		Amount:    amount.String(),
		EventType: eventType,
	})
}

// RecordEventOnly persists a transfer event without modifying balances.
// Used for ForcedTransfer which emits alongside TransferSingle (that already handles balances).
func (s *Store) RecordEventOnly(tokenID string, from, to common.Address, amount *big.Int, eventType string, txHash common.Hash, block uint64, logIndex uint) error {
	return s.putEvent(&TransferEvent{
		TxHash:    txHash.Hex(),
		Block:     block,
		LogIndex:  logIndex,
		From:      from.Hex(),
		To:        to.Hex(),
		TokenID:   tokenID,
		Amount:    amount.String(),
		EventType: eventType,
	})
}
