// match_repository.go - Eşleşme veritabanı işlemleri
package repository

import (
    "database/sql"
    "eros/match-service/model"
)

type MatchRepository struct {
    db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
    return &MatchRepository{db: db}
}

// CreateMatch - Yeni eşleşme oluştur
func (r *MatchRepository) CreateMatch(match *model.Match) error {
    query := `
        INSERT INTO matches (user1_id, user2_id, match_type, status, created_at, expires_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `
    
    result, err := r.db.Exec(query, match.User1ID, match.User2ID, match.MatchType, 
                           match.Status, match.CreatedAt, match.ExpiresAt)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    match.ID = int(id)
    return nil
}

// GetMatchByID - ID ile eşleşme getir
func (r *MatchRepository) GetMatchByID(matchID int) (*model.Match, error) {
    query := `
        SELECT id, user1_id, user2_id, match_type, status, created_at, expires_at
        FROM matches WHERE id = ?
    `
    
    match := &model.Match{}
    err := r.db.QueryRow(query, matchID).Scan(
        &match.ID, &match.User1ID, &match.User2ID, &match.MatchType,
        &match.Status, &match.CreatedAt, &match.ExpiresAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return match, nil
}

// GetActiveMatch - Aktif eşleşme getir
func (r *MatchRepository) GetActiveMatch(userID int) (*model.Match, error) {
    query := `
        SELECT id, user1_id, user2_id, match_type, status, created_at, expires_at
        FROM matches 
        WHERE (user1_id = ? OR user2_id = ?) AND status = 'active'
        ORDER BY created_at DESC LIMIT 1
    `
    
    match := &model.Match{}
    err := r.db.QueryRow(query, userID, userID).Scan(
        &match.ID, &match.User1ID, &match.User2ID, &match.MatchType,
        &match.Status, &match.CreatedAt, &match.ExpiresAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    return match, nil
}

// UpdateMatchStatus - Eşleşme durumunu güncelle
func (r *MatchRepository) UpdateMatchStatus(matchID int, status string) error {
    query := `UPDATE matches SET status = ? WHERE id = ?`
    _, err := r.db.Exec(query, status, matchID)
    return err
}

// CreateSwipe - Swipe kaydet
func (r *MatchRepository) CreateSwipe(swipe *model.Swipe) error {
    query := `
        INSERT INTO swipes (user_id, target_id, direction, created_at)
        VALUES (?, ?, ?, ?)
    `
    
    result, err := r.db.Exec(query, swipe.UserID, swipe.TargetID, 
                           swipe.Direction, swipe.CreatedAt)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    swipe.ID = int(id)
    return nil
}

// CheckMutualSwipe - Karşılıklı swipe kontrolü
func (r *MatchRepository) CheckMutualSwipe(user1ID, user2ID int) (bool, error) {
    query := `
        SELECT COUNT(*) FROM swipes 
        WHERE (user_id = ? AND target_id = ? AND direction = 'right')
        AND (user_id = ? AND target_id = ? AND direction = 'right')
    `
    
    var count int
    err := r.db.QueryRow(query, user1ID, user2ID, user2ID, user1ID).Scan(&count)
    if err != nil {
        return false, err
    }
    
    return count == 2, nil
}

// CreateBlindMessage - Blind mesaj oluştur
func (r *MatchRepository) CreateBlindMessage(message *model.BlindMessage) error {
    query := `
        INSERT INTO blind_messages (match_id, user_id, message, is_ai, created_at)
        VALUES (?, ?, ?, ?, ?)
    `
    
    result, err := r.db.Exec(query, message.MatchID, message.UserID, 
                           message.Message, message.IsAI, message.CreatedAt)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    message.ID = int(id)
    return nil
}

// GetBlindMessages - Blind mesajları getir
func (r *MatchRepository) GetBlindMessages(matchID int) ([]model.BlindMessage, error) {
    query := `
        SELECT id, match_id, user_id, message, is_ai, created_at
        FROM blind_messages 
        WHERE match_id = ?
        ORDER BY created_at ASC
    `
    
    rows, err := r.db.Query(query, matchID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var messages []model.BlindMessage
    for rows.Next() {
        var msg model.BlindMessage
        err := rows.Scan(&msg.ID, &msg.MatchID, &msg.UserID, 
                        &msg.Message, &msg.IsAI, &msg.CreatedAt)
        if err != nil {
            return nil, err
        }
        messages = append(messages, msg)
    }
    
    return messages, nil
}

// CreateDateTask - Date görevi oluştur
func (r *MatchRepository) CreateDateTask(task *model.DateTask) error {
    query := `
        INSERT INTO date_tasks (match_id, title, description, location, duration, difficulty, status, created_at)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
    
    result, err := r.db.Exec(query, task.MatchID, task.Title, task.Description,
                           task.Location, task.Duration, task.Difficulty, 
                           task.Status, task.CreatedAt)
    if err != nil {
        return err
    }
    
    id, err := result.LastInsertId()
    if err != nil {
        return err
    }
    
    task.ID = int(id)
    return nil
}

// GetMatchHistory - Eşleşme geçmişi
func (r *MatchRepository) GetMatchHistory(userID int) ([]model.Match, error) {
    query := `
        SELECT id, user1_id, user2_id, match_type, status, created_at, expires_at
        FROM matches 
        WHERE (user1_id = ? OR user2_id = ?) AND status != 'active'
        ORDER BY created_at DESC
    `
    
    rows, err := r.db.Query(query, userID, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var matches []model.Match
    for rows.Next() {
        var match model.Match
        err := rows.Scan(&match.ID, &match.User1ID, &match.User2ID, 
                        &match.MatchType, &match.Status, &match.CreatedAt, &match.ExpiresAt)
        if err != nil {
            return nil, err
        }
        matches = append(matches, match)
    }
    
    return matches, nil
} 