// main.go - API Gateway Entry Point
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	Name string
	URL  string
}

var services = map[string]ServiceConfig{
	"user": {
		Name: "User Service",
		URL:  "http://localhost:8081",
	},
	"match": {
		Name: "Match Service",
		URL:  "http://localhost:8082",
	},
	"chat": {
		Name: "Chat Service",
		URL:  "http://localhost:8083",
	},
}

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	router := mux.NewRouter()

	// CORS middleware
	router.Use(corsMiddleware)

	// Health check
	router.HandleFunc("/health", healthCheck).Methods("GET")

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// User service routes
	apiRouter.PathPrefix("/auth").Handler(createReverseProxy("user"))
	apiRouter.PathPrefix("/users").Handler(createReverseProxy("user"))
	apiRouter.PathPrefix("/photos").Handler(createReverseProxy("user"))
	apiRouter.PathPrefix("/form").Handler(createReverseProxy("user"))

	// Match service routes
	apiRouter.PathPrefix("/swipe").Handler(createReverseProxy("match"))
	apiRouter.PathPrefix("/matches").Handler(createReverseProxy("match"))
	apiRouter.PathPrefix("/blind").Handler(createReverseProxy("match"))

	// Chat service routes
	apiRouter.PathPrefix("/messages").Handler(createReverseProxy("chat"))
	apiRouter.PathPrefix("/ws").Handler(createReverseProxy("chat"))

	// Sunucuyu başlat
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func createReverseProxy(serviceName string) http.Handler {
	service := services[serviceName]
	targetURL, _ := url.Parse(service.URL)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request
		log.Printf("Routing request to %s: %s %s", service.Name, r.Method, r.URL.Path)

		// Forward request
		proxy.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":   "healthy",
		"services": services,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
