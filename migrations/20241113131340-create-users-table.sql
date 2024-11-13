
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
  id UUID PRIMARY KEY,
  username text NOT NULL,
  password text NOT NULL,
  email text NOT NULL,
  phone_number VARCHAR,
  full_name text NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ  
);
-- +migrate Down
