package handler

import (
	"encoding/json"
	"microservice-app/services"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler struct to store the service instance
type Handler struct {
	service *services.Service
}

// NewHandler creates a new instance of Handler
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// GetMessageHandler handles GET requests to fetch the message
func (h *Handler) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := h.service.GetMessage()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}

// DynamicHandler example of using route variables
func (h *Handler) DynamicHandler(w http.ResponseWriter, r *http.Request) {
	// Get the route variable 'name'
	vars := mux.Vars(r)
	name := vars["name"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello " + name,
	})
}
