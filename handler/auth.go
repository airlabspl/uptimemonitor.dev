package handler

import (
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
	}{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	r.ParseForm()

	if err := validate.Struct(req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := database.New().GetUserByEmail(r.Context(), req.Email)
	if err != nil || user.ID == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	passwordHash, err := user.PasswordHash.Value()
	if err != nil || passwordHash == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(passwordHash.(string)), []byte(req.Password)) != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	session, err := database.New().CreateSession(r.Context(), store.CreateSessionParams{
		Uuid:      uuid.NewString(),
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour * 30),
	})
	if err != nil {
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
