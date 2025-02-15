package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"
)

type transactionRepository struct {
	Conn *sqlx.DB
}

func (t transactionRepository) Create(transaction *model.Transaction) error {
	dbTrans := dbTransaction{
		UserID:    transaction.UserID,
		ToUserID:  transaction.ToUserID,
		Amount:    transaction.Amount,
		CreatedAt: time.Now(),
	}

	query := `INSERT INTO transactions (user_id, to_user_id, amount, created_at) 
			  VALUES (:user_id, :to_user_id, :amount, :created_at)`

	_, err := t.Conn.NamedExec(query, dbTrans)
	return err
}

func (t transactionRepository) GetByUserId(userId uuid.UUID) ([]model.Transaction, error) {
	var dbTransactions []dbTransaction
	query := `SELECT id, user_id, to_user_id, amount, created_at 
	          FROM transactions 
	          WHERE user_id = $1 OR to_user_id = $1`
	err := t.Conn.Select(&dbTransactions, query, userId)
	if err != nil {
		return nil, err
	}

	transactions := make([]model.Transaction, len(dbTransactions))
	for i, dbTrans := range dbTransactions {
		transactions[i] = model.Transaction{
			ID:        dbTrans.ID,
			UserID:    dbTrans.UserID,
			ToUserID:  dbTrans.ToUserID,
			Amount:    dbTrans.Amount,
			CreatedAt: dbTrans.CreatedAt,
		}
	}

	return transactions, nil
}

func NewTransactionRepository(Conn *sqlx.DB) repository.TransactionRepository {
	return &transactionRepository{Conn}
}

type dbTransaction struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ToUserID  uuid.UUID `db:"to_user_id"`
	Amount    int       `db:"amount"`
	CreatedAt time.Time `db:"created_at"`
}
