package response

type RentalRoomResponse struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Floor       int                   `json:"floor"`
	RoomType    RoomTypeResponse      `json:"room_type"`
	Photos      []EntityPhotoResponse `json:"photos"`
	Beds        []BedTypeResponse     `json:"beds"`
}
