package repositories

import (
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type BedTypeRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBedTypeRepositoryImplementation(db *gorm.DB) BedTypeRepository {
	return &BedTypeRepositoryImplementation{Db: db}
}

func (t *BedTypeRepositoryImplementation) FindAll() []response.BedTypeResponse {
	var bedTypes []models.BedType
	var response []response.BedTypeResponse

	result := t.Db.Find(&bedTypes)
	if result.Error != nil {
		return []responses.BedTypeResponse{}
	}

	for _, bedType := range bedTypes {
		response = append(response, responses.BedTypeResponse{
			ID:   bedType.ID,
			Name: bedType.Name,
		})
	}

	return response
}

func (t *BedTypeRepositoryImplementation) FindById(id uint) response.BedTypeResponse {
	var bedType models.BedType
	var response response.BedTypeResponse

	result := t.Db.Where("id = ?", id).First(&bedType)
	if result.Error != nil {
		return response
	}

	response.ID = bedType.ID
	response.Name = bedType.Name

	return response
}

func (t *BedTypeRepositoryImplementation) Create(bedType models.BedType) response.BedTypeResponse {
	result := t.Db.Create(&bedType)
	if result.Error != nil {
		return response.BedTypeResponse{}
	}

	return response.BedTypeResponse{
		ID:   bedType.ID,
		Name: bedType.Name,
	}
}
