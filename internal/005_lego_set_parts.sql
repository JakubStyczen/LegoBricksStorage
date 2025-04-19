-- +goose Up
CREATE TABLE lego_set_parts (
    set_id INT REFERENCES lego_sets(id) ON DELETE CASCADE,
    part_id INT REFERENCES lego_parts(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    PRIMARY KEY (set_id, part_id)
);

-- +goose Down
DROP TABLE lego_set_parts;