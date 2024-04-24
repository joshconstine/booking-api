package response

type InquirySnapshotResponse struct {
	ChatID             uint   `json:"chat_id"`
	Message            string `json:"message"`
	Name               string `json:"name"`
	PermissionRequests []EntityBookingPermissionResponse
}
