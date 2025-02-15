package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type ItemRepository interface {
	Create(item *model.Item) error
	GetItemByTitle(title string) (*model.Item, error)
	GetAllItems() ([]model.Item, error)
	Update(item *model.Item) error
	Delete(id uuid.UUID) error
}
