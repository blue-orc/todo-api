package models

import (
	"time"
)

type Todo struct {
	TodoID      int       `json:"todoId"`
	Username    string    `json:"username"`
	CreatedOn   time.Time `json:"createdOn"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
