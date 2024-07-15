-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE users
    ADD COLUMN verified BOOLEAN DEFAULT false,
ADD COLUMN code VARCHAR(6);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
DROP COLUMN verified,
DROP COLUMN code;
