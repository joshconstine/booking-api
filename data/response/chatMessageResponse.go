package response

import "time"

type ChatMessageResponse struct {
	ID      uint   `json:"id"`
	ChatID  uint   `json:"chatId"`
	UserID  uint   `json:"userId"`
	Message string `json:"message"`
	Sent    time.Time
}
