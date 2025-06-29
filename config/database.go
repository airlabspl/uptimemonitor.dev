package config

import "flag"

var (
	DATABASE_DSN = flag.String("database-dsn", "db.sqlite?mode=wal", "Database connection string")
)
