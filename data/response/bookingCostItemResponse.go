package response

type BookingCostItemResponse struct {
	ID              uint                    `json:"id"`
	Amount          float64                 `json:"amount"`
	BookingCostType BookingCostTypeResponse `json:"bookingCostType"`
	TaxRate         TaxRateResponse         `json:"taxRate"`
}
