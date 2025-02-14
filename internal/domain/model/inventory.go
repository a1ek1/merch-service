package model

import "github.com/google/uuid"

type Inventory struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	ItemID   uuid.UUID `json:"item_id"`
	Quantity int       `json:"quantity"`
}
