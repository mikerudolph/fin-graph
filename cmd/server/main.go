package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/mikerudolph/fin-graph/pkg/config"
	"github.com/mikerudolph/fin-graph/pkg/gql"
)

type reqBody struct {
	Query string `json:"query"`
}

func main() {
	conf := config.GetConfig()

	rootQuery := gql.NewRoot(conf)
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: rootQuery.Query,
		},
	)

	// start http handler
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var rBody reqBody
		_ = json.NewDecoder(r.Body).Decode(&rBody)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: rBody.Query,
		})

		// get result
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)

}
