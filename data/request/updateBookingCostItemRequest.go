package request

type UpdateBookingCostItemRequest struct {
	Id                uint    `json:"id"`
	BookingCostTypeId uint    `json:"bookingCostTypeId"`
	Ammount           float64 `json:"ammount"`
	BookingId         uint    `json:"bookingId"`
}
