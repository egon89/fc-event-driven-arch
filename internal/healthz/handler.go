package healthz

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type HealthzHandler struct{}

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

func (h *HealthzHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.healthz)

	return r
}

func (h *HealthzHandler) healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"status": "ok",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
