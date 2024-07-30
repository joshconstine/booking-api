package request

import "time"

type CreateEntityBookingRequest struct {
	BookingID  string    `json:"bookingId"`
	EntityID   uint      `json:"entityId"`
	EntityType string    `json:"entityType"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}
