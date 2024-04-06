package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomTypeController struct {
	RoomTypeService services.RoomTypeService
}

func NewRoomTypeController(roomTypeService services.RoomTypeService) *RoomTypeController {
	return &RoomTypeController{RoomTypeService: roomTypeService}
}

func (r *RoomTypeController) FindAll(ctx *gin.Context) {
	roomTypes := r.RoomTypeService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   roomTypes,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (r *RoomTypeController) FindById(ctx *gin.Context) {
	roomTypeId := ctx.Param("roomTypeId")

	roomTypeIdInt, _ := strconv.Atoi(roomTypeId)

	roomType := r.RoomTypeService.FindById(roomTypeIdInt)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   roomType,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
