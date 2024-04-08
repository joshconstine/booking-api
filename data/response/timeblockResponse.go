package response

import "time"

type TimeblockResponse struct {
	ID        uint      `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	BookingID string    `json:"bookingID"`
}
