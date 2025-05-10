-- +goose Up
CREATE TABLE user_sets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    set_id UUID NOT NULL REFERENCES lego_sets(id) ON DELETE CASCADE,
    price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    owned_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE user_sets;