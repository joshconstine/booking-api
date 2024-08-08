package request

import "booking-api/data/response"

type UpdateRentalRoomRequest struct {
	ID               uint                   `json:"id"`
	Name             string                 `json:"name"`
	Description      string                 `json:"description"`
	Floor            int                    `json:"floor"`
	RentalID         uint                   `json:"rental_id"`
	RentalRoomTypeID uint                   `json:"rental_room_type_id"`
	Photos           []int                  `json:"photos"`
	Beds             []response.BedResponse `json:"beds"`
	PhotoForm        RoomPhotoFormParams    `json:"photoForm"`
}
