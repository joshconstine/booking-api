package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	inbox "booking-api/view/inbox"
	"booking-api/view/ui"
	"fmt"
	"net/http"
	"strconv"
)

type ChatController struct {
	chatService    services.ChatService
	userService    services.UserService
	accountService services.AccountService
}

func NewChatController(chatService services.ChatService, userService services.UserService, accountService services.AccountService) *ChatController {
	return &ChatController{chatService: chatService, userService: userService, accountService: accountService}
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
	user := GetAuthenticatedUser(r)
	var inquiries response.AccountInquiriesSnapshot
	var messages response.AccountMessagesSnapshot

	userAccountRoles, err := t.accountService.GetUserAccountRoles(user.User.UserID)
	if err != nil {
		return err
	}

	uniqueAccountIDs := []uint{}

	for _, role := range userAccountRoles {
		unique := true
		for _, id := range uniqueAccountIDs {
			if id == role.AccountID {
				unique = false
				break
			}
		}
		if unique {
			uniqueAccountIDs = append(uniqueAccountIDs, role.AccountID)
		}
	}

	for _, accountID := range uniqueAccountIDs {
		accinquiries, err := t.accountService.GetInquiriesSnapshot(accountID)
		if err != nil {
			return err
		}
		accmessages, err := t.accountService.GetMessagesSnapshot(accountID)
		if err != nil {
			return err
		}

		inquiries.Inquiries = append(inquiries.Inquiries, accinquiries.Inquiries...)
		inquiries.Notifications += accinquiries.Notifications

		messages.Chats = append(messages.Chats, accmessages.Chats...)
		messages.Notifications += accmessages.Notifications

	}

	return render(r, w, inbox.Index(
		inquiries,
		messages,
	))
}
