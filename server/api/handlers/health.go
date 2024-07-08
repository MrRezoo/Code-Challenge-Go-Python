package handlers

import (
	"encoding/json"
	"github.com/MrRezoo/code-challenge/api/helpers"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	generatedResponse := helpers.GenerateBaseResponse("Boom Boom 💥", true, http.StatusOK)
	response, err := json.Marshal(generatedResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
