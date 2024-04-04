package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingCostItemRepository interface {
	FindAllCostItemsForBooking(bookingId uint) []response.BookingCostItemResponse
	GetTotalCostItemsForBooking(bookingId uint) float64
	Create(bookingCostItem requests.CreateBookingCostItemRequest) response.BookingCostItemResponse
	// Update(bookingCostItem requests.UpdateBookingCostItemRequest) response.BookingCostItemResponse
	// Delete(bookingCostItemId uint) bool
}
