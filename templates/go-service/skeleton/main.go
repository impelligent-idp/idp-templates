package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Service name and version from template
const (
	ServiceName    = "${{values.name}}"
	ServiceVersion = "1.0.0"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// InfoResponse represents the service info response
type InfoResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Owner       string `json:"owner"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "${{values.port}}"
	}

	router := mux.NewRouter()
	
	// Health and readiness endpoints
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/ready", readyHandler).Methods("GET")
	
	// Service info
	router.HandleFunc("/info", infoHandler).Methods("GET")
	
	// Your API endpoints go here
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/hello", helloHandler).Methods("GET")
	
	// Logging middleware
	router.Use(loggingMiddleware)
	
	log.Printf("Starting %s v%s on port %s", ServiceName, ServiceVersion, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Service:   ServiceName,
		Version:   ServiceVersion,
		Timestamp: time.Now(),
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	// Add readiness checks here (database, dependencies, etc.)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := InfoResponse{
		Name:        ServiceName,
		Description: "${{values.description}}",
		Version:     ServiceVersion,
		Owner:       "${{values.owner}}",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from " + ServiceName,
		"version": ServiceVersion,
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

