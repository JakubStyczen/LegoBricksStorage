-- +goose Up
CREATE TABLE lego_sets (
    id SERIAL PRIMARY KEY,
    serial_number TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    price DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    theme TEXT NOT NULL DEFAULT 'Unknown',
    year INT NOT NULL DEFAULT EXTRACT(YEAR FROM CURRENT_DATE),
    total_parts INT NOT NULL

);

-- +goose Down
DROP TABLE lego_sets;