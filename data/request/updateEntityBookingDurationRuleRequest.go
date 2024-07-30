package request

import "time"

type UpdateEntityBookingDurationRuleRequest struct {
	EntityID    uint      `json:"entityId" validate:"required"`
	EntityType  string    `json:"entityType" validate:"required"`
	MinDuration int       `json:"minDuration" `
	MaxDuration int       `json:"maxDuration" `
	Buffer      int       `json:"buffer" `
	StartTime   time.Time `json:"startTime" `
	EndTime     time.Time `json:"endTime"`
}
