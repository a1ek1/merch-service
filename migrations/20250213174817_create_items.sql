-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT UNIQUE NOT NULL,
    price INT NOT NULL CHECK ( price > 0 )
);

-- Добавление начальных записей
INSERT INTO items (title, price) VALUES
                                     ('t-shirt', 80),
                                     ('cup', 20),
                                     ('book', 50),
                                     ('pen', 10),
                                     ('powerbank', 200),
                                     ('hoody', 300),
                                     ('umbrella', 200),
                                     ('socks', 10),
                                     ('wallet', 50),
                                     ('pink-hoody', 500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
