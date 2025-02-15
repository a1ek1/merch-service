package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"
)

type authTokenRepository struct {
	Conn *sqlx.DB
}

func (a authTokenRepository) Create(token *model.AuthToken) error {
	dbToken := dbAuthToken{
		UserID:    token.UserID,
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
	}
	query := `INSERT INTO auth_tokens (user_id, token, expired_at) VALUES (:user_id, :token, :expired_at)`
	_, err := a.Conn.NamedExec(query, dbToken)
	return err
}

func (a authTokenRepository) GetByToken(token string) (*model.AuthToken, error) {
	var dbToken dbAuthToken
	query := `SELECT * FROM auth_tokens WHERE token = $1`
	if err := a.Conn.Get(&dbToken, query, token); err != nil {
		return nil, err
	}

	authToken := &model.AuthToken{
		ID:        dbToken.ID,
		UserID:    dbToken.UserID,
		Token:     dbToken.Token,
		CreatedAt: dbToken.CreatedAt,
		ExpiresAt: dbToken.ExpiresAt,
	}
	return authToken, nil
}

func (a authTokenRepository) Delete(token string) error {
	query := `DELETE FROM auth_tokens WHERE expired_at < NOW()`
	_, err := a.Conn.Exec(query)
	return err
}

func NewAuthTokenRepository(conn *sqlx.DB) repository.AuthTokenRepository {

	return &authTokenRepository{conn}
}

type dbAuthToken struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expired_at"`
}
