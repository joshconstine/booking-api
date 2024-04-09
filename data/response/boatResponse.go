package response

type BoatResponse struct {
	ID                  uint                              `json:"id"`
	Name                string                            `json:"name"`
	Occupancy           int                               `json:"occupancy"`
	MaxWeight           int                               `json:"maxWeight"`
	Timeblocks          []TimeblockResponse               `json:"timeblocks"`
	Photos              []EntityPhotoResponse             `json:"photos"`
	Bookings            []EntityBookingResponse           `json:"bookings"`
	BookingCostItems    []EntityBookingCostResponse       `json:"bookingCostItems"`
	BookingDocuments    []EntityBookingDocumentResponse   `json:"bookingDocuments"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
}
