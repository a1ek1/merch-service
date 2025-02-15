package datastore

import (
	"database/sql"
	"errors"
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
	query := `INSERT INTO auth_tokens (user_id, token, expired_at) VALUES ($1, $2, $3)`
	_, err := a.Conn.Exec(query, token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (a authTokenRepository) GetActiveTokenByUserID(userID uuid.UUID) (*model.AuthToken, error) {
	var token model.AuthToken
	err := a.Conn.QueryRow(
		"SELECT token, expired_at FROM auth_tokens WHERE user_id = $1 AND expired_at > $2 ORDER BY expired_at DESC LIMIT 1",
		userID, time.Now(),
	).Scan(&token.Token, &token.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
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
