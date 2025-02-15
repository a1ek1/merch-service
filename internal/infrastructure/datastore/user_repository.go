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
              VALUES (:id, :username, :password, :coins, :created_at, :updated_at)`

	dbUser := dbUser{
		ID:        uuid.New(),
		Username:  user.Username,
		Password:  user.Password,
		Coins:     1000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := u.Conn.NamedExec(query, dbUser)
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
	query := `UPDATE users SET coins = :coins, updated_at = NOW() WHERE id = :id`
	_, err := u.Conn.NamedExec(query, map[string]interface{}{
		"id":    id,
		"coins": balance,
	})
	return err
}

func (u userRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = :id`
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
