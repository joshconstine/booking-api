package services

import (
	requests "booking-api/data/request"
	"github.com/stripe/stripe-go/v78"

	responses "booking-api/data/response"
)

type BookingCostItemService interface {
	FindAllCostItemsForBooking(bookingId string) []responses.BookingCostItemResponse
	FindAllCheckoutItemsForBooking(bookingId string) []*stripe.CheckoutSessionLineItemParams
	GetTotalCostItemsForBooking(bookingId string) float64
	Create(bookingCostItem requests.CreateBookingCostItemRequest) responses.BookingCostItemResponse
	Update(bookingCostItem requests.UpdateBookingCostItemRequest) responses.BookingCostItemResponse
	Delete(bookingCostItemId uint) bool
}
