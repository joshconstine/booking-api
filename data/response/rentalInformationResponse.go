package response

type RentalInformationResponse struct {
	ID                  uint                              `json:"id"`
	Name                string                            `json:"name"`
	Bedrooms            int                               `json:"bedrooms"`
	Bathrooms           int                               `json:"bathrooms"`
	Description         string                            `json:"description"`
	Location            LocationResponse                  `json:"location"`
	Timeblocks          []TimeblockResponse               `json:"timeblocks"`
	RentalStatus        RentalStatusResponse              `json:"rentalStatus"`
	Amenities           []AmenityResponse                 `json:"amenities"`
	Photos              []EntityPhotoResponse             `json:"photos"`
	RentalRooms         []RentalRoomResponse              `json:"rentalRooms"`
	Bookings            []EntityBookingResponse           `json:"bookings"`
	BookingCostItems    []EntityBookingCostResponse       `json:"bookingCostItems"`
	BookingDocuments    []EntityBookingDocumentResponse   `json:"bookingDocuments"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
}
