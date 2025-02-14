package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (uuid.UUID, error)
	GetUserById(id uuid.UUID) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	Update(user *model.User) error
	Delete(id uuid.UUID) error
}
