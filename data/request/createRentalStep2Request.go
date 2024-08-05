package request

type CreateRentalStep2Params struct {
	RentalID uint
	Rooms    []RentalRoomCreateRequest
	Success  bool
}

type CreateRentalStep2Errors struct {
	Rooms []RentalRoomCreateRequest
}
