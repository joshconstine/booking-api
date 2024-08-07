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

	bookingCostItems := controller.BookingCostItemService.FindAllCostItemsForBooking(bookingId)

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

	total := controller.BookingCostItemService.GetTotalCostItemsForBooking(bookingId)

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

func (controller *BookingCostItemController) Update(ctx *gin.Context) {
	updateBookingCostItemRequest := requests.UpdateBookingCostItemRequest{}
	err := ctx.ShouldBindJSON(&updateBookingCostItemRequest)

	if err != nil {
		panic(err)
	}

	result := controller.BookingCostItemService.Update(updateBookingCostItemRequest)

	var webResponse response.Response

	if (result == response.BookingCostItemResponse{}) {
		webResponse = response.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   nil,
		}
	} else {
		webResponse = response.Response{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   result,
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(webResponse.Code, webResponse)

}

func (controller *BookingCostItemController) Delete(ctx *gin.Context) {
	bookingCostItemId := ctx.Param("bookingCostItemId")
	id, _ := strconv.Atoi(bookingCostItemId)

	result := controller.BookingCostItemService.Delete(uint(id))

	var webResponse response.Response

	if result {
		webResponse = response.Response{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}
	} else {
		webResponse = response.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   nil,
		}
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(webResponse.Code, webResponse)

}
