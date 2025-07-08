// match_service.go - Eşleşme iş mantığı
package service

import (
    "errors"
    "eros/match-service/model"
    "eros/match-service/repository"
    "time"
)

type MatchService struct {
    matchRepo *repository.MatchRepository
    userRepo  *repository.UserRepository
    aiService *AIService
}

func NewMatchService(matchRepo *repository.MatchRepository, userRepo *repository.UserRepository, aiService *AIService) *MatchService {
    return &MatchService{
        matchRepo: matchRepo,
        userRepo:  userRepo,
        aiService: aiService,
    }
}

// ProcessSwipe - Swipe işlemini gerçekleştir
func (s *MatchService) ProcessSwipe(userID, targetID int, direction string) (bool, error) {
    // Swipe kaydını oluştur
    swipe := &model.Swipe{
        UserID:    userID,
        TargetID:  targetID,
        Direction: direction,
        CreatedAt: time.Now(),
    }

    if err := s.matchRepo.CreateSwipe(swipe); err != nil {
        return false, err
    }

    // Eğer sağa kaydırma ise, karşılıklı swipe kontrolü yap
    if direction == "right" {
        return s.checkMutualSwipe(userID, targetID)
    }

    return false, nil
}

// checkMutualSwipe - Karşılıklı swipe kontrolü
func (s *MatchService) checkMutualSwipe(userID, targetID int) (bool, error) {
    // Karşılıklı swipe kontrolü
    isMutual, err := s.matchRepo.CheckMutualSwipe(userID, targetID)
    if err != nil {
        return false, err
    }

    if isMutual {
        // Eşleşme oluştur
        match := &model.Match{
            User1ID:   userID,
            User2ID:   targetID,
            MatchType: "classic",
            Status:    "active",
            CreatedAt: time.Now(),
        }

        if err := s.matchRepo.CreateMatch(match); err != nil {
            return false, err
        }

        return true, nil
    }

    return false, nil
}

// GetPotentialMatches - Potansiyel eşleşmeleri getir
func (s *MatchService) GetPotentialMatches(userID, limit int) ([]model.User, error) {
    user, err := s.userRepo.GetUserByID(userID)
    if err != nil {
        return nil, err
    }

    return s.userRepo.GetPotentialMatches(user, limit)
}

// GetMatchHistory - Eşleşme geçmişini getir
func (s *MatchService) GetMatchHistory(userID int) ([]model.Match, error) {
    return s.matchRepo.GetMatchHistory(userID)
}

// CreateBlindMatch - Blind date eşleştirmesi oluştur
func (s *MatchService) CreateBlindMatch(userID int) (int, string, error) {
    // Uygun eşleşme bul
    targetUser, err := s.findBlindMatch(userID)
    if err != nil {
        return 0, "", err
    }

    if targetUser == nil {
        return 0, "", errors.New("no suitable blind match found")
    }

    // Blind match oluştur
    match := &model.Match{
        User1ID:   userID,
        User2ID:   targetUser.ID,
        MatchType: "blind",
        CreatedAt: time.Now(),
        ExpiresAt: func() *time.Time { t := time.Now().Add(72 * time.Hour); return &t }(), // 3 gün
    }

    if err := s.matchRepo.CreateMatch(match); err != nil {
        return 0, "", err
    }

    return match.ID, match.ExpiresAt.Format(time.RFC3339), nil
}

// findBlindMatch - Blind date için uygun eşleşme bul
func (s *MatchService) findBlindMatch(userID int) (*model.User, error) {
    user, err := s.userRepo.GetUserByID(userID)
    if err != nil {
        return nil, err
    }

    // Kullanıcının tercihlerine göre potansiyel eşleşmeleri getir
    candidates, err := s.userRepo.GetPotentialMatches(user, 50)
    if err != nil {
        return nil, err
    }

    // AI skorlama algoritması ile en uygun eşleşmeyi bul
    bestMatch := s.aiService.FindBestBlindMatch(user, candidates)
    return bestMatch, nil
}

// HasActiveBlindMatch - Aktif blind date kontrolü
func (s *MatchService) HasActiveBlindMatch(userID int) (bool, error) {
    match, err := s.matchRepo.GetActiveMatch(userID)
    if err != nil {
        return false, err
    }
    return match != nil && match.MatchType == "blind", nil
}

// SendBlindMessage - Blind chat mesajı gönder
func (s *MatchService) SendBlindMessage(matchID, userID int, message string) (int, error) {
    // Mesajı kaydet
    chatMessage := &model.BlindMessage{
        MatchID:  matchID,
        UserID:   userID,
        Message:  message,
        CreatedAt: time.Now(),
    }

    if err := s.matchRepo.CreateBlindMessage(chatMessage); err != nil {
        return 0, err
    }

    // AI analizi yap
    go s.aiService.AnalyzeBlindMessage(matchID, message)

    return chatMessage.ID, nil
}

// GetBlindMessages - Blind chat mesajlarını getir
func (s *MatchService) GetBlindMessages(matchID int) ([]model.BlindMessage, error) {
    return s.matchRepo.GetBlindMessages(matchID)
}

// GenerateAIIceBreaker - AI buz kırıcı mesajı oluştur
func (s *MatchService) GenerateAIIceBreaker(matchID int) (string, error) {
    match, err := s.matchRepo.GetMatchByID(matchID)
    if err != nil {
        return "", err
    }

    user1, err := s.userRepo.GetUserByID(match.User1ID)
    if err != nil {
        return "", err
    }

    user2, err := s.userRepo.GetUserByID(match.User2ID)
    if err != nil {
        return "", err
    }

    return s.aiService.GenerateIceBreaker(user1, user2)
}

// CompleteBlindDate - Blind date'i tamamla
func (s *MatchService) CompleteBlindDate(matchID int) (*model.DateTask, error) {
    match, err := s.matchRepo.GetMatchByID(matchID)
    if err != nil {
        return nil, err
    }

    // Mesajları analiz et
    messages, err := s.matchRepo.GetBlindMessages(matchID)
    if err != nil {
        return nil, err
    }

    // AI ile date görevi oluştur
    dateTask, err := s.aiService.GenerateBlindDateTask(match, messages)
    if err != nil {
        return nil, err
    }

    // Match'i tamamlandı olarak işaretle
    match.Status = "completed"
    if err := s.matchRepo.UpdateMatchStatus(match.ID, "completed"); err != nil {
        return nil, err
    }

    return dateTask, nil
}

// GetBlindMatchStatus - Blind date durumunu getir
func (s *MatchService) GetBlindMatchStatus(userID int) (*model.BlindMatchStatus, error) {
    match, err := s.matchRepo.GetActiveMatch(userID)
    if err != nil {
        return nil, err
    }
    
    if match == nil || match.MatchType != "blind" {
        return &model.BlindMatchStatus{
            HasActiveMatch: false,
        }, nil
    }
    
    messages, err := s.matchRepo.GetBlindMessages(match.ID)
    if err != nil {
        return nil, err
    }
    
    return &model.BlindMatchStatus{
        HasActiveMatch: true,
        MatchID:        match.ID,
        ExpiresAt:      match.ExpiresAt,
        MessageCount:   len(messages),
    }, nil
}

// GenerateDateTask - Date görevi oluştur
func (s *MatchService) GenerateDateTask(userID, targetID int) (*model.DateTask, error) {
    user1, err := s.userRepo.GetUserByID(userID)
    if err != nil {
        return nil, err
    }

    user2, err := s.userRepo.GetUserByID(targetID)
    if err != nil {
        return nil, err
    }

    return s.aiService.GenerateDateTask(user1, user2)
} 