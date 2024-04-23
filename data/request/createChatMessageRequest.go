package request

type CreateChatMessageRequest struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
	ChatID  uint   `json:"chat_id"`
}
