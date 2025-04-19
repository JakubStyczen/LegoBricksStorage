-- +goose Up
CREATE TABLE user_sets (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    set_id INT REFERENCES lego_sets(id) ON DELETE CASCADE,
    price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    owned_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE user_sets;