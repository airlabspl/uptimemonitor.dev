package config

import "flag"

var (
	smtpHost = flag.String("smtp-host", "localhost", "SMTP server host")
	smtpPort = flag.String("smtp-port", "587", "SMTP server port")
	smtpUser = flag.String("smtp-user", "", "SMTP server username")
	smtpPass = flag.String("smtp-pass", "", "SMTP server password")
	mailFrom = flag.String("mail-from", "Uptime Monitor <no-reply@example.com>", "Email address to send from")
)

func SmtpHost() string {
	return *smtpHost
}

func SmtpPort() string {
	return *smtpPort
}

func SmtpUser() string {
	return *smtpUser
}

func SmtpPass() string {
	return *smtpPass
}

func MailFrom() string {
	return *mailFrom
}
