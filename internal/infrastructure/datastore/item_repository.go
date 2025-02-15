package datastore

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
)

type itemRepository struct {
	Conn *sqlx.DB
}

func (i itemRepository) Create(item *model.Item) error {
	query := `INSERT INTO items (title, price) VALUES ($1, $2)`
	_, err := i.Conn.Exec(query, item.Title, item.Price)
	return err
}

func (i itemRepository) GetItemByTitle(title string) (*model.Item, error) {
	var dbIt dbItem
	query := `SELECT id, title, price FROM items WHERE title = $1`
	err := i.Conn.Get(&dbIt, query, title)
	if err != nil {
		return nil, err
	}
	item := &model.Item{
		ID:    dbIt.ID,
		Title: dbIt.Title,
		Price: dbIt.Price,
	}
	return item, nil
}

func (i itemRepository) GetAllItems() ([]model.Item, error) {
	var dbItems []dbItem
	query := `SELECT id, title, price FROM items`
	err := i.Conn.Select(&dbItems, query)
	if err != nil {
		return nil, err
	}

	items := make([]model.Item, len(dbItems))
	for idx, dbIt := range dbItems {
		items[idx] = model.Item{
			ID:    dbIt.ID,
			Title: dbIt.Title,
			Price: dbIt.Price,
		}
	}
	return items, nil
}

func (i itemRepository) Update(item *model.Item) error {
	query := `UPDATE items SET title = $1, price = $2 WHERE id = $3`
	_, err := i.Conn.Exec(query, item.Title, item.Price, item.ID)
	return err
}

func (i itemRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := i.Conn.Exec(query, id)
	return err
}

func NewItemRepository(conn *sqlx.DB) repository.ItemRepository {
	return &itemRepository{conn}
}

type dbItem struct {
	ID    uuid.UUID `db:"id"`
	Title string    `db:"title"`
	Price int       `db:"price"`
}
