package response

type AccountMessagesSnapshot struct {
	Notifications uint                   `json:"notifications"`
	Chats         []ChatSnapshotResponse `json:"messages"`
}
