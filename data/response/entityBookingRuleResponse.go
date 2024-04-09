package response

type EntityBookingRuleResponse struct {
	ID                      uint
	AdvertiseAtAllLocations bool
	AllowPets               bool
	AllowInstantBooking     bool
	OfferEarlyCheckIn       bool
}
