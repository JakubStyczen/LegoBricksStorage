-- +goose Up
CREATE TABLE lego_parts (
    id SERIAL PRIMARY KEY,
    serial_number TEXT UNIQUE NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    name TEXT NOT NULL,
    color TEXT,
    price DECIMAL(10,2)
);

-- +goose Down
DROP TABLE lego_parts;
