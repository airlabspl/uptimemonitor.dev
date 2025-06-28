package main

import (
	"net/http"
	"selfhosted/handler"

	"github.com/go-chi/chi/v5"
)

func main() {
	addr := ":4000"

	r := chi.NewMux()

	r.Group(func(r chi.Router) {

	})

	r.Route("/v1", func(r chi.Router) {
		r.Use(handler.AssignUserToContextMiddleware)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthenticatedMiddleware)

			r.Get("/profile", handler.Profile)
		})

		r.Group(func(r chi.Router) {
			r.Use(handler.GuestMiddleware)

			r.Post("/auth/login", handler.LoginForm)
		})
	})

	r.NotFound(handler.UI)

	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	s.ListenAndServe()
}
