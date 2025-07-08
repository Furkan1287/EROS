// main.go - Chat Service Entry Point
package main

import (
	"eros/chat-service/handler"
	"eros/chat-service/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// Chat Service'i oluştur
	chatService := service.NewChatService()

	// Handler'ları oluştur
	messageHandler := handler.NewMessageHandler(chatService)
	wsHandler := handler.NewWebSocketHandler(chatService)

	// Router'ı oluştur
	router := mux.NewRouter()

	// Message routes
	router.HandleFunc("/api/messages/send", messageHandler.SendMessage).Methods("POST")
	router.HandleFunc("/api/messages/{match_id}", messageHandler.GetMessages).Methods("GET")
	router.HandleFunc("/api/messages/analyze", messageHandler.AnalyzeConversation).Methods("POST")

	// WebSocket route
	router.HandleFunc("/ws/{match_id}", wsHandler.HandleWebSocket)

	// CORS middleware
	router.Use(func(next http.Handler) http.Handler {
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
	})

	// Sunucuyu başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Printf("Chat Service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
