package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type RoomTypeRepositoryImplementation struct {
	Db *gorm.DB
}

func NewRoomTypeRepositoryImplementation(db *gorm.DB) RoomTypeRepository {
	return &RoomTypeRepositoryImplementation{Db: db}
}

func (r *RoomTypeRepositoryImplementation) FindAll() []response.RentalRoomTypeResponse {
	var rentalRoomTypes []models.RoomType
	result := r.Db.Find(&rentalRoomTypes)
	if result.Error != nil {
		return []response.RentalRoomTypeResponse{}
	}

	var rentalRoomTypeResponses []response.RentalRoomTypeResponse
	for _, rentalRoomType := range rentalRoomTypes {
		rentalRoomTypeResponses = append(rentalRoomTypeResponses, response.RentalRoomTypeResponse{
			ID:   rentalRoomType.ID,
			Name: rentalRoomType.Name,
		})
	}

	return rentalRoomTypeResponses
}

func (r *RoomTypeRepositoryImplementation) FindById(id uint) response.RentalRoomTypeResponse {
	var rentalRoomType models.RoomType
	result := r.Db.Where("id = ?", id).First(&rentalRoomType)
	if result.Error != nil {
		return response.RentalRoomTypeResponse{}
	}

	return response.RentalRoomTypeResponse{
		ID:   rentalRoomType.ID,
		Name: rentalRoomType.Name,
	}
}

func (r *RoomTypeRepositoryImplementation) Create(rentalRoomTypeName string) response.RentalRoomTypeResponse {

	rentalRoomType := models.RoomType{
		Name: rentalRoomTypeName,
	}

	result := r.Db.Create(&rentalRoomType)
	if result.Error != nil {
		return response.RentalRoomTypeResponse{}
	}

	return response.RentalRoomTypeResponse{
		ID:   rentalRoomType.ID,
		Name: rentalRoomType.Name,
	}
}
