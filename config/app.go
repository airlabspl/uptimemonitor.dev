package config

import "flag"

var (
	appUrl     = flag.String("app-url", "http://localhost:3000", "Base URL of the application")
	selfhosted = flag.Bool("selfhosted", false, "Run in self-hosted mode (no external services)")

	SetupFinished = false
)

func AppUrl() string {
	return *appUrl
}

func Selfhosted() bool {
	return *selfhosted
}
