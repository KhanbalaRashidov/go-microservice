package http

import (
	"encoding/json"
	"errors"
	"go-microservice/movie/internal/controller/movie"
	"log"
	"net/http"
)

type Handler struct {
	ctrl *movie.Controller
}

// New creates a new movie HTTP handler.
func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	details, err := h.ctrl.Get(r.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
