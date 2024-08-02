package request

import (
	"booking-api/data/response"
)

type CreateRentalStep1Params struct {
	Name                string
	Address             string
	Description         string
	Bedrooms            int
	Bathrooms           float32
	Guests              int
	AllowInstantBooking bool
	AllowPets           bool
	ParentProperty      bool
	Amenities           []response.AmenityResponse
	Success             bool
}

type CreateRentalStep1Errors struct {
	Name        string
	Address     string
	Bedrooms    string
	Bathrooms   string
	Guests      string
	Description string
}