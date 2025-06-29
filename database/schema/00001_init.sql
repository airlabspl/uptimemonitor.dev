-- +goose Up
-- +goose StatementBegin
CREATE TABLE  IF NOT EXISTS users (
    id INTEGER NOT NULL  PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    email_verified_at DATETIME DEFAULT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
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

CREATE TABLE IF NOT EXISTS verifications (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT(CURRENT_TIMESTAMP),
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS password_resets (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token TEXT NOT NULL UNIQUE,
    created_at DATETIME NOT NULL DEFAULT(CURRENT_TIMESTAMP),
    expires_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE password_resets;
DROP TABLE email_verifications;
DROP TABLE sessions;
DROP TABLE users;
-- +goose StatementEnd
