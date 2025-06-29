package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"
	"selfhosted/mailer"
	"time"

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
		ctx, _ := context.WithTimeout(context.Background(), time.Minute)
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
