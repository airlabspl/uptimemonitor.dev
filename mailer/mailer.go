package mailer

import (
	"log/slog"
	"net/smtp"
	"selfhosted/config"
)

type Message struct {
	To      string
	Subject string
	Body    string
}

func Send(message Message) error {
	auth := smtp.CRAMMD5Auth(config.SmtpUser, config.SmtpPass)
	addr := config.SmtpHost + ":" + config.SmtpPort
	from := config.MailFrom

	body := "Subject: " + message.Subject + "\r\n" +
		"From: " + from + "\r\n" +
		"To: " + message.To + "\r\n" +
		"\r\n" + // Blank line between headers and body
		message.Body + "\r\n"

	err := smtp.SendMail(addr, auth, from, []string{message.To}, []byte(body))
	if err != nil {
		slog.Error("failed to send email", "to", message.To, "subject", message.Subject, "from", from, "error", err, "context", "mailer.Send")
		return err
	}

	slog.Info("email sent", "to", message.To, "subject", message.Subject, "from", from, "context", "mailer.Send")

	return nil
}
