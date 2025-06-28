package database

import (
	"database/sql"
	"embed"
	"selfhosted/database/store"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var (
	//go:embed schema/*.sql
	schemaFS embed.FS

	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("sqlite", ":memory:")
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
