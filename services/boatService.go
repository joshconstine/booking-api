package services

import responses "booking-api/data/response"

type BoatService interface {
	FindAll() []responses.BoatResponse
	FindByID(id int) responses.BoatResponse
}
