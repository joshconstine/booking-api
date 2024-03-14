package response

import (
	"booking-api/models"
)

// type RentalInformtion struct {
// 	RentalID      int
// 	Name          string
// 	Description   string
// 	LocationID    int
// 	LocationName  string
// 	RentalIsClean bool
// 	Thumbnail     string
// 	Bookings      []RentalBookingDetails
// 	Timeblocks    []RentalTimeblock
// }

type RentalInformationResponse struct {
	RentalID      uint
	Name          string
	Description   string
	LocationID    uint
	LocationName  string
	RentalIsClean bool
	Thumbnail     string
	// Bookings      []models.RentalBookingDetails
	Timeblocks []models.Timeblock
}
