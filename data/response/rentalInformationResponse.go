package response

type RentalInformationResponse struct {
	ID           uint                 `json:"id"`
	Name         string               `json:"name"`
	Bedrooms     uint                 `json:"bedrooms"`
	Bathrooms    float64              `json:"bathrooms"`
	Guests       uint                 `json:"guests"`
	Description  string               `json:"description"`
	Address      string               `json:"address"`
	Location     LocationResponse     `json:"location"`
	RentalStatus RentalStatusResponse `json:"rentalStatus"`
	RentalRooms  []RentalRoomResponse `json:"rentalRooms"`
	Amenities    []AmenityResponse    `json:"amenities"`

	Timeblocks                 []EntityTimeblockResponse             `json:"timeblocks"`
	Photos                     []EntityPhotoResponse                 `json:"photos"`
	Bookings                   []EntityBookingResponse               `json:"bookings"`
	BookingCostItems           []EntityBookingCostResponse           `json:"bookingCostItems"`
	BookingCostItemAdjustments []EntityBookingCostAdjustmentResponse `json:"bookingCostItemAdjustments"`
	BookingDocuments           []EntityBookingDocumentResponse       `json:"bookingDocuments"`
	BookingRule                EntityBookingRuleResponse             `json:"bookingRule"`
	BookingDurationRule        EntityBookingDurationRuleResponse     `json:"bookingDurationRule"`
	BookingRequests            []EntityBookingPermissionResponse     `json:"bookingRequests"`
}
