package response

type ChatResponse struct {
	ID        uint                  `json:"id"`
	AccountID uint                  `json:"accountId"`
	UserID    uint                  `json:"userId"`
	Messages  []ChatMessageResponse `json:"messages"`
}
