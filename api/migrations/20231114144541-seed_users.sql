
-- +migrate Up
INSERT INTO users
    (name) VALUES ("taro"),("jiro"),("saburo");

-- +migrate Down
