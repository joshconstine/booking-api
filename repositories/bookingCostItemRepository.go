package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
)

type BookingCostItemRepository interface {
	FindAllCostItemsForBooking(bookingId string) []response.BookingCostItemResponse
	GetTotalCostItemsForBooking(bookingId string) float64
	Create(bookingCostItem requests.CreateBookingCostItemRequest) response.BookingCostItemResponse
	Update(bookingCostItem requests.UpdateBookingCostItemRequest) response.BookingCostItemResponse
	Delete(bookingCostItemId uint) bool
}
