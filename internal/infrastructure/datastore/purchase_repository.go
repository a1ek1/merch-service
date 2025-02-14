package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type purchaseRepository struct {
	Conn *sqlx.DB
}

func (p purchaseRepository) Create(purchase *model.Purchase) error {
	query := `INSERT INTO purchases (user_id, item_id, quantity, total_price) VALUES (:user_id, :item_id, :quantity, :total_price)`
	_, err := p.Conn.NamedExec(query, map[string]interface{}{
		"user_id":     purchase.UserID,
		"item_id":     purchase.ItemID,
		"quantity":    purchase.Quantity,
		"total_price": purchase.TotalPrice,
	})
	return err
}

func (p purchaseRepository) GetByUserId(userId uuid.UUID) ([]model.Purchase, error) {
	purchases := []model.Purchase{}
	query := `SELECT * FROM purchases WHERE user_id = $1`
	err := p.Conn.Select(&purchases, query, userId)
	return purchases, err
}

func NewPurchaseRepository(Conn *sqlx.DB) repository.PurchaseRepository {
	return &purchaseRepository{Conn}
}
