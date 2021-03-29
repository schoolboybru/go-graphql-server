package mutations

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/schoolboybru/go-graphql-server/db"
	"github.com/schoolboybru/go-graphql-server/graph/gqlTypes"
	"github.com/schoolboybru/go-graphql-server/model"
	"github.com/schoolboybru/go-graphql-server/util"
)

func CreateUser() *graphql.Field {

	var fields = graphql.Field{
		Type: gqlTypes.UserType,
		Description: "Create a user",
		Args: graphql.FieldConfigArgument{
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"userName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			email, _ := params.Args["email"].(string)
			password, _ := params.Args["password"].(string)
			userName, _ := params.Args["userName"].(string)

			newID := uuid.New()

			hashedPass, err := util.HashPassword(password)

			if err != nil {
				return nil, err
			}

			newUser := model.User{
				Id:       newID,
				Email:    email,
				Password: hashedPass,
				UserName: userName,
			}

			database, err := db.New()

			if err != nil {
				panic(err.Error())
			}

			database.CreateUser(newUser)

			return newUser, nil
		},
	}
	return &fields
}
