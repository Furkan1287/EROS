// chat.go - Chat modelleri
package model

import "time"

// ChatMessage - Mesaj modeli
type ChatMessage struct {
    ID        int       `json:"id" db:"id"`
    MatchID   int       `json:"match_id" db:"match_id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Message   string    `json:"message" db:"message"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ConversationAnalysis - Sohbet analizi
type ConversationAnalysis struct {
    CommonInterests      []string `json:"common_interests"`
    CompatibilityTopics  []string `json:"compatibility_topics"`
    PotentialActivities  []string `json:"potential_activities"`
    CompatibilityScore   int      `json:"compatibility_score"`
}

// ConversationStats - Sohbet istatistikleri
type ConversationStats struct {
    MatchID       int       `json:"match_id"`
    MessageCount  int       `json:"message_count"`
    LastMessageAt time.Time `json:"last_message_at"`
} 