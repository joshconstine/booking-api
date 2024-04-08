package response

type EntityBookingResponse struct {
	ID              uint   `json:"id"`
	BookingID       string `json:"booking_id"`
	TimeblockID     uint   `json:"timeblock_id"`
	BookingStatusID uint   `json:"booking_status_id"`
}
