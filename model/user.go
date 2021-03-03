package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}
