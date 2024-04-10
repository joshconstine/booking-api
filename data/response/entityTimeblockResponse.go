package response

import "time"

type EntityTimeblockResponse struct {
	ID        uint      `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	BookingID string    `json:"bookingID"`
}
