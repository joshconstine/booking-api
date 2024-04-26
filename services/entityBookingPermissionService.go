package services

import (
	"booking-api/data/request"
)

type EntityBookingPermissionService interface {
	Update(entityBookingPermission request.UpdateEntityBookingPermissionRequest) error
}
