package response

import "time"

type BookingSnapshotResponse struct {
	ID             string `json:"id"`
	UserID         string
	Name           string
	PaymentStatus  string
	StartDate      *time.Time            `json:"startDate"`
	DateRecieved   time.Time             `json:"dateRecieved"`
	Status         BookingStatusResponse `json:"status"`
	BookedEntities []EntityInfoResponse  `json:"entities"`
}
