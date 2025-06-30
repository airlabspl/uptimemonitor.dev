package main

import (
	"context"
	"flag"
	"net/http"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/handler"
)

func init() {
	flag.StringVar(&config.Addr, "addr", ":4000", "server address")
	flag.StringVar(&config.AppUrl, "app-url", "http://localhost:3000", "Base URL of the application")
	flag.BoolVar(&config.Selfhosted, "selfhosted", false, "Run in self-hosted mode (no external services)")
	flag.StringVar(&config.DatabaseDsn, "database-dsn", "db.sqlite?mode=wal", "Database connection string")
	flag.StringVar(&config.SmtpHost, "smtp-host", "localhost", "SMTP server host")
	flag.StringVar(&config.SmtpPort, "smtp-port", "587", "SMTP server port")
	flag.StringVar(&config.SmtpUser, "smtp-user", "", "SMTP server username")
	flag.StringVar(&config.SmtpPass, "smtp-pass", "", "SMTP server password")
	flag.StringVar(&config.MailFrom, "mail-from", "Uptime Monitor <no-reply@example.com>", "Email address to send from")

	flag.Parse()
}

func main() {
	adminCount, err := database.New().CountAdminUsers(context.Background())
	if err != nil {
		panic(err)
	}

	config.SetupFinished = adminCount > 0

	s := &http.Server{
		Addr:    config.Addr,
		Handler: handler.NewRouter(),
	}

	s.ListenAndServe()
}
