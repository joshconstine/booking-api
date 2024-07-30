package controllers

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentalStatusController struct {
	rentalStatusService services.RentalStatusService
}

func NewRentalStatusController(service services.RentalStatusService) *RentalStatusController {
	return &RentalStatusController{rentalStatusService: service}
}

func (controller *RentalStatusController) FindAll(ctx *gin.Context) {
	rentalStatuses := controller.rentalStatusService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentalStatuses,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalStatusController) FindByRentalId(ctx *gin.Context) {
	rentalId := ctx.Param("rentalId")

	rentalIdInt, _ := strconv.Atoi(rentalId)

	rentalStatus := controller.rentalStatusService.FindByRentalId(uint(rentalIdInt))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentalStatus,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalStatusController) UpdateStatusForRentalId(ctx *gin.Context) {
	var updateRentalStatusRequest requests.UpdateRentalStatusRequest
	var response response.Response
	err := ctx.ShouldBindJSON(&updateRentalStatusRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Status = "Bad Request"
		response.Data = err.Error()
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	rentalStatus := controller.rentalStatusService.UpdateStatusForRentalId(updateRentalStatusRequest.RentalID, updateRentalStatusRequest.IsClean)

	response.Code = http.StatusOK
	response.Status = "Ok"
	response.Data = rentalStatus

	ctx.JSON(http.StatusOK, response)

}
