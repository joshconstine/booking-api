package request

type UpdateBookingCostItemRequest struct {
	Id                uint    `json:"id"`
	BookingCostTypeId uint    `json:"bookingCostTypeId"`
	Amount            float64 `json:"amount"`
	BookingId         uint    `json:"bookingId"`
}
