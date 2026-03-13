package store

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/linxGnu/grocksdb"
)

// Key prefixes for RocksDB namespace isolation.
const (
	prefixTokenMeta   = "t:meta:"    // t:meta:{tokenId} → TokenMeta JSON
	prefixHolder      = "t:holder:"  // t:holder:{tokenId}:{address} → balance string
	prefixEvent       = "e:"         // e:{block}:{logIndex} → TransferEvent JSON
	prefixCursor      = "cursor"     // cursor → last indexed block (uint64 string)
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
	return s.db.Put(s.wo, holderKey(tokenID, address), []byte(balance.String()))
}

func (s *Store) deleteHolder(tokenID, address string) error {
	return s.db.Delete(s.wo, holderKey(tokenID, address))
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
	endPrefix := []byte("f:") // 'f' > 'e', so this is past all events
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
