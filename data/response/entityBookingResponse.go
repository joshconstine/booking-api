package response

type EntityBookingResponse struct {
	ID               uint                      `json:"id"`
	BookingID        string                    `json:"booking_id"`
	EntityID         uint                      `json:"entity_id"`
	EntityType       string                    `json:"entity_type"`
	TimeblockID      uint                      `json:"timeblock_id"`
	BookingStatusID  uint                      `json:"booking_status_id"`
	BookingCostItems []BookingCostItemResponse `json:"booking_cost_items"`
}
