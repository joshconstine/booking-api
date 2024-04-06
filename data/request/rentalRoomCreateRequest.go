package request

type RentalRoomCreateRequest struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Floor            int    `json:"floor"`
	RentalID         uint   `json:"rental_id"`
	RentalRoomTypeID uint   `json:"rental_room_type_id"`
	Photos           []int  `json:"photos"`
	Beds             []int  `json:"beds"`
}
