package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model
	ChatID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	Message string `gorm:"not null"`
}

func (chatMessage *ChatMessage) TableName() string {
	return "chat_messages"
}

func (chatMessage *ChatMessage) MapChatMessageToResponse() response.ChatMessageResponse {
	return response.ChatMessageResponse{
		ID:      chatMessage.ID,
		ChatID:  chatMessage.ChatID,
		UserID:  chatMessage.UserID,
		Message: chatMessage.Message,
	}
}
