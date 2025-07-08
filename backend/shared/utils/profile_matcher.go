package utils

import (
	"math"
)

// ProfileMatcher - Detaylı profil eşleştirme
type ProfileMatcher struct{}

// MatchScore - İki kullanıcı arasında detaylı uyumluluk skoru hesapla
func (pm *ProfileMatcher) MatchScore(user1, user2 map[string]interface{}) float64 {
	totalScore := 0.0
	maxScore := 0.0

	// 1. Yaş Uyumu (20 puan)
	ageScore := pm.calculateAgeCompatibility(user1, user2)
	totalScore += ageScore
	maxScore += 20

	// 2. Boy Uyumu (15 puan)
	heightScore := pm.calculateHeightCompatibility(user1, user2)
	totalScore += heightScore
	maxScore += 15

	// 3. Hobi Uyumu (25 puan)
	hobbyScore := pm.calculateHobbyCompatibility(user1, user2)
	totalScore += hobbyScore
	maxScore += 25

	// 4. Eğitim Uyumu (10 puan)
	educationScore := pm.calculateEducationCompatibility(user1, user2)
	totalScore += educationScore
	maxScore += 10

	// 5. Yaşam Tarzı Uyumu (15 puan)
	lifestyleScore := pm.calculateLifestyleCompatibility(user1, user2)
	totalScore += lifestyleScore
	maxScore += 15

	// 6. Ciddiyet Seviyesi Uyumu (15 puan)
	seriousnessScore := pm.calculateSeriousnessCompatibility(user1, user2)
	totalScore += seriousnessScore
	maxScore += 15

	// Yüzdelik skor hesapla
	if maxScore == 0 {
		return 0
	}

	return (totalScore / maxScore) * 100
}

// calculateAgeCompatibility - Yaş uyumu hesapla
func (pm *ProfileMatcher) calculateAgeCompatibility(user1, user2 map[string]interface{}) float64 {
	age1, ok1 := user1["age"].(float64)
	age2, ok2 := user2["age"].(float64)

	if !ok1 || !ok2 {
		return 0
	}

	ageDiff := math.Abs(age1 - age2)

	switch {
	case ageDiff <= 2:
		return 20 // Mükemmel uyum
	case ageDiff <= 5:
		return 15 // Çok iyi uyum
	case ageDiff <= 10:
		return 10 // İyi uyum
	case ageDiff <= 15:
		return 5 // Orta uyum
	default:
		return 0 // Düşük uyum
	}
}

// calculateHeightCompatibility - Boy uyumu hesapla
func (pm *ProfileMatcher) calculateHeightCompatibility(user1, user2 map[string]interface{}) float64 {
	height1, ok1 := user1["height"].(float64)
	height2, ok2 := user2["height"].(float64)

	if !ok1 || !ok2 {
		return 0
	}

	heightDiff := math.Abs(height1 - height2)

	switch {
	case heightDiff <= 5:
		return 15 // Mükemmel uyum
	case heightDiff <= 10:
		return 12 // Çok iyi uyum
	case heightDiff <= 15:
		return 8 // İyi uyum
	case heightDiff <= 20:
		return 4 // Orta uyum
	default:
		return 0 // Düşük uyum
	}
}

// calculateHobbyCompatibility - Hobi uyumu hesapla
func (pm *ProfileMatcher) calculateHobbyCompatibility(user1, user2 map[string]interface{}) float64 {
	hobbies1, ok1 := user1["hobbies"].([]interface{})
	hobbies2, ok2 := user2["hobbies"].([]interface{})

	if !ok1 || !ok2 {
		return 0
	}

	// Hobi listelerini string'e çevir
	hobbySet1 := make(map[string]bool)
	hobbySet2 := make(map[string]bool)

	for _, hobby := range hobbies1 {
		if h, ok := hobby.(string); ok {
			hobbySet1[h] = true
		}
	}

	for _, hobby := range hobbies2 {
		if h, ok := hobby.(string); ok {
			hobbySet2[h] = true
		}
	}

	// Ortak hobi sayısını hesapla
	commonHobbies := 0
	for hobby := range hobbySet1 {
		if hobbySet2[hobby] {
			commonHobbies++
		}
	}

	// Toplam hobi sayısı
	totalHobbies := len(hobbySet1) + len(hobbySet2) - commonHobbies
	if totalHobbies == 0 {
		return 0
	}

	// Ortak hobi oranı
	commonRatio := float64(commonHobbies) / float64(totalHobbies)

	return commonRatio * 25 // Maksimum 25 puan
}

// calculateEducationCompatibility - Eğitim uyumu hesapla
func (pm *ProfileMatcher) calculateEducationCompatibility(user1, user2 map[string]interface{}) float64 {
	edu1, ok1 := user1["education"].(string)
	edu2, ok2 := user2["education"].(string)

	if !ok1 || !ok2 {
		return 0
	}

	// Eğitim seviyelerini sayısal değere çevir
	eduLevel1 := pm.getEducationLevel(edu1)
	eduLevel2 := pm.getEducationLevel(edu2)

	eduDiff := math.Abs(eduLevel1 - eduLevel2)

	switch {
	case eduDiff == 0:
		return 10 // Aynı seviye
	case eduDiff == 1:
		return 8 // Yakın seviye
	case eduDiff == 2:
		return 5 // Orta fark
	default:
		return 2 // Büyük fark
	}
}

// getEducationLevel - Eğitim seviyesini sayısal değere çevir
func (pm *ProfileMatcher) getEducationLevel(education string) float64 {
	switch education {
	case "İlkokul":
		return 1
	case "Ortaokul":
		return 2
	case "Lise":
		return 3
	case "Üniversite (Önlisans)":
		return 4
	case "Üniversite (Lisans)":
		return 5
	case "Yüksek Lisans":
		return 6
	case "Doktora":
		return 7
	default:
		return 3 // Varsayılan
	}
}

// calculateLifestyleCompatibility - Yaşam tarzı uyumu hesapla
func (pm *ProfileMatcher) calculateLifestyleCompatibility(user1, user2 map[string]interface{}) float64 {
	score := 0.0

	// Sigara uyumu
	smokes1, ok1 := user1["smokes"].(bool)
	smokes2, ok2 := user2["smokes"].(bool)
	if ok1 && ok2 && smokes1 == smokes2 {
		score += 5
	}

	// İçki uyumu
	drinks1, ok1 := user1["drinks"].(bool)
	drinks2, ok2 := user2["drinks"].(bool)
	if ok1 && ok2 && drinks1 == drinks2 {
		score += 5
	}

	// İş kategorisi uyumu
	job1, ok1 := user1["job_category"].(string)
	job2, ok2 := user2["job_category"].(string)
	if ok1 && ok2 && job1 == job2 {
		score += 5
	}

	return score
}

// calculateSeriousnessCompatibility - Ciddiyet seviyesi uyumu hesapla
func (pm *ProfileMatcher) calculateSeriousnessCompatibility(user1, user2 map[string]interface{}) float64 {
	serious1, ok1 := user1["seriousness"].(float64)
	serious2, ok2 := user2["seriousness"].(float64)

	if !ok1 || !ok2 {
		return 0
	}

	seriousDiff := math.Abs(serious1 - serious2)

	switch {
	case seriousDiff <= 1:
		return 15 // Mükemmel uyum
	case seriousDiff <= 2:
		return 12 // Çok iyi uyum
	case seriousDiff <= 3:
		return 8 // İyi uyum
	case seriousDiff <= 4:
		return 4 // Orta uyum
	default:
		return 0 // Düşük uyum
	}
}
