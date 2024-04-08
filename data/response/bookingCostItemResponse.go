package response

type BookingCostItemResponse struct {
	ID              uint                    `json:"id"`
	BookingID       string                  `json:"bookingId"`
	Amount          float64                 `json:"amount"`
	BookingCostType BookingCostTypeResponse `json:"bookingCostType"`
	TaxRate         TaxRateResponse         `json:"taxRate"`
}
