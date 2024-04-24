package response

type AccountInquiriesSnapshot struct {
	Notifications uint                      `json:"notifications"`
	Inquiries     []InquirySnapshotResponse `json:"inquiries"`
}
