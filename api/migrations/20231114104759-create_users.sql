
-- +migrate Up
 CREATE TABLE IF NOT EXISTS users (id int,name VARCHAR(100),deleted_at DATETIME);

-- +migrate Down
DROP TABLE IF EXISTS users;