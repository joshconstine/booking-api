package response

type BookingResponse struct {
	ID               uint `json:"id"`
	UserID           uint `json:"userID"`
	BookingStatusID  uint `json:"bookingStatusID"`
	BookingDetailsID uint `json:"bookingDetailsID"`
}
