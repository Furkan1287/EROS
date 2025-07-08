// message.go - Mesaj işlemleri
package handler

import (
	"encoding/json"
	"eros/chat-service/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageHandler struct {
	chatService *service.ChatService
}

func NewMessageHandler(chatService *service.ChatService) *MessageHandler {
	return &MessageHandler{
		chatService: chatService,
	}
}

// SendMessage - Mesaj gönder
func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var request struct {
		MatchID int    `json:"match_id"`
		UserID  int    `json:"user_id"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message, err := h.chatService.SendMessage(request.MatchID, request.UserID, request.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// GetMessages - Mesajları getir
func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["match_id"])
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	messages, err := h.chatService.GetMessages(matchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// AnalyzeConversation - Sohbet analizi
func (h *MessageHandler) AnalyzeConversation(w http.ResponseWriter, r *http.Request) {
	var request struct {
		MatchID int `json:"match_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	analysis, err := h.chatService.AnalyzeConversation(request.MatchID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}
