package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
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

	// chatResponse := t.chatService.CreateChatMessage()
	// webResponse := response.Response{
	// 	Code:   200,
	// 	Status: "Ok",
	// 	Data:   chatResponse,
	// }
	// ctx.Header("Content-Type", "application/json")
	// ctx.JSON(http.StatusOK, webResponse)
}
