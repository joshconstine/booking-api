package response

type BookingInformationResponse struct {
	ID string `json:"id"`
	// User UserResponse `json:"user"`
	Status    BookingStatusResponse     `json:"status"`
	Details   BookingDetailsResponse    `json:"details"`
	CostItems []BookingCostItemResponse `json:"costItems"`
	Payments  []BookingPaymentResponse  `json:"payments"`
	Documents []BookingDocumentResponse `json:"documents"`
}
