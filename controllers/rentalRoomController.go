package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RentalRoomController struct {
	rentalRoomService services.RentalRoomService
}

func NewRentalRoomController(rentalRoomService services.RentalRoomService) *RentalRoomController {
	return &RentalRoomController{rentalRoomService: rentalRoomService}
}

func (controller *RentalRoomController) FindAll(ctx *gin.Context) {
	rentalRooms := controller.rentalRoomService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentalRooms,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *RentalRoomController) FindById(ctx *gin.Context) {
	rentalRoomId := ctx.Param("rentalRoomId")
	id, _ := strconv.Atoi(rentalRoomId)

	rentalRoom := controller.rentalRoomService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentalRoom,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalRoomController) Create(ctx *gin.Context) {
	var rentalRoomCreateRequest request.RentalRoomCreateRequest

	err := ctx.ShouldBindJSON(&rentalRoomCreateRequest)

	if err != nil {
		webResponse := response.Response{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	rentalRoom, err := controller.rentalRoomService.Create(rentalRoomCreateRequest)

	if err != nil {
		webResponse := response.Response{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   rentalRoom,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalRoomController) Update(ctx *gin.Context) {
	var updateRentalRoomRequest request.UpdateRentalRoomRequest

	err := ctx.ShouldBindJSON(&updateRentalRoomRequest)

	if err != nil {
		webResponse := response.Response{
			Code:   400,
			Status: "Bad Request",
			Data:   err.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	rentalRoom, err := controller.rentalRoomService.Update(updateRentalRoomRequest)

	if err != nil {
		webResponse := response.Response{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   rentalRoom,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
