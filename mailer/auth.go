package mailer

import "selfhosted/config"

func VerificationMessage(email, token string) Message {
	return Message{
		To:      email,
		Subject: "Email Verification",
		Body:    "Please verify your email by clicking the following link: " + *config.APP_URL + "/auth/verify/" + token,
	}
}
