package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"

	"github.com/google/uuid"
)

func CreateMonitor(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*store.User)

	req := struct {
		Url string `json:"url" validate:"required,url"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("decode error", "context", "CreateMonitor", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	monitorUuid := uuid.NewString()
	if err := database.New().CreateMonitor(r.Context(), store.CreateMonitorParams{
		Uuid:   monitorUuid,
		Url:    req.Url,
		UserID: user.ID,
	}); err != nil {
		slog.Error("cannot create monitor", "context", "CreateMonitor", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Uuid string `json:"uuid"`
	}{
		Uuid: monitorUuid,
	})
}
