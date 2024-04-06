package services

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/repositories"

	"github.com/go-playground/validator/v10"
)

type BookingCostItemServiceImplementation struct {
	bookingCostItemRepository repositories.BookingCostItemRepository
	Validate                  *validator.Validate
}

func NewBookingCostItemServiceImplementation(bookingCostItemRepository repositories.BookingCostItemRepository, validate *validator.Validate) BookingCostItemService {
	return &BookingCostItemServiceImplementation{
		bookingCostItemRepository: bookingCostItemRepository,
		Validate:                  validate,
	}
}

func (t BookingCostItemServiceImplementation) Create(bookingCostItem request.CreateBookingCostItemRequest) response.BookingCostItemResponse {
	err := t.Validate.Struct(bookingCostItem)

	if err != nil {
		panic(err)
	}

	return t.bookingCostItemRepository.Create(bookingCostItem)

}

func (t BookingCostItemServiceImplementation) FindAllCostItemsForBooking(bookingId string) []response.BookingCostItemResponse {
	result := t.bookingCostItemRepository.FindAllCostItemsForBooking(bookingId)

	return result
}

func (t BookingCostItemServiceImplementation) GetTotalCostItemsForBooking(bookingId string) float64 {
	result := t.bookingCostItemRepository.GetTotalCostItemsForBooking(bookingId)

	return result
}

func (t BookingCostItemServiceImplementation) Update(bookingCostItem request.UpdateBookingCostItemRequest) response.BookingCostItemResponse {
	err := t.Validate.Struct(bookingCostItem)

	if err != nil {
		panic(err)
	}

	return t.bookingCostItemRepository.Update(bookingCostItem)
}

func (t BookingCostItemServiceImplementation) Delete(bookingCostItemId uint) bool {
	result := t.bookingCostItemRepository.Delete(bookingCostItemId)

	return result
}
