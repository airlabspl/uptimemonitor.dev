package main

import (
	"context"
	"net/http"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/handler"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	adminCount, err := database.New().CountAdminUsers(context.Background())
	if err != nil {
		panic(err)
	}

	config.SetupFinished = adminCount > 0

	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(handler.SelfhostedDisabledMiddleware)

		r.Get("/auth/verify/{token}", handler.Verification)
	})

	r.Route("/v1", func(r chi.Router) {
		r.Use(handler.AssignUserToContextMiddleware)

		r.Get("/app", handler.App)

		r.Group(func(r chi.Router) {
			r.Use(handler.AuthenticatedMiddleware)

			r.Get("/profile", handler.Profile)
			r.Delete("/auth/logout", handler.Logout)
			r.Post("/auth/resend-verification", handler.ResendVerification)
		})

		r.Group(func(r chi.Router) {
			r.Use(handler.GuestMiddleware)

			r.Post("/setup", handler.Setup)
			r.Post("/auth/login", handler.LoginForm)

			r.Group(func(r chi.Router) {
				r.Use(handler.SelfhostedDisabledMiddleware)

				r.Post("/auth/register", handler.RegisterForm)
				r.Post("/auth/reset-password-link", handler.ResetPasswordLink)
				r.Post("/auth/reset-password", handler.ResetPassword)
			})
		})
	})

	r.NotFound(handler.UI)

	s := &http.Server{
		Addr:    config.Addr(),
		Handler: r,
	}

	s.ListenAndServe()
}
