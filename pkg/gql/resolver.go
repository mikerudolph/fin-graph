package gql

import (
	iex "github.com/goinvest/iexcloud"
	"github.com/graphql-go/graphql"
)

// Resolver ...
type Resolver struct {
	iexClient *iex.Client
}

// CompanyResolver ...
func (r *Resolver) CompanyResolver(p graphql.ResolveParams) (interface{}, error) {
	var companies []iex.Company

	symbol, ok := p.Args["symbol"].(string)
	if ok {
		company, _ := r.iexClient.Company(symbol)

		companies = append(companies, company)

		return companies, nil
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
