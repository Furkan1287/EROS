// photo.go - Kullanıcı fotoğraf modeli
package model

import "time"

type Photo struct {
    ID          int       `json:"id" db:"id"`
    UserID      int       `json:"user_id" db:"user_id"`
    URL         string    `json:"url" db:"url"`
    IsPrimary   bool      `json:"is_primary" db:"is_primary"`
    OrderIndex  int       `json:"order_index" db:"order_index"` // 1-6 arası sıralama
    AIScore     float64   `json:"ai_score" db:"ai_score"`       // AI tarafından verilen kalite skoru
    IsVerified  bool      `json:"is_verified" db:"is_verified"` // Kullanıcının kendisi olup olmadığı
    CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
