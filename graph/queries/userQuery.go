package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/schoolboybru/go-graphql-server/db"
	"github.com/schoolboybru/go-graphql-server/graph/gqlTypes"
)

func GetUserByName() *graphql.Field{
	var fields = graphql.Field{
		Type: gqlTypes.UserType,
		Description: "Get a user by ID",
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

				user := database.GetUserByName(id)

				return user, nil
			}
			return nil, nil
		},
	}
	return &fields
}
