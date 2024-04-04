package response

type BookingCostItemResponse struct {
	Id              uint                    `json:"id"`
	BookingId       uint                    `json:"bookingId"`
	Amount          float64                 `json:"amount"`
	BookingCostType BookingCostTypeResponse `json:"bookingCostType"`
}
