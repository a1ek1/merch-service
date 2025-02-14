package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type inventoryRepository struct {
	Conn *sqlx.DB
}

func (i inventoryRepository) AddItem(userID uuid.UUID, itemID uuid.UUID, quantity int) error {
	query := `
		INSERT INTO inventory (user_id, item_id, quantity) 
		VALUES (:user_id, :item_id, :quantity)
		ON CONFLICT (user_id, item_id) 
		DO UPDATE SET quantity = inventory.quantity + :quantity, updated_at = NOW()
	`

	_, err := i.Conn.NamedExec(query, map[string]interface{}{
		"user_id":  userID,
		"item_id":  itemID,
		"quantity": quantity,
	})

	return err
}

func (i inventoryRepository) GetByUserID(userID uuid.UUID) ([]model.Inventory, error) {
	var inventory []model.Inventory
	query := `SELECT * FROM inventory WHERE user_id = $1`
	err := i.Conn.Select(&inventory, query, userID)
	return inventory, err
}

func NewInventoryRepository(Conn *sqlx.DB) repository.InventoryRepository {
	return &inventoryRepository{Conn}
}
