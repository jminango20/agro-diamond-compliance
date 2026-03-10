package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/indexer"
)

// Server exposes indexed state via HTTP.
type Server struct {
	state *indexer.State
	mux   *http.ServeMux
}

// New creates an API server.
func New(state *indexer.State) *Server {
	s := &Server{
		state: state,
		mux:   http.NewServeMux(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /health", s.handleHealth)
	s.mux.HandleFunc("GET /api/tokens", s.handleGetTokens)
	s.mux.HandleFunc("GET /api/tokens/{tokenId}", s.handleGetToken)
	s.mux.HandleFunc("GET /api/tokens/{tokenId}/holders", s.handleGetHolders)
	s.mux.HandleFunc("GET /api/events", s.handleGetEvents)
}

// Handler returns the HTTP handler.
func (s *Server) Handler() http.Handler {
	return s.mux
}

func (s *Server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status":    "ok",
		"lastBlock": strconv.FormatUint(s.state.LastBlock, 10),
	})
}

func (s *Server) handleGetTokens(w http.ResponseWriter, _ *http.Request) {
	tokens := s.state.GetAllTokens()

	type tokenSummary struct {
		TokenID     string `json:"tokenId"`
		TotalSupply string `json:"totalSupply"`
		HolderCount uint64 `json:"holderCount"`
	}

	result := make([]tokenSummary, 0, len(tokens))
	for _, t := range tokens {
		result = append(result, tokenSummary{
			TokenID:     t.TokenID.String(),
			TotalSupply: t.TotalSupply.String(),
			HolderCount: t.HolderCount,
		})
	}
	writeJSON(w, http.StatusOK, result)
}

func (s *Server) handleGetToken(w http.ResponseWriter, r *http.Request) {
	tokenID := r.PathValue("tokenId")
	t, ok := s.state.GetTokenState(tokenID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "token not found"})
		return
	}
	writeJSON(w, http.StatusOK, t)
}

func (s *Server) handleGetHolders(w http.ResponseWriter, r *http.Request) {
	tokenID := r.PathValue("tokenId")
	t, ok := s.state.GetTokenState(tokenID)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "token not found"})
		return
	}

	type holderEntry struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}

	holders := make([]holderEntry, 0, len(t.Holders))
	for addr, bal := range t.Holders {
		holders = append(holders, holderEntry{
			Address: addr,
			Balance: bal.String(),
		})
	}
	writeJSON(w, http.StatusOK, holders)
}

func (s *Server) handleGetEvents(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit := 50
	if limitStr != "" {
		if n, err := strconv.Atoi(limitStr); err == nil && n > 0 {
			limit = n
		}
	}
	events := s.state.GetEvents(limit)
	writeJSON(w, http.StatusOK, events)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("[api] json encode error: %v", err)
	}
}
