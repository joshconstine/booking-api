package response

type RentalResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Location    LocationResponse `json:"location"`
	Bedrooms    uint             `json:"bedrooms"`
	Bathrooms   uint             `json:"bathrooms"`
	Description string           `json:"description"`
}
