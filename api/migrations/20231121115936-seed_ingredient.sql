
-- +migrate Up
INSERT INTO ingredients (name, type, alcohol)VALUES
("ドライジン", 1, 40),
("ウォッカ", 1, 40),
("ウィスキー", 1, 40),
("ラム",1, 40),
("トニックウォーター", 6, 0),
("ジンジャーエール", 6, 0),
("ソーダ", 6, 0),
("コアントロー", 2, 40),
("レモンジュース", 6, 0),
("アマレット", 2, 24),
("ドライベルモット", 4, 17),
("カンパリ", 2, 25),
("グレナデンシロップ", 6, 0),
("ライムジュース", 6, 0),
("カシス", 2, 18);

-- +migrate Down
