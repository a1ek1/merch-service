package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetUserById(id uuid.UUID) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	UpdateBalance(id uuid.UUID, balance int) error
	Delete(id uuid.UUID) error
}
