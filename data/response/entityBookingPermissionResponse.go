package response

import (
	"time"
)

type EntityBookingPermissionResponse struct {
	ID            uint                  `json:"id"`
	AccountID     uint                  `json:"accountId"`
	UserID        string                `json:"userId"`
	Entity        EntityInfoResponse    `json:"entity"`
	InquiryStatus InquiryStatusResponse `json:"inquiryStatus"`
	StartTime     time.Time             `json:"startTime"`
	EndTime       time.Time             `json:"endTime"`
}
