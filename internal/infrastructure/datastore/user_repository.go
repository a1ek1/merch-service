package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type userRepository struct {
	Conn *sqlx.DB
}

func (u userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (username, password) VALUES (:username, :password)`
	_, err := u.Conn.Exec(query, map[string]interface{}{
		"username": user.Username,
		"password": user.Password,
	})
	return err
}

func (u userRepository) GetUserById(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	query := `SELECT * FROM users WHERE id = $1`
	err := u.Conn.Get(user, query, id)
	return user, err
}

func (u userRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT * FROM users WHERE username = $1`
	err := u.Conn.Get(user, query, username)
	return user, err
}

func (u userRepository) UpdateBalance(id uuid.UUID, balance int) error {
	query := `UPDATE users SET coins = $1, updated_at = NOW() WHERE id = $2`
	_, err := u.Conn.Exec(query, balance, id)
	return err
}

func (u userRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := u.Conn.Exec(query, id)
	return err
}

func NewUserRepository(Conn *sqlx.DB) repository.UserRepository {
	return &userRepository{Conn}
}
