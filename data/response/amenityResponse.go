package response

type AmenityResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	AmenityType AmenityTypeResponse `json:"type"`
}

type SortedAmenityResponse struct {
	TypeId    uint              `json:"typeId"`
	TypeName  string            `json:"typeName"`
	Amenities []AmenityResponse `json:"amenities"`
}
