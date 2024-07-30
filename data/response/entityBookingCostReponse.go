package response

type EntityBookingCostResponse struct {
	ID              uint    `json:"id"`
	Amount          float64 `json:"amount"`
	BookingCostType BookingCostTypeResponse
	TaxRate         TaxRateResponse
}
