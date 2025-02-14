package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type PurchaseRepository interface {
	Create(purchase *model.Purchase) error
	GetByUserId(userId uuid.UUID) ([]model.Purchase, error)
}
