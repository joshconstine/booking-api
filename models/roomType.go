package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type RoomType struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

func (r *RoomType) TableName() string {
	return "room_types"
}

func (r *RoomType) MapRoomTypeToResponse() response.RoomTypeResponse {

	return response.RoomTypeResponse{
		ID:   r.ID,
		Name: r.Name,
	}
}
