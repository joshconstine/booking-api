package request

type CreateEntityBookingCostRequest struct {
	EntityID          uint    `json:"entityId"`
	EntityType        string  `json:"entityType"`
	BookingCostTypeID uint    `json:"bookingCostTypeId"`
	Amount            float64 `json:"amount"`
	TaxRateID         uint    `json:"taxRateId"`
	TaxRatePercentage float64 `json:"taxRatePercentage"`
}
