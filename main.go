package main

import (
	"log"
	handler "microservice-app/handlers"
	"microservice-app/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new service
	svc := &services.Service{}

	// Create a new handler with the service instance
	hndlr := handler.NewHandler(svc)

	// Initialize the Gorilla Mux router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1><a href='/message'>click to messages</a></h1>"))
	})
	router.HandleFunc("/message", hndlr.GetMessageHandler).Methods("GET")
	router.HandleFunc("/greet/{name}", hndlr.DynamicHandler).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
