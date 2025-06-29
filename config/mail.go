package config

import "flag"

var (
	SMTP_HOST              = flag.String("smtp-host", "localhost", "SMTP server host")
	SMTP_PORT              = flag.String("smtp-port", "587", "SMTP server port")
	SMTP_USER              = flag.String("smtp-user", "", "SMTP server username")
	SMTP_PASS              = flag.String("smtp-pass", "", "SMTP server password")
	MAIL_FROM              = flag.String("mail-from", "Uptime Monitor <no-reply@example.com>", "Email address to send from")
	AUTOMATIC_VERIFICATION = flag.Bool("automatic-verification", false, "Automatically verify user emails without sending a verification email")
)
