package request

import "booking-api/data/response"

type CreateRentalStep2Params struct {
	RentalID uint
	Rooms    []response.RentalRoomResponse
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
