package response

import "time"

type BookingDetailsResponse struct {
	ID               uint      `json:"id"`
	BookingID        string    `json:"bookingID"`
	PaymentComplete  bool      `json:"paymentComplete"`
	DepositPaid      bool      `json:"depositPaid"`
	PaymentDueDate   time.Time `json:"paymentDueDate"`
	DocumentsSigned  bool      `json:"documentsSigned"`
	BookingStartDate time.Time `json:"bookingStartDate"`
	GuestCount       int       `json:"guestCount"`
	LocationID       uint      `json:"locationID"`
	InvoiceID        string    `json:"invoiceID"`
}
