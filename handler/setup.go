package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/database/store"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	adminCount, err := database.New().CountAdminUsers(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if adminCount > 0 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	request := struct {
		Name            string `json:"name" validate:"required"`
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	}{}

	json.NewDecoder(r.Body).Decode(&request)

	if err := validate.Struct(request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user, err := database.New().CreateAdminUser(r.Context(), store.CreateAdminUserParams{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = database.New().VerifyUser(r.Context(), store.VerifyUserParams{
		ID:              user.ID,
		EmailVerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		slog.Error("Failed to verify user", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	adminCount, err = database.New().CountAdminUsers(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	config.SetupFinished = adminCount > 0

	slog.Info("Created admin user", "user", user)
	w.WriteHeader(http.StatusCreated)
}
