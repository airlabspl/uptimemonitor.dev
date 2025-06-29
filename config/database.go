package config

import "flag"

var (
	databaseDsn = flag.String("database-dsn", "db.sqlite?mode=wal", "Database connection string")
)

func DatabaseDsn() string {
	return *databaseDsn
}
