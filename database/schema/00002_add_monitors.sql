-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS monitors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    url TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd