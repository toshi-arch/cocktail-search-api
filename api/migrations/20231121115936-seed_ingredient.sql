
-- +migrate Up
INSERT INTO ingredients (name, type, alcohol)
VALUES("ドライジン", 1, 40);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ウォッカ", 1, 40);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ウィスキー", 1, 40);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ラム",1, 40);

INSERT INTO ingredients (name, type, alcohol)
VALUES("トニックウォーター", 6, 0);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ジンジャーエール", 6, 0);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ソーダ", 6, 0);

INSERT INTO ingredients (name, type, alcohol)
VALUES("コアントロー", 2, 40);

INSERT INTO ingredients (name, type, alcohol)
VALUES("レモンジュース", 6, 0);

INSERT INTO ingredients (name, type, alcohol)
VALUES("アマレット", 2, 24);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ドライベルモット", 4, 17);

INSERT INTO ingredients (name, type, alcohol)
VALUES("カンパリ", 2, 25);

INSERT INTO ingredients (name, type, alcohol)
VALUES("グレナデンシロップ", 6, 0);

INSERT INTO ingredients (name, type, alcohol)
VALUES("ライムジュース", 6, 0);
-- +migrate Down
