package controllers

import (
	requests "booking-api/data/request"
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

func (controller *AmenityTypeController) Create(ctx *gin.Context) {
	createAmenityTypeRequest := requests.CreateAmenityTypeRequest{}
	err := ctx.ShouldBindJSON(&createAmenityTypeRequest)

	if err != nil {
		panic(err)
	}

	created := controller.amenityTypeService.Create(createAmenityTypeRequest)

	var webResponse response.Response
	if (created == response.AmenityTypeResponse{}) {
		webResponse = response.Response{
			Code:   400,
			Status: "Bad Request",
			Data:   "Amenity Type already exists",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	} else {

		webResponse = response.Response{
			Code:   201,
			Status: "Ok",
			Data:   created,
		}
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
