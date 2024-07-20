-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE users
    ADD COLUMN github_username VARCHAR(50) UNIQUE;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
DROP COLUMN github_username;
