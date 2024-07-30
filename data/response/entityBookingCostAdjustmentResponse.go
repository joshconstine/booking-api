package response

import "time"

type EntityBookingCostAdjustmentResponse struct {
	ID              uint      `json:"id"`
	Amount          float64   `json:"amount"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	BookingCostType BookingCostTypeResponse
	TaxRate         TaxRateResponse
}
