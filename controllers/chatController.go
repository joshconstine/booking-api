package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
	inbox "booking-api/view/inbox"
	"booking-api/view/ui"
	"fmt"
	"net/http"
	"strconv"
)

type ChatController struct {
	chatService services.ChatService
	userService services.UserService
}

func NewChatController(chatService services.ChatService, userService services.UserService) *ChatController {
	return &ChatController{chatService: chatService, userService: userService}
}

func (t ChatController) HandleChatMessageCreate(w http.ResponseWriter, r *http.Request) error {

	chatID := r.FormValue("chat_id")
	chatIDInt, _ := strconv.Atoi(chatID)

	params := request.CreateChatMessageRequest{

		Message: r.FormValue("message"),
		ChatID:  uint(chatIDInt),
	}

	fmt.Println(params)
	user := GetAuthenticatedUser(r)

	params.UserID = user.User.UserID

	createMessageRequest := request.CreateChatMessageRequest{
		UserID:  user.User.UserID,
		Message: params.Message,
		ChatID:  params.ChatID,
	}

	createdChat, err := t.chatService.CreateChatMessage(&createMessageRequest)

	if err != nil {
		return err
	}

	return render(r, w, ui.Chat(createdChat))
}
func (t ChatController) HandleChatMessageDelete(w http.ResponseWriter, r *http.Request) error {

	chatID := r.URL.Query().Get("messageID")

	chatIDInt, _ := strconv.Atoi(chatID)

	user := GetAuthenticatedUser(r)

	params := request.DeleteChatMessageRequest{
		MessageID: uint(chatIDInt),
		UserID:    user.User.UserID,
	}

	chat, err := t.chatService.DeleteChatMessage(&params)

	if err != nil {
		return err
	}

	return render(r, w, ui.Chat(chat))
}

func (t ChatController) HandleChatIndex(w http.ResponseWriter, r *http.Request) error {
	// user := GetAuthenticatedUser(r)
	// user.User = t.userService.FindByUserID(user.User.UserID)
	return render(r, w, inbox.Index())
}
