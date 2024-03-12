package controllers

import (
	"booking-api/services"
	"net/http"

	"booking-api/data/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoatController struct {
	boatService services.BoatService
}

func NewBoatController(service services.BoatService) *BoatController {
	return &BoatController{boatService: service}
}

// func (controller *TagController) Create(ctx *gin.Context) {
// 	createTagRequest := request.CreateTagsRequest{}
// 	err := ctx.ShouldBindJSON(&createTagRequest)
// 	helper.ErrorPanic(err)

// 	controller.tagService.Create(createTagRequest)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)
// }

// func (controller *TagController) Update(ctx *gin.Context) {
// 	updateTagRequest := request.UpdateTagsRequest{}
// 	err := ctx.ShouldBindJSON(&updateTagRequest)
// 	helper.ErrorPanic(err)

// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)

// 	updateTagRequest.Id = id

// 	controller.tagService.Update(updateTagRequest)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)
// }

// func (controller *TagController) Delete(ctx *gin.Context) {
// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)
// 	controller.tagService.Delete(id)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)

// }

func (controller *BoatController) FindById(ctx *gin.Context) {
	boatId := ctx.Param("boatId")
	id, err := strconv.Atoi(boatId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boat id"})
		return
	}
	tagResponse := controller.boatService.FindByID(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BoatController) FindAll(ctx *gin.Context) {
	boatResponse := controller.boatService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   boatResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
