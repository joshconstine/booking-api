package response

type EntityBookingResponse struct {
	ID         uint                      `json:"id"`
	EntityID   uint                      `json:"entity_id"`
	EntityType string                    `json:"entity_type"`
	BookingID  string                    `json:"booking_id"`
	Name       string                    `json:"name"`
	Thumbnail  string                    `json:"thumbnail"`
	Timeblock  EntityTimeblockResponse   `json:"timeblock"`
	Status     BookingStatusResponse     `json:"status"`
	CostItems  []BookingCostItemResponse `json:"costItems"`
	Documents  []BookingDocumentResponse `json:"documents"`
}
