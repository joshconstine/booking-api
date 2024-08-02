package request

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
	Amenities           []uint
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
