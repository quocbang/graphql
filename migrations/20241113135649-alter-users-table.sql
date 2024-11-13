
-- +migrate Up
ALTER TABLE users ALTER COLUMN phone_number TYPE VARCHAR(10);
-- +migrate Down
