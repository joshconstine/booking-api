package repositories

import (
	"booking-api/data/request"
	"booking-api/models"

	"gorm.io/gorm"
)

type EntityBookingPermissionRepositoryImplementation struct {
	Db *gorm.DB
}

func NewEntityBookingPermissionRepositoryImplementation(db *gorm.DB) EntityBookingPermissionRepository {
	return &EntityBookingPermissionRepositoryImplementation{Db: db}
}

func (e *EntityBookingPermissionRepositoryImplementation) Update(entityBookingPermission request.UpdateEntityBookingPermissionRequest) error {

	var entityBookingPermissionModel models.EntityBookingPermission

	e.Db.Model(&models.EntityBookingPermission{}).Where("id = ?", entityBookingPermission.EntityBookingPermissionID).First(&entityBookingPermissionModel)

	entityBookingPermissionModel.InquiryStatusID = entityBookingPermission.InquiryStatusID
	// if entityBookingPermission.StartTime != nil {
	// 	entityBookingPermissionModel.StartTime = entityBookingPermission.StartTime
	// }
	// entityBookingPermissionModel.EndTime = entityBookingPermission.EndTime

	result := e.Db.Model(&models.EntityBookingPermission{}).Where("id = ? ", entityBookingPermission.EntityBookingPermissionID).Save(&entityBookingPermissionModel)
	return result.Error
}
