// blind.go - Blind date işlemleri (Sürpriz mod)
package handler

import (
	"encoding/json"
	"eros/match-service/service"
	"log"
	"net/http"
	"strconv"
	"time"
)

type BlindHandler struct {
	matchService *service.MatchService
}

func NewBlindHandler(matchService *service.MatchService) *BlindHandler {
	return &BlindHandler{matchService: matchService}
}

// BlindMatchRequest - Blind date eşleştirme isteği
type BlindMatchRequest struct {
	UserID int `json:"user_id"`
}

// BlindMatchResponse - Blind date eşleştirme yanıtı
type BlindMatchResponse struct {
	IsMatched bool   `json:"is_matched"`
	MatchID   int    `json:"match_id,omitempty"`
	Message   string `json:"message"`
	ExpiresAt string `json:"expires_at,omitempty"`
}

// BlindChatRequest - Blind chat mesajı
type BlindChatRequest struct {
	MatchID int    `json:"match_id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

// BlindChatResponse - Blind chat yanıtı
type BlindChatResponse struct {
	MessageID int    `json:"message_id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	IsAI      bool   `json:"is_ai"`
}

// RequestBlindMatch - Blind date eşleştirme isteği
func (h *BlindHandler) RequestBlindMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BlindMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Kullanıcının aktif blind date'i var mı kontrol et
	hasActiveMatch, err := h.matchService.HasActiveBlindMatch(req.UserID)
	if err != nil {
		http.Error(w, "Failed to check active matches", http.StatusInternalServerError)
		return
	}

	if hasActiveMatch {
		http.Error(w, "User already has an active blind date", http.StatusBadRequest)
		return
	}

	// Blind date eşleştirmesi yap
	matchID, expiresAt, err := h.matchService.CreateBlindMatch(req.UserID)
	if err != nil {
		http.Error(w, "Failed to create blind match", http.StatusInternalServerError)
		return
	}

	response := BlindMatchResponse{
		IsMatched: matchID > 0,
		MatchID:   matchID,
		Message:   "Blind date request processed",
		ExpiresAt: expiresAt,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// SendBlindMessage - Blind chat mesajı gönder
func (h *BlindHandler) SendBlindMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BlindChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Mesajı gönder ve AI analizi yap
	messageID, err := h.matchService.SendBlindMessage(req.MatchID, req.UserID, req.Message)
	if err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	// AI buz kırıcı mesajı oluştur
	_, err = h.matchService.GenerateAIIceBreaker(req.MatchID)
	if err != nil {
		// AI mesajı başarısız olsa bile ana mesaj gönderildi
		log.Printf("Failed to generate AI message: %v", err)
	}

	response := BlindChatResponse{
		MessageID: messageID,
		Message:   req.Message,
		Timestamp: time.Now().Format(time.RFC3339),
		IsAI:      false,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetBlindMessages - Blind chat mesajlarını getir
func (h *BlindHandler) GetBlindMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	matchIDStr := r.URL.Query().Get("match_id")
	matchID, err := strconv.Atoi(matchIDStr)
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	messages, err := h.matchService.GetBlindMessages(matchID)
	if err != nil {
		http.Error(w, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

// CompleteBlindDate - Blind date'i tamamla ve date görevi öner
func (h *BlindHandler) CompleteBlindDate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	matchIDStr := r.URL.Query().Get("match_id")
	matchID, err := strconv.Atoi(matchIDStr)
	if err != nil {
		http.Error(w, "Invalid match ID", http.StatusBadRequest)
		return
	}

	// Blind date'i tamamla ve date görevi oluştur
	dateTask, err := h.matchService.CompleteBlindDate(matchID)
	if err != nil {
		http.Error(w, "Failed to complete blind date", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Blind date completed successfully",
		"date_task": dateTask,
	})
}

// GetBlindMatchStatus - Blind date durumunu getir
func (h *BlindHandler) GetBlindMatchStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	status, err := h.matchService.GetBlindMatchStatus(userID)
	if err != nil {
		http.Error(w, "Failed to get match status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
