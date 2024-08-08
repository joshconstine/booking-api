package request

type RoomPhotoFormParams struct {
	RoomId           uint   `json:"roomId"`
	ThumbnailID      uint   `json:"thumbnailId"`
	SelectedPhotoIDs []uint `json:"photoIds"`
}
