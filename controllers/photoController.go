package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"booking-api/view/ui"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoService       services.PhotoService
	EntityPhotoService services.EntityPhotoService
}

func NewPhotoController(photoService services.PhotoService, entityPhotoService services.EntityPhotoService) *PhotoController {
	return &PhotoController{
		PhotoService:       photoService,
		EntityPhotoService: entityPhotoService,
	}
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
func (controller *PhotoController) FindAllForEntity(ctx *gin.Context, entity string, entityID uint) {

	photos := controller.EntityPhotoService.FindAllForEntity(entity, entityID)

	response := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   photos,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}

func (controller *PhotoController) AddPhoto(w http.ResponseWriter, r *http.Request, entity string, entityID int) {

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.

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
		w.Write([]byte("Failed to get file from form"))

	}
	defer file.Close()

	photoResult := controller.PhotoService.AddPhoto(&file, header, entity, entityID)
	entityPhotoResult := controller.EntityPhotoService.AddPhotoToEntity(photoResult.ID, entity, uint(entityID))

	if entityPhotoResult.ID == 0 {
		w.Write([]byte("Failed to add photo to entity"))
		return
	}

	// ctx.Header("Content-Type", "application/json")
	// ctx.JSON(http.StatusOK, response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Photo added successfully"))

}

func (controller *PhotoController) AddPhotoForm(w http.ResponseWriter, r *http.Request) error {

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.

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
		//w.Write([]byte("Failed to get file from form"))

	}
	defer file.Close()

	entity := r.FormValue("entityType")
	entityID := r.FormValue("entityID")

	entityIDInt, _ := strconv.Atoi(entityID)

	params := ui.AddPhotoToEntityFormParams{
		EntityID:   uint(entityIDInt),
		EntityType: entity,
		Errors:     []string{},
	}
	photoResult := controller.PhotoService.AddPhoto(&file, header, entity, entityIDInt)
	entityPhotoResult := controller.EntityPhotoService.AddPhotoToEntity(photoResult.ID, entity, uint(entityIDInt))

	if entityPhotoResult.ID == 0 {
		w.Write([]byte("Failed to add photo to entity"))
		params.Errors = append(params.Errors, "Failed to add photo to entity")

		return render(r, w, ui.AddPhotoToEntityForm(params))
	}

	// ctx.Header("Content-Type", "application/json")
	// ctx.JSON(http.StatusOK, response)

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Photo added successfully"))

	params.Success = true
	return render(r, w, ui.AddPhotoToEntityForm(params))
}
