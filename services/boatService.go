package services

import responses "booking-api/data/response"

type BoatService interface {
	FindAll() []responses.BoatResponse
	FindById(id int) responses.BoatResponse
}
