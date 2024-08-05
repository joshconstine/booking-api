package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type RentalRoomRepositoryImplementation struct {
	Db *gorm.DB
}

func NewRentalRoomRepositoryImplementation(db *gorm.DB) RentalRoomRepository {
	return &RentalRoomRepositoryImplementation{Db: db}
}

func (r *RentalRoomRepositoryImplementation) FindAll() []response.RentalRoomResponse {
	var rentalRooms []models.RentalRoom
	result := r.Db.Preload("Photos.Photo").Preload("Beds").Find(&rentalRooms)
	if result.Error != nil {
		return []response.RentalRoomResponse{}
	}

	var rentalRoomResponses []response.RentalRoomResponse

	for _, rentalRoom := range rentalRooms {
		rentalRoomResponses = append(rentalRoomResponses, rentalRoom.MapRentalRoomToResponse())
	}

	return rentalRoomResponses
}

func (r *RentalRoomRepositoryImplementation) FindById(id uint) response.RentalRoomResponse {
	var rentalRoom models.RentalRoom
	result := r.Db.Where("id = ?", id).First(&rentalRoom)
	if result.Error != nil {
		return response.RentalRoomResponse{}
	}
	return rentalRoom.MapRentalRoomToResponse()
}
func (r *RentalRoomRepositoryImplementation) FindByRentalId(rentalId uint) []response.RentalRoomResponse {
	var rentalRoom []models.RentalRoom
	result := r.Db.Preload("Photos.Photo").Preload("Beds").Where("rental_id = ?", rentalId).Find(&rentalRoom)
	if result.Error != nil {
		return []response.RentalRoomResponse{}
	}
	var rentalRoomResponses []response.RentalRoomResponse
	for _, room := range rentalRoom {
		rentalRoomResponses = append(rentalRoomResponses, room.MapRentalRoomToResponse())
	}
	return rentalRoomResponses
}

func (r *RentalRoomRepositoryImplementation) Create(rentalRoom request.RentalRoomCreateRequest) response.RentalRoomResponse {
	rentalRoomModel := models.RentalRoom{
		Name:        rentalRoom.Name,
		Description: rentalRoom.Description,
		Floor:       rentalRoom.Floor,
		RentalID:    rentalRoom.RentalID,
		RoomTypeID:  rentalRoom.RentalRoomTypeID,
	}

	for _, photo := range rentalRoom.Photos {
		rentalRoomModel.Photos = append(rentalRoomModel.Photos, models.EntityPhoto{
			PhotoID: uint(photo),
		})

	}

	for _, bed := range rentalRoom.Beds {
		rentalRoomModel.Beds = append(rentalRoomModel.Beds, models.BedType{
			Model: gorm.Model{
				ID: uint(bed),
			},
		})
	}

	result := r.Db.Create(&rentalRoomModel)
	if result.Error != nil {
		return response.RentalRoomResponse{}
	}

	return rentalRoomModel.MapRentalRoomToResponse()
}

func (r *RentalRoomRepositoryImplementation) Update(rentalRoom request.UpdateRentalRoomRequest) response.RentalRoomResponse {
	var rentalRoomModel models.RentalRoom
	result := r.Db.Where("id = ?", rentalRoom.ID).First(&rentalRoomModel)
	if result.Error != nil {
		return response.RentalRoomResponse{}
	}

	rentalRoomModel.Name = rentalRoom.Name
	rentalRoomModel.Description = rentalRoom.Description
	rentalRoomModel.Floor = rentalRoom.Floor
	rentalRoomModel.RoomTypeID = rentalRoom.RoomTypeID

	result = r.Db.Save(&rentalRoomModel)
	if result.Error != nil {
		return response.RentalRoomResponse{}
	}

	return rentalRoomModel.MapRentalRoomToResponse()
}
