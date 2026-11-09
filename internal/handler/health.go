// Package handler contains the HTTP handlers for the devboard API.
package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthHandler reports the health status of the API.
type HealthHandler struct{}

type healthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// NewHealth creates a HealthHandler
func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	response := healthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("write response:", err)
	}
}
