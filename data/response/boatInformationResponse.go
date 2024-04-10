package response

type BoatInformationResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Occupancy uint               `json:"occupancy"`
	MaxWeight uint               `json:"maxWeight"`
	Status    BoatStatusResponse `json:"status"`

	Timeblocks                 []EntityTimeblockResponse             `json:"timeblocks"`
	Photos                     []EntityPhotoResponse                 `json:"photos"`
	Bookings                   []EntityBookingResponse               `json:"bookings"`
	BookingCostItems           []EntityBookingCostResponse           `json:"bookingCostItems"`
	BookingCostItemAdjustments []EntityBookingCostAdjustmentResponse `json:"bookingCostItemAdjustments"`
	BookingDocuments           []EntityBookingDocumentResponse       `json:"bookingDocuments"`
	BookingRule                EntityBookingRuleResponse             `json:"bookingRule"`
	BookingDurationRule        EntityBookingDurationRuleResponse     `json:"bookingDurationRule"`
	BookingRequests            []EntityBookingRequestResponse        `json:"bookingRequests"`
}
