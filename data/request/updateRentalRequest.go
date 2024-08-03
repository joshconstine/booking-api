package request

type UpdateRentalRequest struct {
	ID          uint    `json: validate:"required"`
	Name        string  `json: validate:"required"`
	LocationID  uint    `json: validate:"required"`
	AccountID   uint    `json: validate:"required"`
	Bedrooms    uint    `json: validate:"required"`
	Bathrooms   float64 `json: validate:"required"`
	NightlyRate float64 `json: validate:"required"`
	Description string  `json: validate:"required"`
}
