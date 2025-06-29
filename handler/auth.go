package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"
	"selfhosted/mailer"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginForm(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("decode error", "context", "LoginForm", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(req); err != nil {
		slog.Error("validation error", "context", "LoginForm", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := database.New().GetUserByEmail(r.Context(), req.Email)
	if err != nil || user.ID == 0 {
		slog.Error("user not found", "context", "LoginForm", "email", req.Email, "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		slog.Error("password mismatch", "context", "LoginForm", "userID", user.ID)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := database.New().CreateSession(r.Context(), store.CreateSessionParams{
		Uuid:      uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour * 30),
	})
	if err != nil {
		slog.Error("session creation error", "context", "LoginForm", "userID", user.ID, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.Uuid,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // todo
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}

func RegisterForm(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Name            string `json:"name" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("decode error", "context", "RegisterForm", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(req); err != nil {
		slog.Error("validation error", "context", "RegisterForm", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("password hash error", "context", "RegisterForm", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	existing, err := database.New().GetUserByEmail(r.Context(), req.Email)
	if err == nil && existing.ID != 0 {
		slog.Error("user already exists", "context", "RegisterForm", "email", req.Email)
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	user, err := database.New().CreateUser(r.Context(), store.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		slog.Error("user creation error", "context", "RegisterForm", "email", req.Email, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	session, err := database.New().CreateSession(r.Context(), store.CreateSessionParams{
		Uuid:      uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour * 30),
	})
	if err != nil {
		slog.Error("session creation error", "context", "RegisterForm", "userID", user.ID, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	go func(user store.User) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		token := uuid.NewString()
		err := database.New().CreateVerification(ctx, store.CreateVerificationParams{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: time.Now().Add(24 * time.Hour),
		})
		if err != nil {
			slog.Error("verification creation error", "context", "RegisterForm", "userID", user.ID, "error", err)
			return
		}
		slog.Info("verification created", "context", "RegisterForm", "userID", user.ID)
		err = mailer.Send(mailer.VerificationMessage(user.Email, token))
		if err != nil {
			slog.Error("verification email send error", "context", "RegisterForm", "userID", user.ID, "error", err)
			return
		}
		slog.Info("verification email sent", "context", "RegisterForm", "userID", user.ID)
	}(user)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    session.Uuid,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // todo
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}

func Verification(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	if token == "" {
		slog.Error("verification token is empty", "context", "Verification")
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	verification, err := database.New().GetVerificationByToken(r.Context(), token)
	if err != nil || verification.ID == 0 {
		slog.Error("verification not found", "context", "Verification", "token", token, "error", err)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if verification.ExpiresAt.Before(time.Now()) {
		slog.Error("verification token expired", "context", "Verification", "token", token)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	err = database.New().VerifyUser(r.Context(), store.VerifyUserParams{
		ID:              verification.UserID,
		EmailVerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		slog.Error("user verification error", "context", "Verification", "userID", verification.UserID, "error", err)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	slog.Info("user verified", "context", "Verification", "userID", verification.UserID)

	err = database.New().DeleteVerification(r.Context(), verification.ID)
	if err != nil {
		slog.Error("verification deletion error", "context", "Verification", "verificationID", verification.ID, "error", err)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	slog.Info("verification deleted", "context", "Verification", "verificationID", verification.ID)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil || c.Value == "" {
		slog.Error("session cookie not found", "context", "Logout", "error", err)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = database.New().DeleteSession(r.Context(), c.Value)
	if err != nil {
		slog.Error("session deletion error", "context", "Logout", "cookie", c.Value, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // todo
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}

func ResetPasswordLink(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Email string `json:"email" validate:"required,email"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("decode error", "context", "ResetPasswordLink", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := database.New().GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		slog.Error("user not found", "context", "ResetPasswordLink", "error", err, "email", req.Email)
		w.WriteHeader(http.StatusOK)
		return
	}

	token := uuid.NewString()
	err = database.New().CreatePasswordReset(r.Context(), store.CreatePasswordResetParams{
		UserID:    user.ID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})

	go func(user store.User, token string) {
		err = mailer.Send(mailer.PasswordResetMessage(user.Email, token))
		if err != nil {
			slog.Error("password reset email send error", "context", "ResetPasswordLink", "userID", user.ID, "error", err)
			return
		}
		slog.Info("password reset email sent", "context", "ResetPasswordLink", "userID", user.ID)
	}(user, token)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Token           string `json:"token" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("decode error", "context", "ResetPassword", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return
	}

	passwordReset, err := database.New().GetPasswordResetByToken(r.Context(), req.Token)
	if err != nil {
		slog.Error("password reset not found", "context", "ResetPassword", "error", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if passwordReset.ExpiresAt.Before(time.Now()) {
		slog.Error("link expired", "context", "ResetPassword", "ExpiresAt", passwordReset.ExpiresAt)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("bcrypt error", "context", "ResetPassword", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err = database.New().UpdateUserPassword(r.Context(), store.UpdateUserPasswordParams{
		ID:           passwordReset.UserID,
		PasswordHash: string(hash),
	}); err != nil {
		slog.Error("unable to update user password", "context", "ResetPassword", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ResendVerification(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*store.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	verification, err := database.New().GetLatestUserVerification(r.Context(), user.ID)
	if err != nil || verification.ID == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	go func(verification store.Verification) {
		err = mailer.Send(mailer.VerificationMessage(user.Email, verification.Token))
		if err != nil {
			slog.Error("verification email send error", "context", "ResendEmailVerification", "userID", user.ID, "error", err)
			return
		}
		slog.Info("verification email sent", "context", "ResendEmailVerification", "userID", user.ID)
	}(verification)
}
