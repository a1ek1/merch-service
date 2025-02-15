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
	dbIt := dbItem{
		Title: item.Title,
		Price: item.Price,
	}
	query := `INSERT INTO items (title, price) VALUES (:title, :price)`
	_, err := i.Conn.NamedExec(query, dbIt)
	return err
}

func (i itemRepository) GetItemByTitle(title string) (*model.Item, error) {
	var dbIt dbItem
	query := `SELECT id, title, price FROM items WHERE title = :title`
	err := i.Conn.Get(&dbIt, query, map[string]interface{}{
		"title": title,
	})
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

	// Преобразуем полученные данные в доменную модель
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
	query := `UPDATE items SET title = :title, price = :price WHERE id = :id`
	_, err := i.Conn.NamedExec(query, map[string]interface{}{
		"id":    item.ID,
		"title": item.Title,
		"price": item.Price,
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

type dbItem struct {
	ID    uuid.UUID `db:"id"`
	Title string    `db:"title"`
	Price int       `db:"price"`
}
