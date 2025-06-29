package mailer

import "selfhosted/config"

func VerificationMessage(email, token string) Message {
	return Message{
		To:      email,
		Subject: "Email Verification",
		Body:    "Please verify your email by clicking the following link: " + config.AppUrl() + "/auth/verify/" + token,
	}
}

func PasswordResetMessage(email, token string) Message {
	return Message{
		To:      email,
		Subject: "Password Reset",
		Body:    "Please reset your password by clicking the following link: " + config.AppUrl() + "/reset-password/" + token,
	}
}
