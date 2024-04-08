package response

type EntityPhotoResponse struct {
	ID    uint          `json:"id"`
	Photo PhotoResponse `json:"photo"`
}
