package main

import (
	"context"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"
	"selfhosted/handler"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	database.New().CreateUser(context.Background(), store.CreateUserParams{
		Email:        "test@example.com",
		Name:         "Test User",
		PasswordHash: string(hash),
	})

	addr := ":4000"

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Use(handler.AssignUserToContextMiddleware)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthenticatedMiddleware)

			r.Get("/profile", handler.Profile)
		})

		r.Group(func(r chi.Router) {
			r.Use(handler.GuestMiddleware)

			r.Post("/auth/login", handler.LoginForm)
			r.Post("/auth/register", handler.RegisterForm)
		})
	})

	r.NotFound(handler.UI)

	s := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	s.ListenAndServe()
}
