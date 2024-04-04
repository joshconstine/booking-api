package request

type CreateAmenityTypeRequest struct {
	Name          string `json:"name"`
	AmenityTypeId uint   `json:"amenityTypeId"`
}
