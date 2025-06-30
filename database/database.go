package database

import (
	"database/sql"
	"embed"
	"log/slog"
	"selfhosted/config"
	"selfhosted/database/store"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var (
	//go:embed schema/*.sql
	schemaFS embed.FS

	db *sql.DB
)

func DB() *sql.DB {
	return db
}

func Connect() {
	var err error

	dsn := config.DatabaseDsn
	slog.Info("Initializing database", "dsn", dsn)

	db, err = sql.Open("sqlite", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	goose.SetDialect("sqlite3")
	goose.SetBaseFS(schemaFS)

	err = goose.Up(db, "schema")
	if err != nil {
		panic(err)
	}
}

func New() *store.Queries {
	return store.New(db)
}
