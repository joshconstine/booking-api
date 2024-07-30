package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type RentalRoom struct {
	gorm.Model
	RentalID    uint
	Name        string
	Description string
	Floor       int
	RoomTypeID  uint
	RoomType    RoomType
	Beds        []BedType     `gorm:"many2many:rental_room_beds;"`
	Photos      []EntityPhoto `gorm:"polymorphic:Entity"`
}

func (r *RentalRoom) TableName() string {
	return "rental_rooms"
}

func (r *RentalRoom) MapRentalRoomToResponse() response.RentalRoomResponse {
	response := response.RentalRoomResponse{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Floor:       r.Floor,
	}

	response.RoomType = r.RoomType.MapRoomTypeToResponse()

	for _, bed := range r.Beds {
		response.Beds = append(response.Beds, bed.MapBedTypeToResponse())
	}

	for _, photo := range r.Photos {
		response.Photos = append(response.Photos, photo.MapEntityPhotoToResponse())
	}

	return response

}
