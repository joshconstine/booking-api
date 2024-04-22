package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type ChatRepositoryImplementation struct {
	Db *gorm.DB
}

func NewChatRepositoryImplementation(db *gorm.DB) ChatRepository {
	return &ChatRepositoryImplementation{Db: db}
}

func (t *ChatRepositoryImplementation) FindAllForUser(userID string) []response.ChatResponse {
	var chats []models.Chat
	result := t.Db.Model(&models.Chat{}).Where("user_id = ?", userID).Find(&chats)
	if result.Error != nil {
		return []response.ChatResponse{}
	}

	var chatResponses []response.ChatResponse
	for _, chat := range chats {
		chatResponses = append(chatResponses, chat.MapChatToResponse())
	}

	return chatResponses
}

func (t *ChatRepositoryImplementation) Create(chat *models.Chat) (response.ChatResponse, error) {
	result := t.Db.Model(&models.Chat{}).Create(&chat)
	if result.Error != nil {
		return response.ChatResponse{}, result.Error
	}

	return chat.MapChatToResponse(), nil
}
