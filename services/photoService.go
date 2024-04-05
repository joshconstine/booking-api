package services

import (
	"booking-api/data/response"
)

type PhotoService interface {
	FindAll() []response.PhotoResponse
}
