package datastore

import (
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type authTokenRepository struct {
	Conn *sqlx.DB
}

func (a authTokenRepository) Create(token *model.AuthToken) error {
	query := `INSERT INTO auth_tokens (user_id, token, expired_at) VALUES (:user_id, :token, :expired_at)`
	_, err := a.Conn.NamedExec(query, map[string]interface{}{
		"user_id":    token.UserID,
		"token":      token.Token,
		"expired_at": token.ExpiresAt,
	})
	return err
}

func (a authTokenRepository) GetByToken(token string) (*model.AuthToken, error) {
	authToken := &model.AuthToken{}
	query := `SELECT * FROM auth_tokens WHERE token = $1`
	err := a.Conn.Get(authToken, query, token)
	return authToken, err

}

func (a authTokenRepository) Delete(token string) error {
	query := `DELETE FROM auth_tokens WHERE expired_at < NOW()`
	_, err := a.Conn.Exec(query)
	return err
}

func NewAuthTokenRepository(conn *sqlx.DB) repository.AuthTokenRepository {
	return &authTokenRepository{conn}
}
