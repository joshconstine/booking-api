package request

import "time"

type UpdateEntityBookingCostAdjustmentRequest struct {
	ID                uint      `json:"id"`
	EntityID          uint      `json:"entityId"`
	EntityType        string    `json:"entityType"`
	BookingCostTypeID uint      `json:"bookingCostTypeId"`
	Amount            float64   `json:"amount"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	TaxRateID         uint      `json:"taxRateId"`
}
