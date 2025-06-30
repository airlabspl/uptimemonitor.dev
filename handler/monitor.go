package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"

	"github.com/google/uuid"
)

type MonitorDTO struct {
	Uuid string `json:"uuid"`
	Url  string `json:"url"`
}

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

	if err := validate.Struct(req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		Uuid string `json:"uuid"`
	}{
		Uuid: monitorUuid,
	})
}

func ListMonitors(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*store.User)

	monitors, err := database.New().ListMonitors(r.Context(), user.ID)
	if err != nil {
		slog.Error("cannot list monitors", "context", "ListMonitors", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := []MonitorDTO{}
	for _, m := range monitors {
		data = append(data, MonitorDTO{
			Uuid: m.Uuid,
			Url:  m.Url,
		})
	}

	json.NewEncoder(w).Encode(struct {
		Monitors []MonitorDTO `json:"monitors"`
	}{
		Monitors: data,
	})
}
