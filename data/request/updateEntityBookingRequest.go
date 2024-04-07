package request

import "time"

type UpdateEntityBookingRequest struct {
	ID               uint                           `json:"id"`
	BookingID        string                         `json:"bookingId"`
	EntityID         uint                           `json:"entityId"`
	EntityType       string                         `json:"entityType"`
	StartTime        time.Time                      `json:"startTime"`
	EndTime          time.Time                      `json:"endTime"`
	BookingStatusID  uint                           `json:"booking_statusId"`
	BookingCostItems []CreateBookingCostItemRequest `json:"bookingCostItems"`
}
