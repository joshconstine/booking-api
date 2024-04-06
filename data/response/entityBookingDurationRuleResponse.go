package response

import "time"

type EntityBookingDurationRuleResponse struct {
	ID          uint      `json:"id"`
	EntityID    uint      `json:"entity_id"`
	EntityType  string    `json:"entity_type"`
	MinDuration int       `json:"min_duration"`
	MaxDuration int       `json:"max_duration"`
	Buffer      int       `json:"buffer"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
