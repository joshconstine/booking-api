package controllers

import (
	"booking-api/services"
	"strconv"

	requests "booking-api/data/request"
	responses "booking-api/data/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

type EntityBookingDurationRuleController struct {
	EntityBookingDurationRuleService services.EntityBookingDurationRuleService
}

func NewEntityBookingDurationRuleController(entityBookingDurationRuleService services.EntityBookingDurationRuleService) *EntityBookingDurationRuleController {
	return &EntityBookingDurationRuleController{
		EntityBookingDurationRuleService: entityBookingDurationRuleService,
	}
}

func (e *EntityBookingDurationRuleController) FindByID(ctx *gin.Context) {
	entityID := ctx.Param("entityId")
	entityType := ctx.Param("entityType")

	entityIdInt, _ := strconv.Atoi(entityID)

	response := e.EntityBookingDurationRuleService.FindByID(uint(entityIdInt), entityType)

	webresponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webresponse)

}

func (e *EntityBookingDurationRuleController) Update(ctx *gin.Context) {
	var request requests.UpdateEntityBookingDurationRuleRequest
	ctx.BindJSON(&request)

	response := e.EntityBookingDurationRuleService.Update(request)

	webresponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webresponse)
}
