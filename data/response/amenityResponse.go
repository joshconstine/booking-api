package response

type AmenityResponse struct {
	ID          uint                `json:"id"`
	Name        string              `json:"name"`
	AmenityType AmenityTypeResponse `json:"type"`
}
