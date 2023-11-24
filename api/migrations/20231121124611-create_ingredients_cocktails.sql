
-- +migrate Up
CREATE TABLE IF NOT EXISTS ingredients_cocktails (
    id int NOT NULL AUTO_INCREMENT,
    ingredient_id int NOT NULL,
    cocktails_id int NOT NULL,
    amount int NOT NULL,
    unit int NOT NULL,
    created_at datetime NOT NULL default CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL default CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(ingredient_id) REFERENCES ingredients(id),
    FOREIGN KEY(cocktails_id) REFERENCES cocktails(id)
);
-- +migrate Down
DROP TABLE IF EXISTS ingredients_cocktails