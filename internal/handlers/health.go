package handlers

import (
	"example/cats/db"
	"example/cats/internal/utilities"
	"log/slog"
	"net/http"
)

type HealthHandler struct {
	DB     *db.Queries
	Logger *slog.Logger
}

func (h *HealthHandler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utilities.RespondJson(w, map[string]string{"status": "system operational!"}, 200)
	}
}
