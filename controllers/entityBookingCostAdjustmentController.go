package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityBookingCostAdjustmentController struct {
	EntityBookingCostAdjustmentService services.EntityBookingCostAdjustmentService
}

func NewEntityBookingCostAdjustmentController(entityBookingCostAdjustmentService services.EntityBookingCostAdjustmentService) *EntityBookingCostAdjustmentController {
	return &EntityBookingCostAdjustmentController{EntityBookingCostAdjustmentService: entityBookingCostAdjustmentService}
}

func (e *EntityBookingCostAdjustmentController) FindAllForEntity(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")

	entityIdInt, _ := strconv.Atoi(entityID)

	response := e.EntityBookingCostAdjustmentService.FindAllForEntity(uint(entityIdInt), entityType)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}

func (e *EntityBookingCostAdjustmentController) FindAllForEntityAndRange(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")
	startDate := c.Param("startDate")
	endDate := c.Param("endDate")

	entityIdInt, _ := strconv.Atoi(entityID)

	response := e.EntityBookingCostAdjustmentService.FindAllForEntityAndRange(uint(entityIdInt), entityType, startDate, endDate)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}

func (e *EntityBookingCostAdjustmentController) Create(c *gin.Context) {
	var request request.CreateEntityBookingCostAdjustmentRequest
	c.BindJSON(&request)

	response, err := e.EntityBookingCostAdjustmentService.Create(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (e *EntityBookingCostAdjustmentController) Update(c *gin.Context) {
	var request request.UpdateEntityBookingCostAdjustmentRequest
	c.BindJSON(&request)

	response, err := e.EntityBookingCostAdjustmentService.Update(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (e *EntityBookingCostAdjustmentController) Delete(c *gin.Context) {
	entityID := c.Param("adjustmentId")

	entityIdInt, _ := strconv.Atoi(entityID)

	err := e.EntityBookingCostAdjustmentService.Delete(uint(entityIdInt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
