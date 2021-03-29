package mutations

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
	"github.com/schoolboybru/go-graphql-server/db"
	"github.com/schoolboybru/go-graphql-server/graph/gqlTypes"
	"github.com/schoolboybru/go-graphql-server/model"
)

func CreatePost() *graphql.Field {
	var fields = graphql.Field{
		Type: gqlTypes.PostType,
		Description: "Create a post",
		Args: graphql.FieldConfigArgument{
			"content": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error){
			content := params.Args["content"].(string)

			newID := uuid.New()

			newPost := model.Post{
				Id: newID,
				Content: content,
			}

			database, err := db.New()

			if err != nil {
				panic(err.Error())
			}

			database.CreateNewPost("temp", newPost)

			return newPost, nil

		},
	}

	return &fields
}
