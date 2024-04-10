package response

type BoatResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Occupancy int    `json:"occupancy"`
	MaxWeight int    `json:"maxWeight"`

	Timeblocks          []EntityTimeblockResponse         `json:"timeblocks"`
	Photos              []EntityPhotoResponse             `json:"photos"`
	Bookings            []EntityBookingResponse           `json:"bookings"`
	BookingCostItems    []EntityBookingCostResponse       `json:"bookingCostItems"`
	BookingDocuments    []EntityBookingDocumentResponse   `json:"bookingDocuments"`
	BookingRule         EntityBookingRuleResponse         `json:"bookingRule"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
	BookingRequests     []EntityBookingRequestResponse    `json:"bookingRequests"`
}
