package config

import "flag"

var (
	APP_URL = flag.String("app-url", "http://localhost:3000", "Base URL of the application")
)
