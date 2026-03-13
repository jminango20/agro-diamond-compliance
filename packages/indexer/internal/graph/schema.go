package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/renancorreadev/diamond-erc3643-protocol/packages/indexer/internal/store"
)

// NewSchema builds the GraphQL schema backed by a RocksDB store.
func NewSchema(db *store.Store) (graphql.Schema, error) {
	// ── Types ──────────────────────────────────────────────────

	holderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Holder",
		Fields: graphql.Fields{
			"address": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"balance": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	transferEventType := graphql.NewObject(graphql.ObjectConfig{
		Name: "TransferEvent",
		Fields: graphql.Fields{
			"txHash":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"block":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"logIndex":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"from":      &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"to":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"tokenId":   &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"amount":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"eventType": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	})

	tokenType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(*store.TokenMeta).TokenID, nil
				},
			},
			"totalSupply": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"holderCount": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"holders": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(holderType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 50,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					meta := p.Source.(*store.TokenMeta)
					holders, err := db.GetHolders(meta.TokenID)
					if err != nil {
						return nil, err
					}
					limit := p.Args["first"].(int)
					if limit > 0 && limit < len(holders) {
						holders = holders[:limit]
					}
					return holders, nil
				},
			},
			"events": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(transferEventType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 20,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					meta := p.Source.(*store.TokenMeta)
					limit := p.Args["first"].(int)
					events, err := db.GetTokenEvents(meta.TokenID, limit)
					if err != nil {
						return nil, err
					}
					return events, nil
				},
			},
		},
	})

	indexerStatusType := graphql.NewObject(graphql.ObjectConfig{
		Name: "IndexerStatus",
		Fields: graphql.Fields{
			"lastBlock":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"tokenCount": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		},
	})

	// ── Queries ────────────────────────────────────────────────

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"token": &graphql.Field{
				Type: tokenType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					meta, err := db.GetTokenMeta(id)
					if err != nil {
						return nil, err
					}
					if meta.TotalSupply == "0" && meta.HolderCount == 0 {
						return nil, nil
					}
					return meta, nil
				},
			},
			"tokens": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(tokenType))),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return db.GetAllTokens()
				},
			},
			"holder": &graphql.Field{
				Type: holderType,
				Args: graphql.FieldConfigArgument{
					"tokenId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"address": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					tokenID := p.Args["tokenId"].(string)
					addr := p.Args["address"].(string)
					bal, err := db.GetHolderBalance(tokenID, addr)
					if err != nil {
						return nil, err
					}
					return store.Holder{Address: addr, Balance: bal}, nil
				},
			},
			"events": &graphql.Field{
				Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(transferEventType))),
				Args: graphql.FieldConfigArgument{
					"first": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 50,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					limit := p.Args["first"].(int)
					return db.GetRecentEvents(limit)
				},
			},
			"status": &graphql.Field{
				Type: graphql.NewNonNull(indexerStatusType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cursor, _ := db.GetCursor()
					tokens, _ := db.GetAllTokens()
					return map[string]interface{}{
						"lastBlock":  cursor,
						"tokenCount": len(tokens),
					}, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
