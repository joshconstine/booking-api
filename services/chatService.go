package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type ChatService interface {
	CreateChatMessage(message *request.CreateChatMessageRequest) (response.ChatResponse, error)
}
