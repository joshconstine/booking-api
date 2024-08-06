package response

type RentalRoomResponse struct {
	ID          uint                  `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Floor       int                   `json:"floor"`
	RoomType    RoomTypeResponse      `json:"room_type"`
	Beds        []BedResponse         `json:"beds"`
	Photos      []EntityPhotoResponse `json:"photos"`
}
