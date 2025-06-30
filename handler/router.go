package handler

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewMux()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(SelfhostedDisabledMiddleware)

		r.Get("/auth/verify/{token}", Verification)
	})

	r.Route("/v1", func(r chi.Router) {
		r.Use(AssignUserToContextMiddleware)

		r.Get("/app", App)

		r.Group(func(r chi.Router) {
			r.Use(AuthenticatedMiddleware)

			r.Get("/profile", Profile)
			r.Delete("/auth/logout", Logout)

			r.Post("/monitor", CreateMonitor)

			r.Group(func(r chi.Router) {
				r.Use(SelfhostedDisabledMiddleware)

				r.Post("/auth/resend-verification", ResendVerification)
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(GuestMiddleware)

			r.Post("/setup", Setup)
			r.Post("/auth/login", LoginForm)

			r.Group(func(r chi.Router) {
				r.Use(SelfhostedDisabledMiddleware)

				r.Post("/auth/register", RegisterForm)
				r.Post("/auth/reset-password-link", ResetPasswordLink)
				r.Post("/auth/reset-password", ResetPassword)
			})
		})
	})

	r.NotFound(UI)

	return r
}
