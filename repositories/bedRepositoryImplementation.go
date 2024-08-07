package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
)

type BedRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBedRepositoryImplementation(db *gorm.DB) BedRepository {
	return &BedRepositoryImplementation{
		Db: db,
	}
}

func (r *BedRepositoryImplementation) Create(bed models.Bed) response.BedResponse {
	result := r.Db.Create(&bed)
	if result.Error != nil {
		return response.BedResponse{}
	}

	return bed.MapBedToResponse()
}

func (r *BedRepositoryImplementation) Update(bed models.Bed) response.BedResponse {
	result := r.Db.Model(&models.Bed{}).Where("id = ?", bed.ID).Update("bed_type_id", bed.BedTypeID)
	if result.Error != nil {
		return response.BedResponse{}
	}

	return bed.MapBedToResponse()
}
