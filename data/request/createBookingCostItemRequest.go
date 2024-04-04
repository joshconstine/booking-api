package request

type CreateBookingCostItemRequest struct {
	BookingCostTypeId uint    `json:"bookingCostTypeId"`
	Amount            float64 `json:"amount"`
	BookingId         uint    `json:"bookingId"`
}
