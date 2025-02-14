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
	query := `INSERT INTO items (title, price) VALUES (:title, :price)`
	_, err := i.Conn.NamedExec(query, map[string]interface{}{
		"title": item.Title,
		"price": item.Price,
	})
	return err
}

func (i itemRepository) GetItemByTitle(title string) (*model.Item, error) {
	item := &model.Item{}
	query := `SELECT title FROM items WHERE title = :title`
	err := i.Conn.Get(item, query, title)
	return item, err
}

func (i itemRepository) GetAllItems() ([]model.Item, error) {
	items := []model.Item{}
	query := `SELECT title FROM items`
	err := i.Conn.Select(&items, query)
	return items, err
}

func (i itemRepository) Update(item *model.Item) error {
	query := `UPDATE items SET title = :title, price = :price WHERE id = :id`
	_, err := i.Conn.NamedExec(query, map[string]interface{}{
		"title": item.Title,
		"price": item.Price,
		"id":    item.ID,
	})
	return err
}

func (i itemRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = :id`
	_, err := i.Conn.Exec(query, id)
	return err
}

func NewItemRepository(conn *sqlx.DB) repository.ItemRepository {
	return &itemRepository{conn}
}
