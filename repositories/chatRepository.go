package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
)

type ChatRepository interface {
	FindAllForUser(userID string) []response.ChatResponse
	Create(chat *models.Chat) (response.ChatResponse, error)
	CreateChatMessage(chatMessage *request.CreateChatMessageRequest) (response.ChatResponse, error)
	DeleteChatMessage(deleteReq *request.DeleteChatMessageRequest) (response.ChatResponse, error)
}
