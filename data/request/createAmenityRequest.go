package request

type CreateAmenityRequest struct {
	Name          string `json:"name"`
	AmenityTypeId uint   `json:"amenityTypeId"`
}
