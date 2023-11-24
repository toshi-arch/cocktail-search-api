
-- +migrate Up
INSERT INTO cocktails(name, alcohol, recipe)
VALUES("ジントニック", 12, "省略1");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("ジンバック", 12, "省略2");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("ウォッカトニック", 12, "省略3");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("モスコミュール", 12, "省略4");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("ハイボール", 12, "省略5");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("XYZ", 25, "省略6");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("フリッパー", 9, "省略7");

INSERT INTO cocktails(name, alcohol, recipe)
VALUES("シャーリーテンプル", 0, "省略8");
-- +migrate Down
