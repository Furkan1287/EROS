// match.go - Eşleşme modelleri
package model

import "time"

// Match - Eşleşme modeli
type Match struct {
    ID        int       `json:"id" db:"id"`
    User1ID   int       `json:"user1_id" db:"user1_id"`
    User2ID   int       `json:"user2_id" db:"user2_id"`
    MatchType string    `json:"match_type" db:"match_type"` // "classic" veya "blind"
    Status    string    `json:"status" db:"status"`         // "active", "completed", "expired"
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    ExpiresAt *time.Time `json:"expires_at,omitempty" db:"expires_at"`
}

// Swipe - Swipe modeli
type Swipe struct {
    ID        int       `json:"id" db:"id"`
    UserID    int       `json:"user_id" db:"user_id"`
    TargetID  int       `json:"target_id" db:"target_id"`
    Direction string    `json:"direction" db:"direction"` // "left" veya "right"
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// BlindMessage - Blind chat mesajı
type BlindMessage struct {
    ID        int       `json:"id" db:"id"`
    MatchID   int       `json:"match_id" db:"match_id"`
    UserID    int       `json:"user_id" db:"user_id"`
    Message   string    `json:"message" db:"message"`
    IsAI      bool      `json:"is_ai" db:"is_ai"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// DateTask - Date görevi
type DateTask struct {
    ID          int    `json:"id" db:"id"`
    MatchID     int    `json:"match_id" db:"match_id"`
    Title       string `json:"title" db:"title"`
    Description string `json:"description" db:"description"`
    Location    string `json:"location" db:"location"`
    Duration    string `json:"duration" db:"duration"`
    Difficulty  string `json:"difficulty" db:"difficulty"`
    Status      string `json:"status" db:"status"` // "pending", "accepted", "completed"
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// BlindMatchStatus - Blind date durumu
type BlindMatchStatus struct {
    HasActiveMatch bool       `json:"has_active_match"`
    MatchID        int        `json:"match_id,omitempty"`
    ExpiresAt      *time.Time `json:"expires_at,omitempty"`
    MessageCount   int        `json:"message_count"`
} 