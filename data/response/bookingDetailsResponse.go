package response

import "time"

type BookingDetailsResponse struct {
	ID               uint       `json:"id"`
	PaymentComplete  bool       `json:"paymentComplete"`
	DepositPaid      bool       `json:"depositPaid"`
	PaymentDueDate   *time.Time `json:"paymentDueDate"`
	DocumentsSigned  bool       `json:"documentsSigned"`
	BookingStartDate *time.Time `json:"bookingStartDate"`
	LocationID       uint       `json:"locationID"`
	InvoiceID        string     `json:"invoiceID"`
}
