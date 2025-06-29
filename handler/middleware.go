package handler

import (
	"context"
	"log/slog"
	"net/http"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/database/store"
)

func AssignUserToContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session")
		if err != nil || c.Value == "" {
			next.ServeHTTP(w, r)
			return
		}

		session, err := database.New().GetSessionByUUID(r.Context(), c.Value)
		if err != nil || session.ID == 0 {
			slog.Error("session not found", "context", "AssignUserToContextMiddleware", "cookie", c.Value, "error", err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), "user", &session.User)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*store.User)
		if !ok || user == nil {
			slog.Error("user not found in context", "context", "AuthenticatedMiddleware")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GuestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SelfhostedDisabledMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Selfhosted() {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
