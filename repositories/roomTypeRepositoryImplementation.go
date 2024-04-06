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

func (r *RoomTypeRepositoryImplementation) FindAll() []response.RoomTypeResponse {
	var rentalRoomTypes []models.RoomType
	result := r.Db.Find(&rentalRoomTypes)
	if result.Error != nil {
		return []response.RoomTypeResponse{}
	}

	var rentalRoomTypeResponses []response.RoomTypeResponse
	for _, rentalRoomType := range rentalRoomTypes {
		rentalRoomTypeResponses = append(rentalRoomTypeResponses, response.RoomTypeResponse{
			ID:   rentalRoomType.ID,
			Name: rentalRoomType.Name,
		})
	}

	return rentalRoomTypeResponses
}

func (r *RoomTypeRepositoryImplementation) FindById(id uint) response.RoomTypeResponse {
	var rentalRoomType models.RoomType
	result := r.Db.Where("id = ?", id).First(&rentalRoomType)
	if result.Error != nil {
		return response.RoomTypeResponse{}
	}

	return response.RoomTypeResponse{
		ID:   rentalRoomType.ID,
		Name: rentalRoomType.Name,
	}
}

func (r *RoomTypeRepositoryImplementation) Create(rentalRoomTypeName string) response.RoomTypeResponse {

	rentalRoomType := models.RoomType{
		Name: rentalRoomTypeName,
	}

	result := r.Db.Create(&rentalRoomType)
	if result.Error != nil {
		return response.RoomTypeResponse{}
	}

	return response.RoomTypeResponse{
		ID:   rentalRoomType.ID,
		Name: rentalRoomType.Name,
	}
}
