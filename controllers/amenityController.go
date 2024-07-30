package controllers

import (
	"booking-api/services"
	"net/http"

	requests "booking-api/data/request"
	"booking-api/data/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AmenityController struct {
	amenityService services.AmenityService
}

func NewAmenityController(service services.AmenityService) *AmenityController {
	return &AmenityController{amenityService: service}
}

func (controller *AmenityController) FindById(ctx *gin.Context) {

	amenityId := ctx.Param("amenityId")
	id, _ := strconv.Atoi(amenityId)

	amenity := controller.amenityService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   amenity,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AmenityController) FindAll(ctx *gin.Context) {
	amenities := controller.amenityService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   amenities,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AmenityController) Create(ctx *gin.Context) {
	createAmenityRequest := requests.CreateAmenityRequest{}
	err := ctx.ShouldBindJSON(&createAmenityRequest)

	if err != nil {
		panic(err)
	}

	controller.amenityService.Create(createAmenityRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
