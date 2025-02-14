package model

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ToUserID  uuid.UUID `json:"to_user_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
