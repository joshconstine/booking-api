package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService        services.BookingService
	bookingDetailsService services.BookingDetailsService
}

func NewBookingController(service services.BookingService, detailsService services.BookingDetailsService) *BookingController {
	return &BookingController{bookingService: service, bookingDetailsService: detailsService}
}

func (t BookingController) FindAll(ctx *gin.Context) {
	bookingResponse := t.bookingService.FindAll()
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
func convertStringToUint(s string) uint {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return uint(id)
}
func (t BookingController) FindById(ctx *gin.Context) {
	id := ctx.Param("bookingId")

	bookingResponse := t.bookingService.FindById(id)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (t BookingController) CreateBookingWithUserInformation(ctx *gin.Context) {
	var request request.CreateUserRequest
	ctx.BindJSON(&request)

	bookingResponse := t.bookingService.Create(request)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   bookingResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (t BookingController) GetDetailsForBookingID(ctx *gin.Context) {
	bookingId := ctx.Param("bookingId")
	id := convertStringToUint(bookingId)

	bookingDetailsResponse := t.bookingDetailsService.FindById(id)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingDetailsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
