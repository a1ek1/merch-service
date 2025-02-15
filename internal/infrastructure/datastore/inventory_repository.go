package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"
)

type inventoryRepository struct {
	Conn *sqlx.DB
}

func (i inventoryRepository) AddItem(userID uuid.UUID, itemID uuid.UUID, quantity int) error {
	query := `
		INSERT INTO inventory (user_id, item_id, quantity) 
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, item_id) 
		DO UPDATE SET quantity = inventory.quantity + $3, updated_at = NOW()
	`
	_, err := i.Conn.Exec(query, userID, itemID, quantity)
	return err
}

func (i inventoryRepository) GetByUserID(userID uuid.UUID) ([]model.Inventory, error) {
	var dbRecords []dbInventory
	query := `SELECT * FROM inventory WHERE user_id = $1`
	if err := i.Conn.Select(&dbRecords, query, userID); err != nil {
		return nil, err
	}

	inventory := make([]model.Inventory, len(dbRecords))
	for idx, rec := range dbRecords {
		inventory[idx] = model.Inventory{
			ID:       rec.ID,
			UserID:   rec.UserID,
			ItemID:   rec.ItemID,
			Quantity: rec.Quantity,
		}
	}
	return inventory, nil
}

func NewInventoryRepository(conn *sqlx.DB) repository.InventoryRepository {
	return &inventoryRepository{conn}
}

type dbInventory struct {
	ID        uuid.UUID `db:"id"`
	UserID    uuid.UUID `db:"user_id"`
	ItemID    uuid.UUID `db:"item_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
