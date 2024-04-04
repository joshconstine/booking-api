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

func (t *BookingCostItemRepositoryImplementation) FindAllCostItemsForBooking(bookingId uint) []response.BookingCostItemResponse {
	var bookingCostItems []models.BookingCostItem
	result := t.Db.Where("booking_id = ?", bookingId).Find(&bookingCostItems)
	if result.Error != nil {
		return []response.BookingCostItemResponse{}
	}

	var bookingCostItemResponses []response.BookingCostItemResponse
	for _, bookingCostItem := range bookingCostItems {
		bookingCostItemResponses = append(bookingCostItemResponses, response.BookingCostItemResponse{
			Id:        bookingCostItem.ID,
			BookingId: bookingCostItem.BookingID,
			Amount:    bookingCostItem.Amount,
			BookingCostType: response.BookingCostTypeResponse{
				ID:   bookingCostItem.BookingCostType.ID,
				Name: bookingCostItem.BookingCostType.Name,
			},
		})
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

	return response.BookingCostItemResponse{
		Id:        bookingCostItemModel.ID,
		BookingId: bookingCostItemModel.BookingID,
		Amount:    bookingCostItemModel.Amount,
		BookingCostType: response.BookingCostTypeResponse{
			ID:   bookingCostItemModel.BookingCostType.ID,
			Name: bookingCostItemModel.BookingCostType.Name,
		},
	}

}

func (t *BookingCostItemRepositoryImplementation) GetTotalCostItemsForBooking(bookingId uint) float64 {
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
