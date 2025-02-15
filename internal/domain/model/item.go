package model

import (
	"github.com/google/uuid"
)

type Item struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"name"`
	Price int       `json:"price"`
}
