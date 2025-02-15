package repository

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
)

type AuthTokenRepository interface {
	Create(token *model.AuthToken) error
	GetActiveTokenByUserID(userID uuid.UUID) (*model.AuthToken, error)
	GetByToken(token string) (*model.AuthToken, error)
	Delete(token string) error
}
