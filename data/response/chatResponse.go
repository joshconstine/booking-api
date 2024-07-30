package response

type ChatResponse struct {
	ID        uint                  `json:"id"`
	AccountID uint                  `json:"accountId"`
	UserID    string                `json:"userId"`
	Messages  []ChatMessageResponse `json:"messages"`
}
