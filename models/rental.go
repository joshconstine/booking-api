package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Rental struct {
	gorm.Model
	Name                string
	LocationID          uint
	Bedrooms            int
	Bathrooms           int
	Description         string
	Location            Location
	RentalStatus        RentalStatus
	Amenities           []Amenity                 `gorm:"many2many:rental_amenities;"`
	Timeblocks          []Timeblock               `gorm:"polymorphic:Entity"`
	Photos              []EntityPhoto             `gorm:"polymorphic:Entity"`
	RentalRooms         []RentalRoom              `gorm:"foreignKey:RentalID"`
	Bookings            []EntityBooking           `gorm:"polymorphic:Entity"`
	BookingCostItems    []EntityBookingCost       `gorm:"polymorphic:Entity"`
	BookingDurationRule EntityBookingDurationRule `gorm:"polymorphic:Entity"`
}

func (r *Rental) TableName() string {
	return "rentals"
}

func (r *Rental) MapRentalsToResponse() response.RentalResponse {
	return response.RentalResponse{
		ID:          r.ID,
		Name:        r.Name,
		Location:    r.Location.MapLocationToResponse(),
		Bedrooms:    r.Bedrooms,
		Bathrooms:   r.Bathrooms,
		Description: r.Description,
	}
}
