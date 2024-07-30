package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"
)

type ChatServiceImplementation struct {
	chatRepository repositories.ChatRepository
}

func NewChatServiceImplementation(chatRepository repositories.ChatRepository) ChatService {
	return &ChatServiceImplementation{chatRepository: chatRepository}
}

func (csi *ChatServiceImplementation) CreateChatMessage(request *request.CreateChatMessageRequest) (response.ChatResponse, error) {
	createdMessage, err := csi.chatRepository.CreateChatMessage(request)

	if err != nil {
		return response.ChatResponse{}, err
	}

	return createdMessage, nil

}

func (csi *ChatServiceImplementation) DeleteChatMessage(request *request.DeleteChatMessageRequest) (response.ChatResponse, error) {
	deletedMessage, err := csi.chatRepository.DeleteChatMessage(request)

	if err != nil {
		return response.ChatResponse{}, err
	}

	return deletedMessage, nil
}
