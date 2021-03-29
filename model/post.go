package model

import "github.com/google/uuid"

type Post struct {
	Id uuid.UUID `json:"id"`
	Content string `json:"content"`
}
