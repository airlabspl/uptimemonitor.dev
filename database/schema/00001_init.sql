-- +goose Up
-- +goose StatementBegin
CREATE TABLE  IF NOT EXISTS users (
    id INTEGER NOT NULL AUTOINCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password_hash TEXT,
    created_at DATETIME DEFAULT(CURRENT_TIMESTAMP)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
