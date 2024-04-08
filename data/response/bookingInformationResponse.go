package response

type BookingInformationResponse struct {
	ID string `json:"id"`
	// User UserResponse `json:"user"`
	Status    BookingStatusResponse     `json:"status"`
	Details   BookingDetailsResponse    `json:"details"`
	CostItems []BookingCostItemResponse `json:"costItems"`
	Entities  []EntityBookingResponse   `json:"entities"`
	Payments  []BookingPaymentResponse  `json:"payments"`
	Documents []BookingDocumentResponse `json:"documents"`
}
