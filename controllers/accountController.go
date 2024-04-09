package controllers

import (
	"booking-api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountController(accountRepository repositories.AccountRepository) *AccountController {
	return &AccountController{
		AccountRepository: accountRepository,
	}
}

func (controller *AccountController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	account, err := controller.AccountRepository.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
