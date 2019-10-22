package gql

import (
	iex "github.com/goinvest/iexcloud"
	"github.com/graphql-go/graphql"
	"github.com/mikerudolph/fin-graph/pkg/config"
)

// Resolver ...
type Resolver struct {
	iexClient *iex.Client
}

// NewResolver ...
func NewResolver(conf config.Config) *Resolver {
	client := iex.NewClient(conf.IEXKey, conf.IEXBaseURL)

	// define the graphql resolver service where fetches to iex will live
	resolver := Resolver{
		iexClient: client,
	}

	return &resolver
}

// CompanyResolver ...
func (r *Resolver) CompanyResolver(p graphql.ResolveParams) (interface{}, error) {

	symbol, ok := p.Args["symbol"].(string)
	if ok {
		company, _ := r.iexClient.Company(symbol)
		return company, nil
	}

	return nil, nil
}

// PriceResolver ...
func (r *Resolver) PriceResolver(p graphql.ResolveParams) (interface{}, error) {
	var symbol string

	symbol, _ = p.Args["symbol"].(string)
	if symbol == "" {
		source := p.Source.(iex.Company)
		symbol = source.Symbol
	}

	price, _ := r.iexClient.Price(symbol)

	return price, nil
}
