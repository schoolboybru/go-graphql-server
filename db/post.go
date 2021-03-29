package db

import (
	"github.com/schoolboybru/go-graphql-server/model"
)

func (db *Db) GetUserPosts(userName string) []model.Post {
	var posts []model.Post
	defer db.Close()

	rows, err := db.Query(
		`SELECT p.id, content FROM POST p
    			INNER JOIN USER_HAS_POST up ON up.post_id = p.id
    			INNER JOIN USERS u ON up.user_id = u.id WHERE u.user_name = $1`, userName)

	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var post model.Post

		err := rows.Scan(&post.Id, &post.Content)

		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, post)
	}

	return posts
}

func (db *Db) CreateNewPost(userId string, post model.Post) {
	defer db.Close()

	err := db.QueryRow(
		`WITH post_key AS (INSERT INTO POST (id, content) VALUES ($1, $2) RETURNING id)
				INSERT INTO USER_HAS_POST (user_id, post_id)
				SELECT $1, post_key.id FROM post_key;`).
		Scan(&post.Id, &post.Content, &userId)

	if err != nil {
		panic(err.Error())
	}
}
