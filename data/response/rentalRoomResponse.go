package response

type RentalRoomResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Floor       int               `json:"floor"`
	RentalID    uint              `json:"rental_id"`
	RoomTypeID  uint              `json:"room_type_id"`
	Photos      []PhotoResponse   `json:"photos"`
	Beds        []BedTypeResponse `json:"beds"`
}
