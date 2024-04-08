package response

type EntityBookingCostResponse struct {
	ID              uint    `json:"id"`
	Amount          float64 `json:"ammount"`
	BookingCostType BookingCostTypeResponse
	TaxRate         TaxRateResponse
}
