package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
)

type BedRepository interface {
	Create(bed models.Bed) response.BedResponse
	Update(bed models.Bed) response.BedResponse
}
