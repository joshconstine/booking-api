package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentalController struct {
	rentalService services.RentalService
}

func NewRentalController(rentalService services.RentalService) *RentalController {
	return &RentalController{rentalService: rentalService}
}

func (controller *RentalController) FindAll(ctx *gin.Context) {
	rentals := controller.rentalService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentals,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *RentalController) FindById(ctx *gin.Context) {
	rentalId := ctx.Param("rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rental,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalController) Create(ctx *gin.Context) {
	var request request.CreateRentalRequest
	ctx.BindJSON(&request)

	rental, err := controller.rentalService.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, rental)
}

func (controller *RentalController) Update(ctx *gin.Context) {
	var request request.UpdateRentalRequest
	ctx.BindJSON(&request)

	rental, err := controller.rentalService.Update(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, rental)
}
