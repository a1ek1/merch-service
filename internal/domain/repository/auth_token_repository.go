package repository

import (
	"merch-service/internal/domain/model"
)

type AuthTokenRepository interface {
	Create(token *model.AuthToken) error
	GetByToken(token string) (*model.AuthToken, error)
	Delete(token string) error
}
