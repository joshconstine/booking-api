package request

import "time"

type UpdateEntityBookingPermissionRequest struct {
	EntityBookingPermissionID uint
	StartTime                 time.Time
	EndTime                   time.Time
	InquiryStatusID           uint
}
