package model

import (
	"github.com/google/uuid"
	"time"
)

type Purchase struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ItemID      uuid.UUID `json:"item_id"`
	Quantity    int       `json:"quantity"`
	TotalPrice  int       `json:"total_price"`
	PurchasedAt time.Time `json:"created_at"`
}
