package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AmenityTypeController struct {
	amenityTypeService services.AmenityTypeService
}

func NewAmenityTypeController(service services.AmenityTypeService) *AmenityTypeController {
	return &AmenityTypeController{amenityTypeService: service}
}

func (controller *AmenityTypeController) FindById(ctx *gin.Context) {

	amenityTypeId := ctx.Param("amenityTypeId")
	id, _ := strconv.Atoi(amenityTypeId)

	amenityType := controller.amenityTypeService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   amenityType,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AmenityTypeController) FindAll(ctx *gin.Context) {
	amenityTypes := controller.amenityTypeService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   amenityTypes,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
