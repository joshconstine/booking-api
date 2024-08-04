package controllers

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"booking-api/view/ui"
	"github.com/go-chi/chi/v5"
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

func (controller *RentalStatusController) ToggleCleanStatusForRental(w http.ResponseWriter, r *http.Request) error {

	var updateRentalStatusRequest requests.UpdateRentalStatusRequest

	var rentalID = chi.URLParam(r, "rentalId")

	// Convert the rentalID to an integer
	rentalIDInt, _ := strconv.Atoi(rentalID)

	updateRentalStatusRequest.RentalID = uint(rentalIDInt)

	currentStatus := controller.rentalStatusService.FindByRentalId(updateRentalStatusRequest.RentalID)
	updateRentalStatusRequest.IsClean = !currentStatus.IsClean

	rentalStatus := controller.rentalStatusService.UpdateStatusForRentalId(updateRentalStatusRequest.RentalID, updateRentalStatusRequest.IsClean)
	return ui.RentalStatusBadge(rentalStatus, uint(rentalIDInt)).Render(r.Context(), w)
}
