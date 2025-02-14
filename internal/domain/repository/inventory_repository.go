package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type InventoryRepository interface {
	AddItem(userID uuid.UUID, itemID uuid.UUID, quantity int) error
	GetByUserID(userID uuid.UUID) ([]model.Inventory, error)
}
