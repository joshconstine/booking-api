package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	locationService services.LocationService
}

func NewLocationController(locationService services.LocationService) *LocationController {
	return &LocationController{
		locationService: locationService,
	}
}

func (t LocationController) FindAll(ctx *gin.Context) {
	locations := t.locationService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   locations,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)

}

func (t LocationController) FindById(ctx *gin.Context) {
	id := ctx.Param("locationId")

	idInt, _ := strconv.Atoi(id)

	location := t.locationService.FindById(uint(idInt))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   location,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, webResponse)
}

func (t LocationController) Create(ctx *gin.Context) {
	var locationName string

	err := ctx.ShouldBindJSON(&locationName)

	if err != nil {
		panic(err)
	}

	location := t.locationService.Create(locationName)

	webResponse := response.Response{
		Code:   201,
		Status: "Created",
		Data:   location,
	}

	ctx.JSON(201, webResponse)
}
