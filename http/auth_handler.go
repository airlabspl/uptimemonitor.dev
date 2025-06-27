package http

import (
	"net/http"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "some-session-token",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // todo
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
}
