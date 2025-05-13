-- +goose Up
CREATE TABLE lego_set_parts (
    set_serial VARCHAR NOT NULL REFERENCES lego_sets(serial_number) ON DELETE CASCADE,
    part_serial VARCHAR NOT NULL REFERENCES lego_parts(serial_number) ON DELETE CASCADE,
    quantity INT NOT NULL,
    PRIMARY KEY (set_serial, part_serial)
);

-- +goose Down
DROP TABLE lego_set_parts;