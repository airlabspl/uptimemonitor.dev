package handler

import (
	"encoding/json"
	"net/http"
	"selfhosted/config"
)

func App(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		SetupFinished bool `json:"setup_finished"`
		Selfhosted    bool `json:"selfhosted"`
	}{
		SetupFinished: config.SetupFinished,
		Selfhosted:    config.Selfhosted(),
	})
}
