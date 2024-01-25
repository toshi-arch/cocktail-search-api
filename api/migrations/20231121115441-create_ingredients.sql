
-- +migrate Up
CREATE TABLE IF NOT EXISTS ingredients (
    id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255)  NOT NULL UNIQUE,
    type int NOT NULL,
    alcohol int,
    created_at datetime NOT NULL default CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL default CURRENT_TIMESTAMP,
    deleted_at datetime,

    PRIMARY KEY(id)
 );
-- +migrate Down
DROP TABLE IF EXISTS ingredients;