CREATE TABLE rsv_null_in_json(
    i INT PRIMARY KEY,
    j JSON NOT NULL,
    jb jsonb NOT NULL
);

INSERT INTO rsv_null_in_json(i, j, jb) VALUES
(1, 'null', 'null'),
(2, '"null"', '"null"');