-- +goose Up
CREATE TABLE cats (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);

-- +goose Down
DROP TABLE cats;
