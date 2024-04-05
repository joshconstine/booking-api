package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"time"

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

func (controller *PhotoController) AddPhoto(ctx *gin.Context, entity string, entityID int) {

	var response response.Response

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	r := ctx.Request
	w := ctx.Writer
	timeout := 20 * time.Second

	var cancelFn func()
	if timeout > 0 {
		// ctx, cancelFn = ctx.WithTimeout(ctx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}
	err := r.ParseMultipartForm(10 * 1024 * 1024) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)

	}

	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusInternalServerError)
		// return "", err
		response.Code = http.StatusBadRequest
		response.Status = http.StatusText(http.StatusBadRequest)
		response.Data = "Failed to get file from form"

	}
	defer file.Close()

	photoResult := controller.PhotoService.AddPhoto(&file, header, entity, entityID)

	response.Code = http.StatusOK
	response.Status = http.StatusText(http.StatusOK)
	response.Data = photoResult

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)

}
