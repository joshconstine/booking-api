package repositories

import "booking-api/data/request"

type EntityBookingPermissionRepository interface {
	Update(entityBookingPermission request.UpdateEntityBookingPermissionRequest) error
}
