package request

type CreateBookingCostItemRequest struct {
	BookingCostTypeId uint    `json:"bookingCostTypeId"`
	Amount            float64 `json:"amount"`
	BookingId         string  `json:"bookingId"`
	EntityBookingId   uint    `json:"entityBookingId"`
}
