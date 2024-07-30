package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingCostTypeController struct {
	BookingStatusService services.BookingCostTypeService
}

func NewBookingCostTypeController(bookingCostTypeService services.BookingCostTypeService) *BookingCostTypeController {
	return &BookingCostTypeController{BookingStatusService: bookingCostTypeService}
}

func (controller *BookingCostTypeController) FindAll(ctx *gin.Context) {
	statusResponse := controller.BookingStatusService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   statusResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingCostTypeController) FindById(ctx *gin.Context) {
	costTypeId := ctx.Param("costTypeId")

	id := convertStringToUint(costTypeId)

	statusResponse := controller.BookingStatusService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   statusResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
