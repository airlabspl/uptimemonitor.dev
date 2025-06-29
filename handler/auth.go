package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"
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
		slog.Error("decode error", "context", "LoginForm ", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := validate.Struct(req); err != nil {
		slog.Error("validation error", "context", "LoginForm ", "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := database.New().GetUserByEmail(r.Context(), req.Email)
	if err != nil || user.ID == 0 {
		slog.Error("user not found", "context", "LoginForm ", "email", req.Email, "error", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		slog.Error("password mismatch", "context", "LoginForm ", "userID", user.ID)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := database.New().CreateSession(r.Context(), store.CreateSessionParams{
		Uuid:      uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour * 30),
	})
	if err != nil {
		slog.Error("session creation error", "context", "LoginForm ", "userID", user.ID, "error", err)
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
