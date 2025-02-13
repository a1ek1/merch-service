-- +goose Up
-- +goose StatementBegin
CREATE TABLE purchases (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    item_id UUID REFERENCES items(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK ( quantity > 0 ),
    total_price INT NOT NULL CHECK ( total_price > 0 ),
    purchased_at TIMESTAMPTZ DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS purchases;
-- +goose StatementEnd
