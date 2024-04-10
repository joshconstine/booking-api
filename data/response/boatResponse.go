package response

type BoatResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Occupancy uint   `json:"occupancy"`
	MaxWeight uint   `json:"maxWeight"`

	Timeblocks          []EntityTimeblockResponse         `json:"timeblocks"`
	Photos              []EntityPhotoResponse             `json:"photos"`
	BookingRule         EntityBookingRuleResponse         `json:"bookingRule"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
}
