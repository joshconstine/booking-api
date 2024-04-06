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
	result := r.Db.Preload("EntityPhotos").Find(&rentalRooms)
	if result.Error != nil {
		return []response.RentalRoomResponse{}
	}

	var rentalRoomResponses []response.RentalRoomResponse
	var photos []response.PhotoResponse
	var beds []response.BedTypeResponse
	for _, rentalRoom := range rentalRooms {

		for _, photo := range rentalRoom.Photos {
			photos = append(photos, response.PhotoResponse{
				ID:  photo.ID,
				URL: photo.Photo.URL,
			})

		}

		for _, bed := range rentalRoom.Beds {
			beds = append(beds, response.BedTypeResponse{
				ID:   bed.ID,
				Name: bed.Name,
			})
		}

		rentalRoomResponses = append(rentalRoomResponses, response.RentalRoomResponse{
			ID:       rentalRoom.ID,
			Name:     rentalRoom.Name,
			RentalID: rentalRoom.RentalID,
			Photos:   photos,
			Beds:     beds,
		})
		photos = nil
		beds = nil
	}

	return rentalRoomResponses
}

func (r *RentalRoomRepositoryImplementation) FindById(id uint) response.RentalRoomResponse {
	var rentalRoom models.RentalRoom
	result := r.Db.Where("id = ?", id).First(&rentalRoom)
	if result.Error != nil {
		return response.RentalRoomResponse{}
	}

	var photos []response.PhotoResponse
	var beds []response.BedTypeResponse

	for _, photo := range rentalRoom.Photos {
		photos = append(photos, response.PhotoResponse{
			ID:  photo.ID,
			URL: photo.Photo.URL,
		})
	}

	for _, bed := range rentalRoom.Beds {
		beds = append(beds, response.BedTypeResponse{
			ID:   bed.ID,
			Name: bed.Name,
		})
	}

	return response.RentalRoomResponse{
		ID:       rentalRoom.ID,
		Name:     rentalRoom.Name,
		RentalID: rentalRoom.RentalID,
		Photos:   photos,
		Beds:     beds,
	}
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

	return response.RentalRoomResponse{
		ID:       rentalRoomModel.ID,
		Name:     rentalRoomModel.Name,
		RentalID: rentalRoomModel.RentalID,
	}
}
