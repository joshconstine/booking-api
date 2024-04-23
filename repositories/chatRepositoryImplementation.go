package repositories

import (
	"booking-api/data/request"
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

func (t *ChatRepositoryImplementation) CreateChatMessage(chatMessageRequest *request.CreateChatMessageRequest) (response.ChatResponse, error) {
	chatMessage := models.ChatMessage{
		ChatID:  chatMessageRequest.ChatID,
		Message: chatMessageRequest.Message,
		UserID:  chatMessageRequest.UserID,
	}

	result := t.Db.Model(&models.ChatMessage{}).Create(&chatMessage)
	if result.Error != nil {
		return response.ChatResponse{}, result.Error
	}

	var newChat models.Chat

	chat := t.Db.Model(&models.Chat{}).Preload("Messages").First(&newChat)

	if chat.Error != nil {
		return response.ChatResponse{}, chat.Error
	}

	return newChat.MapChatToResponse(), nil
}

func (t *ChatRepositoryImplementation) DeleteChatMessage(deleteReq *request.DeleteChatMessageRequest) (response.ChatResponse, error) {
	result := t.Db.Model(&models.ChatMessage{}).Where("id = ? AND user_id = ?", deleteReq.MessageID, deleteReq.UserID).Delete(&models.ChatMessage{})
	if result.Error != nil {
		return response.ChatResponse{}, result.Error
	}

	var newChat models.Chat

	chat := t.Db.Model(&models.Chat{}).Preload("Messages").First(&newChat)

	if chat.Error != nil {
		return response.ChatResponse{}, chat.Error
	}

	return newChat.MapChatToResponse(), nil
}
