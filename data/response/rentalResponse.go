package response

type RentalResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Location    LocationResponse `json:"location"`
	Bedrooms    int              `json:"bedrooms"`
	Bathrooms   int              `json:"bathrooms"`
	Description string           `json:"description"`
}
