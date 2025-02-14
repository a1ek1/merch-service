package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type ItemRepository interface {
	Create(item *model.Item) (uuid.UUID, error)
	GetItemById(id uuid.UUID) (*model.Item, error)
	GetAllItems() ([]model.Item, error)
	Update(item *model.Item) error
	Delete(id uuid.UUID) error
}
