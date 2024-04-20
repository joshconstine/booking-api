package request

import (
	"booking-api/models"
	"time"
)

type CreateEntityTimeblockRequest struct {
	EntityType      string
	EntityBookingID uint
	EntityID        uint
	StartDate       time.Time
	EndDate         time.Time
}

func (cet CreateEntityTimeblockRequest) MapCreateEntityTimeblockRequestToEntityTimeblock() models.EntityTimeblock {
	return models.EntityTimeblock{
		EntityType:      cet.EntityType,
		EntityBookingID: cet.EntityBookingID,
		EntityID:        cet.EntityID,
		StartTime:       cet.StartDate,
		EndTime:         cet.EndDate,
	}
}
