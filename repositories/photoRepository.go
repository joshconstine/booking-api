package repositories

import "booking-api/data/response"

type PhotoRepository interface {
	FindAll() []response.PhotoResponse
}
