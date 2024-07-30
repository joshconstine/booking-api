package request

type UpdateEntityBookingRuleRequest struct {
	EntityID                uint
	EntityType              string
	AdvertiseAtAllLocations bool
	AllowPets               bool
	AllowInstantBooking     bool
	OfferEarlyCheckIn       bool
}
