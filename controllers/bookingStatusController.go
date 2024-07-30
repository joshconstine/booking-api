package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingStatusController struct {
	BookingStatusService services.BookingStatusService
}

func NewBookingStatusController(bookingStatusService services.BookingStatusService) *BookingStatusController {
	return &BookingStatusController{BookingStatusService: bookingStatusService}
}

func (controller *BookingStatusController) FindAll(ctx *gin.Context) {
	statusResponse := controller.BookingStatusService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   statusResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingStatusController) FindById(ctx *gin.Context) {
	statusId := ctx.Param("statusId")

	id := convertStringToUint(statusId)

	statusResponse := controller.BookingStatusService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   statusResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
