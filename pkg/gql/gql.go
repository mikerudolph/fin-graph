package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/mikerudolph/fin-graph/pkg/config"
)

// Root ... I am root
type Root struct {
	Query *graphql.Object
}

// NewRoot returns a new root
func NewRoot(conf config.Config) *Root {
	resolver := NewResolver(conf)

	// define a type service that will return types with associated resolve functionality
	typeService := TypeService{
		resolverService: resolver,
	}

	// define a root query object with top level queries that can be executed
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"company": &graphql.Field{
						Type: typeService.Company(),
						Args: graphql.FieldConfigArgument{
							"symbol": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.CompanyResolver,
					},

					"price": &graphql.Field{
						Type: graphql.Float,
						Args: graphql.FieldConfigArgument{
							"symbol": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.PriceResolver,
					},
				},
			},
		),
	}

	return &root
}
