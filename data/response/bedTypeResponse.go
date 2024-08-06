package response

type BedTypeResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
type BedResponse struct {
	ID        uint   `json:"id"`
	BedTypeID uint   `json:"bed_type_id"`
	Name      string `json:"name"`
}
