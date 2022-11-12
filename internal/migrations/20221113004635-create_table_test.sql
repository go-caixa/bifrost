
-- +migrate Up
CREATE TABLE test (
   role_id serial PRIMARY KEY,
   role_name VARCHAR (255) UNIQUE NOT NULL
);

-- +migrate Down
DROP TABLE test;
