package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BedTypeController struct {
	bedTypeService services.BedTypeService
}

func NewBedTypeController(service services.BedTypeService) *BedTypeController {
	return &BedTypeController{bedTypeService: service}
}

func (controller *BedTypeController) FindById(ctx *gin.Context) {

	bedTypeId := ctx.Param("bedTypeId")
	id, _ := strconv.Atoi(bedTypeId)

	bedType := controller.bedTypeService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bedType,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BedTypeController) FindAll(ctx *gin.Context) {
	bedTypes := controller.bedTypeService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bedTypes,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
