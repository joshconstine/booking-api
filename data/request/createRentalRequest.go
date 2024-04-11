package request

type CreateRentalRequest struct {
	Name        string  `json: validate:"required"`
	LocationID  uint    `json: validate:"required"`
	AccountID   uint    `json: validate:"required"`
	Bedrooms    uint    `json: validate:"required"`
	Bathrooms   uint    `json: validate:"required"`
	NightlyRate float64 `json: validate:"required"`
	Description string  `json: validate:"required"`
}
