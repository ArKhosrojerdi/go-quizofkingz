-- +migrate Up
ALTER TABLE users ADD COLUMN password TEXT NOT NULL;

-- +migrate Down
ALTER TABLE users DROP COLUMN password;