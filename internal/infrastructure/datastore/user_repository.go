package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"
)

type userRepository struct {
	Conn *sqlx.DB
}

func (u userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (id, username, password, coins, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := u.Conn.Exec(query, user.ID, user.Username, user.Password, 1000, time.Now(), time.Now())
	if err != nil {
		log.Printf("Ошибка при вставке пользователя: %v", err)
	}
	return err
}

func (u userRepository) GetUserById(id uuid.UUID) (*model.User, error) {
	var dbUser dbUser
	query := `SELECT id, username, password, coins, created_at, updated_at FROM users WHERE id = $1`
	err := u.Conn.Get(&dbUser, query, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		Coins:     dbUser.Coins,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func (u userRepository) GetUserByUsername(username string) (*model.User, error) {
	var dbUser dbUser
	query := `SELECT id, username, password, coins, created_at, updated_at FROM users WHERE username = $1`
	err := u.Conn.Get(&dbUser, query, username)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		Coins:     dbUser.Coins,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
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

type dbUser struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Coins     int       `db:"coins"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
