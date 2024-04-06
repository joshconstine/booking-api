package response

import "time"

type BookingDetailsResponse struct {
	ID               uint      `json:"id"`
	BookingID        string    `json:"bookingID"`
	PaymentComplete  bool      `json:"paymentComplete"`
	PaymentDueDate   time.Time `json:"paymentDueDate"`
	DocumentsSigned  bool      `json:"documentsSigned"`
	BookingStartDate time.Time `json:"bookingStartDate"`
	InvoiceID        *string   `json:"invoiceID"`
}
