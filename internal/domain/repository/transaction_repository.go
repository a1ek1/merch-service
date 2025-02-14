package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetByUserId(userId uuid.UUID) ([]model.Transaction, error)
}
