package http

import (
	"encoding/json"
	"errors"
	"go-microservice/metadata/internal/controller/metadata"
	"go-microservice/metadata/internal/repository"
	"log"
	"net/http"
)

type Handler struct {
	controller *metadata.Controller
}

func NewHandler(controller *metadata.Controller) *Handler {
	return &Handler{controller: controller}
}

func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	m, err := h.controller.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
