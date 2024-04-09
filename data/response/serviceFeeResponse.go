package response

// ServiceFeeResponse is a struct that represents a service fee response.
type ServiceFeeResponse struct {
	ID                    uint    `json:"id"`
	FeePercentage         float64 `json:"feePercentage"`
	AppliesToAllCostTypes bool    `json:"appliesToAllCostTypes"`
	BookingCostTypeID     uint    `json:"bookingCostTypeID"`
}
