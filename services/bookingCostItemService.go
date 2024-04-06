package services

import (
	requests "booking-api/data/request"

	responses "booking-api/data/response"
)

type BookingCostItemService interface {
	FindAllCostItemsForBooking(bookingId string) []responses.BookingCostItemResponse
	GetTotalCostItemsForBooking(bookingId string) float64
	Create(bookingCostItem requests.CreateBookingCostItemRequest) responses.BookingCostItemResponse
	Update(bookingCostItem requests.UpdateBookingCostItemRequest) responses.BookingCostItemResponse
	Delete(bookingCostItemId uint) bool
}
