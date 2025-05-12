-- +goose Up
CREATE TABLE lego_parts (
    id UUID PRIMARY KEY,
    serial_number TEXT UNIQUE NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    name TEXT NOT NULL,
    color TEXT NOT NULL DEFAULT 'Unknown',
    price DECIMAL(10,2) NOT NULL DEFAULT 0.00
);

-- +goose Down
DROP TABLE lego_parts;
