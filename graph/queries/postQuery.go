package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/schoolboybru/go-graphql-server/db"
	"github.com/schoolboybru/go-graphql-server/graph/gqlTypes"
)

func GetUserPosts() *graphql.Field {
	var fields = graphql.Field{
		Type: graphql.NewList(gqlTypes.PostType),
		Description: "Get all posts for a user",
		Args: graphql.FieldConfigArgument{
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error){
			id, ok := params.Args["userName"].(string)

			if ok {
				database, err := db.New()

				if err != nil {
					panic(err.Error())
				}

				posts := database.GetUserPosts(id)

				return posts, nil
			}

			return nil, nil
		},
	}

	return &fields
}
