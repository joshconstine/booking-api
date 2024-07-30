package response

type ChatSnapshotResponse struct {
	ChatID  uint   `json:"chat_id"`
	Message string `json:"message"`
	Name    string `json:"name"`
	Sent    string `json:"sent"`
}

type InquirySnapshotResponse struct {
	Chat               ChatSnapshotResponse `json:"chats"`
	PermissionRequests []EntityBookingPermissionResponse
}
