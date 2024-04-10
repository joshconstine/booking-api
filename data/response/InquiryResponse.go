package response

type InquiryResponse struct {
	ID              uint                    `json:"id"`
	Note            string                  `json:"note"`
	NumGuests       int                     `json:"numGuests"`
	EntityInquiries []EntityInquiryResponse `json:"entityInquiries"`
	User            UserResponse            `json:"user"`
}
