package request

type UpdateRentalRoomRequest struct {
	ID          uint
	Name        string
	Description string
	Floor       int
	RentalID    uint
	RoomTypeID  uint
}
