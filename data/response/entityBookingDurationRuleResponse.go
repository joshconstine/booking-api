package response

import "time"

type EntityBookingDurationRuleResponse struct {
	ID          uint      `json:"id"`
	EntityID    uint      `json:"entityId"`
	EntityType  string    `json:"entityType"`
	MinDuration int       `json:"minDuration"`
	MaxDuration int       `json:"maxDuration"`
	Buffer      int       `json:"buffer"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}
