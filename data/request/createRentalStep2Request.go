package request

type CreateRentalStep2Params struct {
	RentalID uint
	Rooms    []RentalRoomCreateRequest
	Success  bool
}

type CreateRentalStep2Errors struct {
	Rooms []RentalRoomCreateRequestError
}

type RentalRoomCreateRequestError struct {
	Name        string
	Description string
	Floor       string
}
