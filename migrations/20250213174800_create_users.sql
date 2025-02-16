-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    coins INT DEFAULT 1000,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO users (id, username, password) VALUES
                                     ('000d41d4-dd00-4ed4-8d14-5483e3147b31', 'oleg', 'passwd123'),
                                     ('e150ea00-ad91-4459-852e-06f05d60c77b', 'vlad', 'passwd123');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
