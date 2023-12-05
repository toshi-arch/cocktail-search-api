
-- +migrate Up
CREATE TABLE IF NOT EXISTS cocktails (
    id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE,
    alcohol int NOT NULL,
    recipe text(255) NOT NULL,
    created_at datetime NOT NULL default CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL default CURRENT_TIMESTAMP,
    deleted_at datetime,

    PRIMARY KEY(id)
 );
-- +migrate Down
 DROP TABLE IF EXISTS cocktails;