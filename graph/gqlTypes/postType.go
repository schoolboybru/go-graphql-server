package gqlTypes

import "github.com/graphql-go/graphql"

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
	},
})
