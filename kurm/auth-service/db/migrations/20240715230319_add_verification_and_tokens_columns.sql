-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE users
ADD COLUMN refresh_token VARCHAR(255);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
DROP COLUMN refresh_token;
