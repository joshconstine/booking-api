package controllers

import (
	"booking-api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InquiryController struct {
	InquiryRepository repositories.InquiryRepository
}

func NewInquiryController(accountRepository repositories.InquiryRepository) *InquiryController {
	return &InquiryController{
		InquiryRepository: accountRepository,
	}
}

func (controller *InquiryController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	account, err := controller.InquiryRepository.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Inquiry not found"})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
