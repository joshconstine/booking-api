package response

import "time"

type EntityInquiryResponse struct {
	ID            uint                  `json:"id"`
	InquiryID     uint                  `json:"inquiryID"`
	EntityID      uint                  `json:"entityID"`
	EntityType    string                `json:"entityType"`
	StartTime     time.Time             `json:"startTime"`
	EndTime       time.Time             `json:"endTime"`
	InquiryStatus InquiryStatusResponse `json:"inquiryStatus"`
}
