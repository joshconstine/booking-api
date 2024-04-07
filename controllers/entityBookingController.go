package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EntityBookingController struct {
	entityBookingService services.EntityBookingService
}

func NewEntityBookingController(entityBookingService services.EntityBookingService) *EntityBookingController {
	return &EntityBookingController{entityBookingService: entityBookingService}
}

func (e *EntityBookingController) CreateEntityBooking(ctx *gin.Context) {
	var entityBooking request.CreateEntityBookingRequest
	ctx.BindJSON(&entityBooking)

	booking, err := e.entityBookingService.AttemptToCreate(entityBooking)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "Ok",
		Data:   booking,
	})
}
