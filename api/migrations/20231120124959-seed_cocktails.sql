
-- +migrate Up
INSERT INTO cocktails(name, alcohol, recipe)VALUES
("ジントニック", 12, "省略1"),
("ジンバック", 12, "省略2"),
("ウォッカトニック", 12, "省略3"),
("モスコミュール", 12, "省略4"),
("ハイボール", 12, "省略5"),
("XYZ", 25, "省略6"),
("フリッパー", 9, "省略7"),
("シャーリーテンプル", 0, "省略8");
-- +migrate Down
