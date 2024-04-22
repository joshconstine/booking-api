package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type ChatRepository interface {
	FindAllForUser(userID string) []response.ChatResponse
	Create(chat *models.Chat) (response.ChatResponse, error)
}
