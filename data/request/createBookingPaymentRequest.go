package request

type CreateBookingPaymentRequest struct {
	BookingID       string  `json:"bookingId"`
	PaymentAmount   float64 `json:"paymentAmount"`
	PaymentMethodID uint    `json:"paymentMethodId"`
	PaypalReference *string `json:"paypalReference"`
	PaymentDate     string  `json:"paymentDate"`
}
