package controllers

import (
	"booking-api/data/response"
	"net/http"
	"strconv"

	requests "booking-api/data/request"
	services "booking-api/services"

	"github.com/gin-gonic/gin"
)

type BookingCostItemController struct {
	BookingCostItemService services.BookingCostItemService
}

func NewBookingCostItemController(service services.BookingCostItemService) *BookingCostItemController {
	return &BookingCostItemController{BookingCostItemService: service}
}

func (controller *BookingCostItemController) FindByBookingId(ctx *gin.Context) {

	bookingId := ctx.Param("bookingId")
	id, _ := strconv.Atoi(bookingId)

	bookingCostItems := controller.BookingCostItemService.FindAllCostItemsForBooking(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingCostItems,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingCostItemController) TotalForBookingId(ctx *gin.Context) {
	bookingId := ctx.Param("bookingId")
	id, _ := strconv.Atoi(bookingId)

	total := controller.BookingCostItemService.GetTotalCostItemsForBooking(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   total,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingCostItemController) Create(ctx *gin.Context) {
	createBookingCostItemRequest := requests.CreateBookingCostItemRequest{}
	err := ctx.ShouldBindJSON(&createBookingCostItemRequest)

	if err != nil {
		panic(err)
	}

	result := controller.BookingCostItemService.Create(createBookingCostItemRequest)

	var webResponse response.Response

	if (result == response.BookingCostItemResponse{}) {
		webResponse = response.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   nil,
		}
	} else {
		webResponse = response.Response{
			Code:   http.StatusCreated,
			Status: http.StatusText(http.StatusCreated),
			Data:   result,
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(webResponse.Code, webResponse)

}
