package controllers

import (
	"booking-api/database"
	"booking-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBoat(context *gin.Context) {
	var boat models.Boat
	boatIdString := context.Param("id")
	boatId, err := strconv.Atoi(boatIdString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boat id"})
		context.Abort()
		return
	}

	boat, err = GetBoatById(boatId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Boat not found"})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, boat)
}

func GetBoats(context *gin.Context) {
	var boats []models.Boat
	database.Instance.Find(&boats)
	context.JSON(http.StatusOK, boats)
}

func CreateBoat(context *gin.Context) {
	var boat models.Boat
	if err := context.ShouldBindJSON(&boat); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&boat)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, boat)
}

func GetBoatById(id int) (models.Boat, error) {
	var boat models.Boat
	record := database.Instance.Where("id = ?", id).First(&boat)
	if record.Error != nil {
		return boat, record.Error
	}
	return boat, nil
}
