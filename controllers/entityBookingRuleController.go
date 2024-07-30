package controllers

import (
	requests "booking-api/data/request"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityBookingRuleController struct {
	EntityBookingRuleService services.EntityBookingRuleService
}

func NewEntityBookingRuleController(entityBookingRuleService services.EntityBookingRuleService) *EntityBookingRuleController {
	return &EntityBookingRuleController{EntityBookingRuleService: entityBookingRuleService}
}

func (e *EntityBookingRuleController) FindByID(ctx *gin.Context) {
	entityID := ctx.Param("entityId")
	entityType := ctx.Param("entityType")

	entityIdInt, _ := strconv.Atoi(entityID)

	response, err := e.EntityBookingRuleService.FindByID(uint(entityIdInt), entityType)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)

}

func (e *EntityBookingRuleController) Update(ctx *gin.Context) {
	var request requests.UpdateEntityBookingRuleRequest
	ctx.BindJSON(&request)

	response, err := e.EntityBookingRuleService.Update(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
