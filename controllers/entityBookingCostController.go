package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityBookingCostController struct {
	EntityBookingCostService services.EntityBookingCostService
}

func NewEntityBookingCostController(entityBookingCostService services.EntityBookingCostService) *EntityBookingCostController {
	return &EntityBookingCostController{EntityBookingCostService: entityBookingCostService}
}

func (e *EntityBookingCostController) FindAllForEntity(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")

	entityIdInt, _ := strconv.Atoi(entityID)

	response := e.EntityBookingCostService.FindAllForEntity(uint(entityIdInt), entityType)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}

func (e *EntityBookingCostController) Create(c *gin.Context) {
	var request request.CreateEntityBookingCostRequest
	c.BindJSON(&request)

	response, err := e.EntityBookingCostService.Create(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (e *EntityBookingCostController) Update(c *gin.Context) {
	var request request.UpdateEntityBookingCostRequest
	c.BindJSON(&request)

	response, err := e.EntityBookingCostService.Update(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (e *EntityBookingCostController) Delete(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")
	bookingCostTypeID := c.Param("bookingCostTypeId")

	entityIdInt, _ := strconv.Atoi(entityID)
	bookingCostTypeIDInt, _ := strconv.Atoi(bookingCostTypeID)

	err := e.EntityBookingCostService.Delete(uint(entityIdInt), entityType, uint(bookingCostTypeIDInt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Entity booking cost deleted successfully"})
}
