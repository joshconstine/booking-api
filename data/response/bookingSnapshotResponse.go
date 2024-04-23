package response

import "time"

type BookingSnapshotResponse struct {
	ID             string                `json:"id"`
	Status         BookingStatusResponse `json:"status"`
	StartDate      time.Time             `json:"startDate"`
	DateRecieved   time.Time             `json:"dateRecieved"`
	BookedEntities []EntityInfoResponse  `json:"entities"`
}
