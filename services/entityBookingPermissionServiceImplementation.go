package services

import (
	"booking-api/data/request"
	"booking-api/repositories"
)

type EntityBookingPermissionServiceImplementation struct {
	EntityBookingPermissionRepository repositories.EntityBookingPermissionRepository
}

func NewEntityBookingPermissionServiceImplementation(entityBookingPermissionRepository repositories.EntityBookingPermissionRepository) EntityBookingPermissionService {
	return &EntityBookingPermissionServiceImplementation{EntityBookingPermissionRepository: entityBookingPermissionRepository}
}

func (e *EntityBookingPermissionServiceImplementation) Update(entityBookingPermission request.UpdateEntityBookingPermissionRequest) error {
	return e.EntityBookingPermissionRepository.Update(entityBookingPermission)
}
