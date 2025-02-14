package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type transactionRepository struct {
	Conn *sqlx.DB
}

func (t transactionRepository) Create(transaction *model.Transaction) error {
	query := `INSERT INTO transactions (user_id, to_user_id, amount) VALUES (:user_id, :to_user_id, :amount)`
	_, err := t.Conn.NamedExec(query, map[string]interface{}{
		"user_id":    transaction.UserID,
		"to_user_id": transaction.ToUserID,
		"amount":     transaction.Amount,
	})
	return err
}

func (t transactionRepository) GetByUserId(userId uuid.UUID) ([]model.Transaction, error) {
	transactions := []model.Transaction{}
	query := `SELECT * FROM transactions WHERE user_id = $1 OR to_user_id = $1`
	err := t.Conn.Select(&transactions, query, userId)
	return transactions, err
}

func NewTransactionRepository(Conn *sqlx.DB) repository.TransactionRepository {
	return &transactionRepository{Conn}
}
