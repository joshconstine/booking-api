package response

import (
	"time"
)

type EntityBookingPermissionResponse struct {
	ID            uint                  `json:"id"`
	AccountID     uint                  `json:"accountId"`
	UserID        uint                  `json:"userId"`
	EntityID      uint                  `json:"entityId"`
	EntityType    string                `json:"entityType"`
	InquiryStatus InquiryStatusResponse `json:"inquiryStatus"`
	StartTime     time.Time             `json:"startTime"`
	EndTime       time.Time             `json:"endTime"`
}
