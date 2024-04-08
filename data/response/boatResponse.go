package response

type BoatResponse struct {
	ID                  uint                              `json:"id"`
	Name                string                            `json:"name"`
	Occupancy           int                               `json:"occupancy"`
	MaxWeight           int                               `json:"maxWeight"`
	Timeblocks          []TimeblockResponse               `json:"timeblocks"`
	Photos              []PhotoResponse                   `json:"photos"`
	Bookings            []EntityBookingResponse           `json:"bookings"`
	BookingCostItems    []EntityBookingCostResponse       `json:"bookingCostItems"`
	BookingDurationRule EntityBookingDurationRuleResponse `json:"bookingDurationRule"`
}
