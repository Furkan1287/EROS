// openrouter_client.go - OpenRouter API Client
package utils

import (
	"bytes"
	"encoding/json"
	"eros/shared/types"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OpenRouterClient struct {
	BaseURL string
	Keys    map[string]string // model -> key
}

type OpenRouterRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenRouterResponse struct {
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func NewOpenRouterClientFromEnv() *OpenRouterClient {
	apiKey := os.Getenv("OPENROUTER_API_KEY")
	return &OpenRouterClient{
		BaseURL: "https://openrouter.ai/api/v1/chat/completions",
		Keys: map[string]string{
			"google/gemma-3-27b-it:free":  apiKey,
			"google/gemma-3n-e4b-it:free": apiKey,
		},
	}
}

// ChatAnalysis - Sohbet analizi (mistralai/mistral-7b-instruct)
func (c *OpenRouterClient) ChatAnalysis(conversation string) (string, error) {
	prompt := fmt.Sprintf(`
    Bu sohbeti analiz et ve şu bilgileri çıkar:
    
    Sohbet: %s
    
    Analiz:
    1. Ortak ilgi alanları neler?
    2. Hangi konularda uyum var?
    3. Potansiyel date aktiviteleri neler olabilir?
    4. Genel uyumluluk skoru (1-10)
    
    JSON formatında döndür:
    {
        "common_interests": ["hobi1", "hobi2"],
        "compatibility_topics": ["konu1", "konu2"],
        "potential_activities": ["aktivite1", "aktivite2"],
        "compatibility_score": 8
    }
    `, conversation)

	return c.callAPI(types.ModelChatAnalysis, prompt, 0.7)
}

// DateSuggestion - Date önerisi (google/gemma-7b-it)
func (c *OpenRouterClient) DateSuggestion(user1, user2 map[string]interface{}) (string, error) {
	prompt := fmt.Sprintf(`
    İki kişi için İstanbul'da eğlenceli bir date önerisi oluştur:
    
    Kişi 1: %v
    Kişi 2: %v
    
    Öneri şu formatta olsun:
    {
        "title": "Başlık",
        "description": "Açıklama",
        "location": "İstanbul'da bir yer",
        "duration": "2-3 saat",
        "difficulty": "Kolay/Orta/Zor",
        "cost": "Ücretsiz/Uygun/Pahalı",
        "why_perfect": "Neden bu ikili için ideal"
    }
    `, user1, user2)

	return c.callAPI(types.ModelDateSuggestion, prompt, 0.8)
}

// SecurityFilter - Keyword-based güvenlik filtresi
func (c *OpenRouterClient) SecurityFilter(message string) (bool, error) {
	// Yeni keyword-based filter kullan
	return SecurityFilter(message), nil
}

// IceBreaker - Buz kırıcı mesaj (huggingfaceh4/zephyr-7b-beta)
func (c *OpenRouterClient) IceBreaker(user1, user2 map[string]interface{}) (string, error) {
	prompt := fmt.Sprintf(`
    İki kişi arasında doğal ve samimi bir buz kırıcı mesaj oluştur:
    
    Kişi 1: %v
    Kişi 2: %v
    
    Mesaj:
    - Kısa ve etkili olsun (1-2 cümle)
    - Ortak ilgi alanlarına odaklan
    - Doğal ve samimi ton kullan
    - Soru içersin
    - Türkçe yaz
    `, user1, user2)

	return c.callAPI(types.ModelIceBreaker, prompt, 0.9)
}

// ProfileMatching - Profil eşleştirme (Yeni algoritma)
func (c *OpenRouterClient) ProfileMatching(user1, user2 map[string]interface{}) (float64, error) {
	matcher := &ProfileMatcher{}
	score := matcher.MatchScore(user1, user2)
	return score, nil
}

// callAPI - Genel API çağrısı
func (c *OpenRouterClient) callAPI(model, prompt string, temperature float64) (string, error) {
	apiKey := c.Keys[model]
	if apiKey == "" {
		return "", fmt.Errorf("API key for model %s not found", model)
	}
	request := OpenRouterRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens:   500,
		Temperature: temperature,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("request marshal error: %v", err)
	}

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("request creation error: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "https://eros-app.com")
	req.Header.Set("X-Title", "EROS Dating App")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("API call error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("response read error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var response OpenRouterResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("response unmarshal error: %v", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return response.Choices[0].Message.Content, nil
}
