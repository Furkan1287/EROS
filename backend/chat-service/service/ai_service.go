// ai_service.go - AI işlemleri
package service

import (
	"eros/shared/utils"
)

type AIService struct {
	openRouterClient *utils.OpenRouterClient
}

func NewAIService(apiKey string) *AIService {
	return &AIService{
		openRouterClient: utils.NewOpenRouterClientFromEnv(),
	}
}
