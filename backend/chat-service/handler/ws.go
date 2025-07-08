// ws.go - WebSocket işlemleri
package handler

import (
	"eros/chat-service/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	chatService *service.ChatService
	upgrader    websocket.Upgrader
}

func NewWebSocketHandler(chatService *service.ChatService) *WebSocketHandler {
	return &WebSocketHandler{
		chatService: chatService,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // CORS için
			},
		},
	}
}

// HandleWebSocket - WebSocket bağlantısını yönet
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["match_id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	// WebSocket bağlantısını yükselt
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	log.Printf("WebSocket connected for match %d", matchID)

	// Mesaj dinleme döngüsü
	for {
		var message struct {
			Type    string `json:"type"`
			UserID  int    `json:"user_id"`
			Message string `json:"message"`
		}

		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		// Mesajı işle
		if message.Type == "send_message" {
			chatMessage, err := h.chatService.SendMessage(matchID, message.UserID, message.Message)
			if err != nil {
				// Hata mesajını gönder
				errorResponse := map[string]interface{}{
					"type":  "error",
					"error": err.Error(),
				}
				conn.WriteJSON(errorResponse)
				continue
			}

			// Başarılı mesajı gönder
			successResponse := map[string]interface{}{
				"type":    "message_sent",
				"message": chatMessage,
			}
			conn.WriteJSON(successResponse)
		}
	}

	log.Printf("WebSocket disconnected for match %d", matchID)
}
