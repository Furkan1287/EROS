// chat_service.go - Chat işlemleri
package service

import (
	"encoding/json"
	"eros/chat-service/model"
	"eros/shared/utils"
	"fmt"
	"time"
)

type ChatService struct {
	aiService *utils.OpenRouterClient
}

func NewChatService() *ChatService {
	return &ChatService{
		aiService: utils.NewOpenRouterClientFromEnv(),
	}
}

// SendMessage - Mesaj gönder ve AI analizi yap
func (s *ChatService) SendMessage(matchID int, userID int, message string) (*model.ChatMessage, error) {
	// Güvenlik filtresi
	isSafe, err := s.aiService.SecurityFilter(message)
	if err != nil {
		return nil, fmt.Errorf("security check failed: %v", err)
	}

	if !isSafe {
		return nil, fmt.Errorf("inappropriate content detected")
	}

	// Mesajı kaydet
	chatMessage := &model.ChatMessage{
		MatchID:   matchID,
		UserID:    userID,
		Message:   message,
		CreatedAt: time.Now(),
	}

	// TODO: Database'e kaydet
	// if err := s.messageRepo.CreateMessage(chatMessage); err != nil {
	//     return nil, err
	// }

	// AI analizi (asenkron)
	go s.analyzeMessage(matchID, message)

	return chatMessage, nil
}

// GetMessages - Mesajları getir
func (s *ChatService) GetMessages(matchID int) ([]model.ChatMessage, error) {
	// TODO: Database'den getir
	// return s.messageRepo.GetMessagesByMatchID(matchID)
	return []model.ChatMessage{}, nil
}

// AnalyzeConversation - Sohbet analizi
func (s *ChatService) AnalyzeConversation(matchID int) (*model.ConversationAnalysis, error) {
	messages, err := s.GetMessages(matchID)
	if err != nil {
		return nil, err
	}

	// Sohbeti birleştir
	conversation := ""
	for _, msg := range messages {
		conversation += msg.Message + " "
	}

	// AI analizi
	analysis, err := s.aiService.ChatAnalysis(conversation)
	if err != nil {
		return nil, err
	}

	// JSON parse et
	var result model.ConversationAnalysis
	if err := json.Unmarshal([]byte(analysis), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GenerateIceBreaker - Buz kırıcı mesaj oluştur
func (s *ChatService) GenerateIceBreaker(user1, user2 map[string]interface{}) (string, error) {
	return s.aiService.IceBreaker(user1, user2)
}

// analyzeMessage - Mesaj analizi (asenkron)
func (s *ChatService) analyzeMessage(matchID int, message string) {
	// Mesaj içeriğini analiz et
	isSafe, err := s.aiService.SecurityFilter(message)
	if err != nil {
		// Log error
		return
	}

	if !isSafe {
		// TODO: Mesajı engelle ve kullanıcıyı uyar
		return
	}

	// TODO: Mesaj analizi sonuçlarını kaydet
}

// GetConversationStats - Sohbet istatistikleri
func (s *ChatService) GetConversationStats(matchID int) (*model.ConversationStats, error) {
	messages, err := s.GetMessages(matchID)
	if err != nil {
		return nil, err
	}

	stats := &model.ConversationStats{
		MatchID:       matchID,
		MessageCount:  len(messages),
		LastMessageAt: time.Now(),
	}

	if len(messages) > 0 {
		stats.LastMessageAt = messages[len(messages)-1].CreatedAt
	}

	return stats, nil
}
