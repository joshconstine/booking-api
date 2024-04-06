package models

import (
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
