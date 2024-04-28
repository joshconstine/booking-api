package services

type InvoiceService interface {
	CreateInvoiceForBooking(bookingID string) (string, error)
}
