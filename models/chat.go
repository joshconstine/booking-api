package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	AccountID uint `gorm:"not null"`
	UserID    uint `gorm:"not null"`
	Messages  []ChatMessage
}

func (chat *Chat) TableName() string {
	return "chats"
}

func (chat *Chat) MapChatToResponse() response.ChatResponse {
	var messages []response.ChatMessageResponse
	for _, message := range chat.Messages {
		messages = append(messages, message.MapChatMessageToResponse())
	}
	return response.ChatResponse{
		ID:        chat.ID,
		AccountID: chat.AccountID,
		UserID:    chat.UserID,
		Messages:  messages,
	}
}
