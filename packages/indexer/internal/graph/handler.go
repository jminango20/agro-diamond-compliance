package graph

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Handler returns an http.Handler that serves GraphQL queries.
func Handler(schema graphql.Schema) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		var req struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  req.Query,
			OperationName:  req.OperationName,
			VariableValues: req.Variables,
		})

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("[graphql] encode error: %v", err)
		}
	})
}

// PlaygroundHandler returns an HTTP handler serving GraphQL Playground HTML.
func PlaygroundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
  <title>Diamond ERC-3643 — GraphQL Playground</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.26/build/static/css/index.css"/>
  <script src="https://cdn.jsdelivr.net/npm/graphql-playground-react@1.7.26/build/static/js/middleware.js"></script>
</head>
<body>
  <div id="root"></div>
  <script>
    window.addEventListener('load', function () {
      GraphQLPlayground.init(document.getElementById('root'), { endpoint: '/graphql' })
    })
  </script>
</body>
</html>`))
	})
}
