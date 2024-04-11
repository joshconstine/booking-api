package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type BookingCostItemRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingCostItemRepositoryImplementation(db *gorm.DB) BookingCostItemRepository {
	return &BookingCostItemRepositoryImplementation{Db: db}
}

func (t *BookingCostItemRepositoryImplementation) FindAllCostItemsForBooking(bookingId string) []response.BookingCostItemResponse {
	var bookingCostItems []models.BookingCostItem
	result := t.Db.Where("booking_id = ?", bookingId).
		Preload("BookingCostType").
		Preload("TaxRate").
		Find(&bookingCostItems)
	if result.Error != nil {
		return []response.BookingCostItemResponse{}
	}

	var bookingCostItemResponses []response.BookingCostItemResponse
	for _, bookingCostItem := range bookingCostItems {
		bookingCostItemResponses = append(bookingCostItemResponses, bookingCostItem.MapBookingCostItemToResponse())
	}
	return bookingCostItemResponses
}

func (t *BookingCostItemRepositoryImplementation) Create(bookingCostItem request.CreateBookingCostItemRequest) response.BookingCostItemResponse {
	bookingCostItemModel := models.BookingCostItem{
		BookingID:         bookingCostItem.BookingId,
		BookingCostTypeID: bookingCostItem.BookingCostTypeId,
		Amount:            bookingCostItem.Amount,
	}
	result := t.Db.Create(&bookingCostItemModel)
	if result.Error != nil {
		return response.BookingCostItemResponse{}
	}

	return bookingCostItemModel.MapBookingCostItemToResponse()

}

func (t *BookingCostItemRepositoryImplementation) GetTotalCostItemsForBooking(bookingId string) float64 {
	var total float64
	var bookingCostItems []models.BookingCostItem
	result := t.Db.Where("booking_id = ?", bookingId).Find(&bookingCostItems)
	if result.Error != nil {
		return 0
	}

	for _, bookingCostItem := range bookingCostItems {
		total += bookingCostItem.Amount
	}
	return total
}

func (t *BookingCostItemRepositoryImplementation) Update(updateRequest request.UpdateBookingCostItemRequest) response.BookingCostItemResponse {
	var bookingCostItem models.BookingCostItem
	result := t.Db.First(&bookingCostItem, updateRequest.Id)
	if result.Error != nil {
		return response.BookingCostItemResponse{}
	}

	bookingCostItem.Amount = updateRequest.Amount
	bookingCostItem.BookingCostTypeID = updateRequest.BookingCostTypeId

	result = t.Db.Save(&bookingCostItem)
	if result.Error != nil {
		return response.BookingCostItemResponse{}
	}

	return bookingCostItem.MapBookingCostItemToResponse()
}

func (t *BookingCostItemRepositoryImplementation) Delete(bookingCostItemId uint) bool {
	result := t.Db.Delete(&models.BookingCostItem{}, bookingCostItemId)
	if result.Error != nil {
		return false
	}

	return true
}
