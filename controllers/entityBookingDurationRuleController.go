package controllers

import (
	"booking-api/services"
	"strconv"

	requests "booking-api/data/request"

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

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)

}

func (e *EntityBookingDurationRuleController) Update(ctx *gin.Context) {
	var request requests.UpdateEntityBookingDurationRuleRequest
	ctx.BindJSON(&request)

	response, err := e.EntityBookingDurationRuleService.Update(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
