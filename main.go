package main

import (
	"context"
	"database/sql"
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
	database.New().VerifyUser(context.Background(), store.VerifyUserParams{
		ID:              1,
		EmailVerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	addr := ":4000"

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/auth/verify/{token}", handler.Verification)

	r.Route("/v1", func(r chi.Router) {
		r.Use(handler.AssignUserToContextMiddleware)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthenticatedMiddleware)

			r.Get("/profile", handler.Profile)
			r.Delete("/auth/logout", handler.Logout)
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
