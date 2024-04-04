package services

import (
	requests "booking-api/data/request"

	responses "booking-api/data/response"
)

type BookingCostItemService interface {
	FindAllCostItemsForBooking(bookingId uint) []responses.BookingCostItemResponse
	GetTotalCostItemsForBooking(bookingId uint) float64
	Create(bookingCostItem requests.CreateBookingCostItemRequest) responses.BookingCostItemResponse
}
