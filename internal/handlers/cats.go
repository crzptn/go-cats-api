package handlers

import (
	"database/sql"
	"errors"
	"example/cats/db"
	"example/cats/internal/utilities"
	"log/slog"
	"net/http"
	"strconv"
)

type CatsHandler struct {
	DB     *db.Queries
	Logger *slog.Logger
}

func (h *CatsHandler) GetAllCats() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := h.DB.GetCats(r.Context())
		if err != nil {
			h.Logger.Error("failed to get cats", "error", err)

			utilities.RespondJson(w, map[string]string{
				"error": "failed to retrieve cats",
			}, http.StatusInternalServerError)
			return
		}

		utilities.RespondJson(w, cats, http.StatusOK)
	}
}

func (h *CatsHandler) ReadCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			utilities.RespondJson(w, map[string]string{
				"error": "invalid cat id",
			}, http.StatusBadRequest)
			return
		}

		cat, err := h.DB.GetCat(r.Context(), id)
		if err != nil {
			h.Logger.Error("failed to get cat", "id", id, "error", err)

			if errors.Is(err, sql.ErrNoRows) {
				utilities.RespondJson(w, map[string]string{
					"error": "cat not found",
				}, http.StatusNotFound)
				return
			}

			utilities.RespondJson(w, map[string]string{
				"error": "failed to retrieve cat",
			}, http.StatusInternalServerError)
			return
		}

		utilities.RespondJson(w, cat, http.StatusOK)
	}
}

func (h *CatsHandler) CreateCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cat db.CreateCatParams

		if err := utilities.ReadJson(r, &cat); err != nil {
			utilities.RespondJson(w, map[string]string{
				"error": "invalid request body",
			}, http.StatusBadRequest)
			return
		}

		created, err := h.DB.CreateCat(r.Context(), cat)
		if err != nil {
			h.Logger.Error("failed to create cat", "error", err)

			utilities.RespondJson(w, map[string]string{
				"error": "failed to create cat",
			}, http.StatusInternalServerError)
			return
		}

		utilities.RespondJson(w, created, http.StatusCreated)
	}
}

func (h *CatsHandler) UpdateCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			utilities.RespondJson(w, map[string]string{
				"error": "invalid cat id",
			}, http.StatusBadRequest)
			return
		}

		var cat db.UpdateCatParams

		if err := utilities.ReadJson(r, &cat); err != nil {
			utilities.RespondJson(w, map[string]string{
				"error": "invalid request body",
			}, http.StatusBadRequest)
			return
		}

		cat.ID = id

		if err := h.DB.UpdateCat(r.Context(), cat); err != nil {
			h.Logger.Error("failed to update cat", "id", id, "error", err)

			utilities.RespondJson(w, map[string]string{
				"error": "failed to update cat",
			}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *CatsHandler) DeleteCat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			utilities.RespondJson(w, map[string]string{
				"error": "invalid cat id",
			}, http.StatusBadRequest)
			return
		}

		if err := h.DB.DeleteCat(r.Context(), id); err != nil {
			h.Logger.Error("failed to delete cat", "id", id, "error", err)

			utilities.RespondJson(w, map[string]string{
				"error": "failed to delete cat",
			}, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
