package main

import (
	"eros/shared/utils"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	client := utils.NewOpenRouterClientFromEnv()

	fmt.Println("🤖 TÜM AI SERVİSLERİ TEST")
	fmt.Println("==================================================")

	// Test mesajı
	testMessage := "Merhaba! Nasılsın? Bugün hava çok güzel, birlikte bir şeyler yapalım mı?"

	// Test kullanıcıları
	user1 := map[string]interface{}{
		"name":         "Ahmet",
		"age":          28.0,
		"height":       180.0,
		"seriousness":  7.0,
		"smokes":       false,
		"drinks":       true,
		"job_category": "Teknoloji",
		"education":    "Üniversite (Lisans)",
		"hobbies":      []interface{}{"müzik", "gitar", "konser", "yazılım"},
	}

	user2 := map[string]interface{}{
		"name":         "Zeynep",
		"age":          26.0,
		"height":       165.0,
		"seriousness":  8.0,
		"smokes":       false,
		"drinks":       false,
		"job_category": "Eğitim",
		"education":    "Üniversite (Lisans)",
		"hobbies":      []interface{}{"kitap", "doğa", "yürüyüş", "yemek yapma"},
		"penis_size":   13.0,
	}

	// 1. Chat Analysis
	fmt.Println("\n1️⃣ CHAT ANALYSIS")
	analysis, err := client.ChatAnalysis(testMessage)
	if err != nil {
		fmt.Printf("❌ Hata: %v\n", err)
	} else {
		fmt.Printf("✅ Analiz: %s\n", analysis)
	}

	// 2. Date Suggestion
	fmt.Println("\n2️⃣ DATE SUGGESTION")
	suggestion, err := client.DateSuggestion(user1, user2)
	if err != nil {
		fmt.Printf("❌ Hata: %v\n", err)
	} else {
		fmt.Printf("✅ Öneri: %s\n", suggestion)
	}

	// 3. Ice Breaker
	fmt.Println("\n3️⃣ ICE BREAKER")
	iceBreaker, err := client.IceBreaker(user1, user2)
	if err != nil {
		fmt.Printf("❌ Hata: %v\n", err)
	} else {
		fmt.Printf("✅ Ice Breaker: %s\n", iceBreaker)
	}

	// 4. Security Filter
	fmt.Println("\n4️⃣ SECURITY FILTER")
	isSafe, err := client.SecurityFilter(testMessage)
	if err != nil {
		fmt.Printf("❌ Hata: %v\n", err)
	} else {
		fmt.Printf("✅ Güvenli: %t\n", isSafe)
	}

	// 5. Profile Matching
	fmt.Println("\n5️⃣ PROFILE MATCHING")
	score, err := client.ProfileMatching(user1, user2)
	if err != nil {
		fmt.Printf("❌ Hata: %v\n", err)
	} else {
		fmt.Printf("✅ Uyum Skoru: %.1f/100\n", score)
	}

	fmt.Println("\n==================================================")
	fmt.Println("🎉 TEST TAMAMLANDI!")
}
