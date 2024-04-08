package response

type RentalInformationResponse struct {
	ID                  uint                              `json:"id"`
	Name                string                            `json:"name"`
	Bedrooms            int                               `json:"bedrooms"`
	Bathrooms           int                               `json:"bathrooms"`
	Description         string                            `json:"description"`
	Location            LocationResponse                  `json:"location"`
	RentalStatus        RentalStatusResponse              `json:"rentalStatus"`
	Amenities           []AmenityResponse                 `json:"amenities"`
	Photos              []EntityPhotoResponse             `json:"photos"`
	RentalRooms         []RentalRoomResponse              `json:"rentalRooms"`
	Bookings            []EntityBookingResponse           `json:"bookings"`
	BookingCostItems    []EntityBookingCostResponse       `json:"bookingCostItems"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
}
