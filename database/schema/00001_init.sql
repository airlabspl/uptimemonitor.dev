-- +goose Up
-- +goose StatementBegin
CREATE TABLE  IF NOT EXISTS users (
    id INTEGER NOT NULL  PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT,
    created_at DATETIME DEFAULT(CURRENT_TIMESTAMP)
);

CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER NOT NULL  PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT(CURRENT_TIMESTAMP),
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;
DROP TABLE users;
-- +goose StatementEnd
