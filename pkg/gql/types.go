package gql

// Company ...
// var Company = graphql.NewObject(
// 	graphql.ObjectConfig{
// 		Name: "Company",
// 		Fields: graphql.Fields{
// 			"symbol": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"name": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"exchange": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"industry": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"website": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"description": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"ceo": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"issueType": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"sector": &graphql.Field{
// 				Type: graphql.String,
// 			},
// 			"employees": &graphql.Field{
// 				Type: graphql.Int,
// 			},
// 			"tags": &graphql.Field{
// 				Type: graphql.NewList(graphql.String),
// 			},
// 			"test": &graphql.Field{
// 				Type: graphql.String,
// 				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 					return "hello world", nil
// 				},
// 			},
// 			"price": &graphql.Field{
// 				Type: graphql.Float,
// 				Resolve: ,
// 			}
// 		},
// 	},
// )
