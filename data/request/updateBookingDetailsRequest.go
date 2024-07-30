package request

import "time"

type UpdateBookingDetailsRequest struct {
	ID               uint
	BookingID        string
	PaymentComplete  bool
	DepositPaid      bool
	PaymentDueDate   time.Time
	DocumentsSigned  bool
	BookingStartDate time.Time
	GuestCount       int
}
