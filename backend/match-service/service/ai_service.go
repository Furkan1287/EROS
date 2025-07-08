// ai_service.go - AI işlemleri
package service

import (
	"encoding/json"
	"eros/match-service/model"
	"eros/shared/utils"
	"fmt"
	"time"
)

type AIService struct {
	openRouterClient *utils.OpenRouterClient
}

func NewAIService() *AIService {
	return &AIService{
		openRouterClient: utils.NewOpenRouterClientFromEnv(),
	}
}

// FindBestBlindMatch - Blind date için en uygun eşleşmeyi bul
func (s *AIService) FindBestBlindMatch(user *model.User, candidates []model.User) *model.User {
	if len(candidates) == 0 {
		return nil
	}

	bestScore := 0.0
	var bestMatch *model.User

	for _, candidate := range candidates {
		score, err := s.openRouterClient.ProfileMatching(s.userToMap(user), s.userToMap(&candidate))
		if err != nil {
			// API hatası durumunda basit skorlama kullan
			score = s.calculateSimpleCompatibilityScore(user, &candidate)
		}

		if score > bestScore {
			bestScore = score
			bestMatch = &candidate
		}
	}

	return bestMatch
}

// calculateSimpleCompatibilityScore - Basit uyumluluk skoru (API hatası durumunda)
func (s *AIService) calculateSimpleCompatibilityScore(user1, user2 *model.User) float64 {
	score := 0.0

	// Ciddiyet seviyesi uyumu
	if user1.Seriousness == user2.Seriousness {
		score += 30
	} else {
		score += 10
	}

	// Hobi uyumu
	commonHobbies := s.findCommonHobbies(user1.Hobbies, user2.Hobbies)
	score += float64(len(commonHobbies)) * 10

	// Hobi kategorisi uyumu
	commonCategories := s.findCommonHobbies(user1.HobbyCategories, user2.HobbyCategories)
	score += float64(len(commonCategories)) * 15

	// Boy uyumu (varsa)
	if user1.Height > 0 && user2.Height > 0 {
		heightDiff := abs(user1.Height - user2.Height)
		if heightDiff <= 10 {
			score += 15
		} else if heightDiff <= 20 {
			score += 10
		}
	}

	// Yaş uyumu
	if user1.Age > 0 && user2.Age > 0 {
		ageDiff := abs(user1.Age - user2.Age)
		if ageDiff <= 5 {
			score += 20
		} else if ageDiff <= 10 {
			score += 15
		} else {
			score += 5
		}
	}

	// Sigara/içki uyumu
	if user1.Smokes == user2.Smokes {
		score += 10
	}
	if user1.Drinks == user2.Drinks {
		score += 10
	}

	return score
}

// GenerateIceBreaker - Buz kırıcı mesaj oluştur
func (s *AIService) GenerateIceBreaker(user1, user2 *model.User) (string, error) {
	return s.openRouterClient.IceBreaker(s.userToMap(user1), s.userToMap(user2))
}

// GenerateDateTask - Date görevi oluştur
func (s *AIService) GenerateDateTask(user1, user2 *model.User) (*model.DateTask, error) {
	response, err := s.openRouterClient.DateSuggestion(s.userToMap(user1), s.userToMap(user2))
	if err != nil {
		return nil, err
	}

	// JSON parse et
	var task model.DateTask
	if err := json.Unmarshal([]byte(response), &task); err != nil {
		return nil, err
	}

	return &task, nil
}

// GenerateBlindDateTask - Blind date için görev oluştur
func (s *AIService) GenerateBlindDateTask(match *model.Match, messages []model.BlindMessage) (*model.DateTask, error) {
	// Mesajları analiz et
	conversation := ""
	for _, msg := range messages {
		conversation += msg.Message + " "
	}

	// Chat analysis ile sohbeti analiz et (şimdilik kullanmıyoruz)
	_, err := s.openRouterClient.ChatAnalysis(conversation)
	if err != nil {
		return nil, err
	}

	// Basit date önerisi döndür
	task := &model.DateTask{
		MatchID:     match.ID,
		Title:       "Blind Date Buluşması",
		Description: "Sohbet analizine göre uygun bir buluşma yeri",
		Location:    "İstanbul'da uygun bir mekan",
		Duration:    "2-3 saat",
		Difficulty:  "Orta",
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	return task, nil
}

// AnalyzeBlindMessage - Blind mesajı analiz et
func (s *AIService) AnalyzeBlindMessage(matchID int, message string) error {
	isSafe, err := s.openRouterClient.SecurityFilter(message)
	if err != nil {
		return err
	}

	if !isSafe {
		return fmt.Errorf("inappropriate content detected")
	}

	return nil
}

// userToMap - User'ı map'e çevir (AI için)
func (s *AIService) userToMap(user *model.User) map[string]interface{} {
	return map[string]interface{}{
		"name":             user.Name,
		"age":              user.Age,
		"bio":              user.Bio,
		"seriousness":      user.Seriousness,
		"height":           user.Height,
		"weight":           user.Weight,
		"smokes":           user.Smokes,
		"drinks":           user.Drinks,
		"job":              user.Job,
		"job_category":     user.JobCategory,
		"education":        user.Education,
		"hobbies":          user.Hobbies,
		"hobby_categories": user.HobbyCategories,
	}
}

// Yardımcı fonksiyonlar
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *AIService) findCommonHobbies(hobbies1, hobbies2 []string) []string {
	common := []string{}
	hobbyMap := make(map[string]bool)

	for _, hobby := range hobbies1 {
		hobbyMap[hobby] = true
	}

	for _, hobby := range hobbies2 {
		if hobbyMap[hobby] {
			common = append(common, hobby)
		}
	}

	return common
}
