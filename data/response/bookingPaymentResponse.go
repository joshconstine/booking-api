package response

import "time"

type BookingPaymentResponse struct {
	ID            uint                  `json:"id"`
	BookingID     uint                  `json:"bookingId"`
	PaymentAmount float64               `json:"paymentAmount"`
	PaymentMethod PaymentMethodResponse `json:"paymentMethod"`
	PaymentDate   time.Time             `json:"paymentDate"`
}
