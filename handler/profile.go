package handler

import (
	"encoding/json"
	"net/http"
	"selfhosted/database/store"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*store.User)
	if !ok || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Verified bool   `json:"verified"`
	}{
		Name:     user.Name,
		Email:    user.Email,
		Verified: user.EmailVerifiedAt.Valid && !user.EmailVerifiedAt.Time.IsZero(),
	})
}
