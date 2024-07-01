package response

type EntityBookingResponse struct {
	ID        uint                      `json:"id"`
	BookingID string                    `json:"booking_id"`
	Name      string                    `json:"name"`
	Timeblock EntityTimeblockResponse   `json:"timeblock"`
	Status    BookingStatusResponse     `json:"status"`
	CostItems []BookingCostItemResponse `json:"costItems"`
	Documents []BookingDocumentResponse `json:"documents"`
}
