package http

import (
	"net/http"
	"selfhosted/ui"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewMux()

	authHandler := NewAuthHandler()

	r.Group(func(r chi.Router) {

	})

	r.Route("/v1", func(r chi.Router) {
		r.Post("/auth/login", authHandler.LoginForm)

		r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		})
	})

	r.NotFound(NewUIHandler(ui.FS()).Handle)

	return r
}
