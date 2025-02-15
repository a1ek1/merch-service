package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"
)

type purchaseRepository struct {
	Conn *sqlx.DB
}

func (p purchaseRepository) Create(purchase *model.Purchase) error {
	dbPur := dbPurchase{
		UserID:      purchase.UserID,
		ItemID:      purchase.ItemID,
		Quantity:    purchase.Quantity,
		TotalPrice:  purchase.TotalPrice,
		PurchasedAt: time.Now(),
	}

	query := `INSERT INTO purchases (user_id, item_id, quantity, total_price, purchased_at) 
			  VALUES (:user_id, :item_id, :quantity, :total_price, :purchased_at)`

	_, err := p.Conn.NamedExec(query, dbPur)
	return err
}

func (p purchaseRepository) GetByUserId(userId uuid.UUID) ([]model.Purchase, error) {
	var dbPurchases []dbPurchase
	query := `SELECT id, user_id, item_id, quantity, total_price, purchased_at FROM purchases WHERE user_id = $1`
	err := p.Conn.Select(&dbPurchases, query, userId)
	if err != nil {
		return nil, err
	}

	purchases := make([]model.Purchase, len(dbPurchases))
	for i, dbPur := range dbPurchases {
		purchases[i] = model.Purchase{
			ID:          dbPur.ID,
			UserID:      dbPur.UserID,
			ItemID:      dbPur.ItemID,
			Quantity:    dbPur.Quantity,
			TotalPrice:  dbPur.TotalPrice,
			PurchasedAt: dbPur.PurchasedAt,
		}
	}

	return purchases, nil
}

func NewPurchaseRepository(Conn *sqlx.DB) repository.PurchaseRepository {
	return &purchaseRepository{Conn}
}

type dbPurchase struct {
	ID          uuid.UUID `db:"id"`
	UserID      uuid.UUID `db:"user_id"`
	ItemID      uuid.UUID `db:"item_id"`
	Quantity    int       `db:"quantity"`
	TotalPrice  int       `db:"total_price"`
	PurchasedAt time.Time `db:"purchased_at"`
}
