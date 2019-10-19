package gql

import (
	iex "github.com/goinvest/iexcloud"
	"github.com/graphql-go/graphql"
	"github.com/mikerudolph/fin-graph/pkg/config"
)

// Root ...
type Root struct {
	Query *graphql.Object
}

// NewRoot ...
func NewRoot(conf config.Config) *Root {
	client := iex.NewClient(conf.IEXKey, conf.IEXBaseURL)

	resolver := Resolver{
		iexClient: client,
	}

	var Company = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Company",
			Fields: graphql.Fields{
				"symbol": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"exchange": &graphql.Field{
					Type: graphql.String,
				},
				"industry": &graphql.Field{
					Type: graphql.String,
				},
				"website": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
				"ceo": &graphql.Field{
					Type: graphql.String,
				},
				"issueType": &graphql.Field{
					Type: graphql.String,
				},
				"sector": &graphql.Field{
					Type: graphql.String,
				},
				"employees": &graphql.Field{
					Type: graphql.Int,
				},
				"tags": &graphql.Field{
					Type: graphql.NewList(graphql.String),
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
	)

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"companies": &graphql.Field{
						Type: graphql.NewList(Company),
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
