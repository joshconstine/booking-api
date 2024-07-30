package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntityBookingDocumentController struct {
	EntityBookingDocumentService services.EntityBookingDocumentService
}

func NewEntityBookingDocumentController(entityBookingDocumentService services.EntityBookingDocumentService) *EntityBookingDocumentController {
	return &EntityBookingDocumentController{EntityBookingDocumentService: entityBookingDocumentService}
}

func (e *EntityBookingDocumentController) FindEntityBookingDocumentsForEntity(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")

	entityIDInt, err := strconv.Atoi(entityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entityID"})
		return
	}

	entityBookingDocuments := e.EntityBookingDocumentService.FindEntityBookingDocumentsForEntity(uint(entityIDInt), entityType)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"data": entityBookingDocuments})
}

func (e *EntityBookingDocumentController) CreateEntityBookingDocument(c *gin.Context) {
	var request request.CreateEntityBookingDocumentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entityBookingDocument, err := e.EntityBookingDocumentService.CreateEntityBookingDocument(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"data": entityBookingDocument})
}
