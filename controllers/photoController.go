package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoService services.PhotoService
}

func NewPhotoController(photoService services.PhotoService) *PhotoController {
	return &PhotoController{PhotoService: photoService}
}

func (controller *PhotoController) FindAll(ctx *gin.Context) {

	photos := controller.PhotoService.FindAll()

	response := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   photos,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
