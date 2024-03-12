package controllers

import (
	"booking-api/database"
	"booking-api/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBoat is a function to get a boat by id

func DeleteBoatPhoto(context *gin.Context) {
	var boatPhoto models.BoatPhoto
	boatPhotoID := context.Param("id")
	record := database.Instance.Where("id = ?", boatPhotoID).Delete(&boatPhoto)
	if record.Error != nil {
		context.JSON(404, gin.H{"error": "BoatPhoto not found"})
		context.Abort()
		return
	}
	context.JSON(200, gin.H{"success": "BoatPhoto deleted"})
}

func CreateBoatPhoto(context *gin.Context) {
	var boatPhoto models.BoatPhoto
	var boatID int
	boatID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid boat id"})
		context.Abort()
		return
	}
	newFilePath := "boat_photos/" + context.Param("id")

	w := context.Writer
	r := context.Request

	//log the request
	log.Println(r)

	uploadedFilePath, err := UploadHandler(w, r, newFilePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error uploading photo", http.StatusInternalServerError)
		return
	}

	boatPhoto.PhotoURL = uploadedFilePath
	boatPhoto.BoatID = boatID
	boat, err := GetBoatById(boatID)
	if err != nil {
		context.JSON(404, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	boatPhoto.Boat = boat

	record := database.Instance.Create(&boatPhoto)
	if record.Error != nil {
		context.JSON(500, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	log.Println(uploadedFilePath)

	context.JSON(201, boatPhoto)
}

func GetBoatPhotosForBoat(context *gin.Context) {
	var boatPhotos []models.BoatPhoto
	boatID := context.Param("id")
	database.Instance.Where("boat_id = ?", boatID).Find(&boatPhotos)
	context.JSON(200, boatPhotos)
}
