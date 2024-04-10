package response

type InquiryResponse struct {
	ID              uint   `json:"id"`
	Note            string `json:"note"`
	NumGuests       int    `json:"numGuests"`
	BookingRequests []EntityBookingRequestResponse
	User            UserResponse `json:"user"`
}
