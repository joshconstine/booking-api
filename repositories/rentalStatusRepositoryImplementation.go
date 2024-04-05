package repositories

import (
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type RentalStatusRepositoryImplementation struct {
	Db *gorm.DB
}

func NewRentalStatusRepositoryImplementation(db *gorm.DB) RentalStatusRepository {
	return &RentalStatusRepositoryImplementation{Db: db}
}

func (r *RentalStatusRepositoryImplementation) FindAll() []responses.RentalStatusResponse {
	var rentalStatuses []models.RentalStatus
	var rentalStatusResponses []responses.RentalStatusResponse

	result := r.Db.Find(&rentalStatuses)
	if result.Error != nil {
		return []responses.RentalStatusResponse{}
	}

	for _, rentalStatus := range rentalStatuses {
		rentalStatusResponses = append(rentalStatusResponses, responses.RentalStatusResponse{
			RentalID: rentalStatus.RentalID,
			IsClean:  rentalStatus.IsClean,
		})
	}

	return rentalStatusResponses
}

func (r *RentalStatusRepositoryImplementation) FindByRentalId(rentalId uint) responses.RentalStatusResponse {
	var rentalStatus models.RentalStatus
	var rentalStatusResponse responses.RentalStatusResponse

	result := r.Db.Where("rental_id = ?", rentalId).First(&rentalStatus)
	if result.Error != nil {
		return responses.RentalStatusResponse{}
	}

	rentalStatusResponse.RentalID = rentalStatus.RentalID
	rentalStatusResponse.IsClean = rentalStatus.IsClean

	return rentalStatusResponse
}

func (r *RentalStatusRepositoryImplementation) UpdateStatusForRentalId(rentalId uint, isClean bool) responses.RentalStatusResponse {
	var rentalStatus models.RentalStatus
	var rentalStatusResponse responses.RentalStatusResponse

	result := r.Db.Where("rental_id = ?", rentalId).First(&rentalStatus)
	if result.Error != nil {
		return responses.RentalStatusResponse{}
	}

	rentalStatus.IsClean = isClean
	result = r.Db.Save(&rentalStatus)
	if result.Error != nil {
		return responses.RentalStatusResponse{}
	}

	rentalStatusResponse.RentalID = rentalStatus.RentalID
	rentalStatusResponse.IsClean = rentalStatus.IsClean

	return rentalStatusResponse
}
