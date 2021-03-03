package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/schoolboybru/go-graphql-server/graph/mutations"
	"github.com/schoolboybru/go-graphql-server/graph/queries"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": mutations.CreateUser(),
}})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getUserByName": queries.GetUserByName(),
	},
})

var UserSchema,_ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	Mutation: rootMutation,
})


