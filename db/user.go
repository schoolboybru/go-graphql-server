package db

import "github.com/schoolboybru/go-graphql-server/model"

func (db *Db) CreateUser(user model.User) {

	defer db.Close()

	db.QueryRow("INSERT INTO users(id, email, password, user_name) VALUES ($1, $2, $3, $4)",
		&user.Id, &user.Email, &user.Password, &user.UserName)
}

func (db *Db) GetUserByName(userName string) model.User {
	var user model.User
	defer db.Close()

	db.QueryRow("SELECT * FROM USERS WHERE user_name=$1", userName).
		Scan(&user.Id, &user.Email, &user.Password, &user.UserName)

	return user
}
