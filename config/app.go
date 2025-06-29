package config

import "flag"

var (
	addr       = flag.String("addr", ":4000", "server address")
	appUrl     = flag.String("app-url", "http://localhost:3000", "Base URL of the application")
	selfhosted = flag.Bool("selfhosted", false, "Run in self-hosted mode (no external services)")

	SetupFinished = false
)

func Addr() string {
	return *addr
}

func AppUrl() string {
	return *appUrl
}

func Selfhosted() bool {
	return *selfhosted
}
