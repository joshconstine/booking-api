package request

type CreateEntityBookingCostRequest struct {
	EntityID          uint    `json:"entityId"`
	EntityType        string  `json:"entityType"`
	BookingCostTypeID uint    `json:"bookingCostTypeId"`
	Ammount           float64 `json:"ammount"`
	TaxRateID         uint    `json:"taxRateId"`
}
