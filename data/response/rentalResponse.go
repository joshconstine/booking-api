package response

type RentalResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	Location    LocationResponse `json:"location"`
	Bedrooms    uint             `json:"bedrooms"`
	Bathrooms   float64          `json:"bathrooms"`
	Description string           `json:"description"`
	Thumbnail   string           `json:"thumbnail"`
}
