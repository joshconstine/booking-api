package response

type BoatResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Occupancy int    `json:"occupancy"`
	MaxWeight int    `json:"maxWeight"`
	Photos    []BoatPhotoResponse
}
