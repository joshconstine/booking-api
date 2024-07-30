package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentMethodController struct {
	PaymentMethodService services.PaymentMethodService
}

func NewPaymentMethodController(service services.PaymentMethodService) *PaymentMethodController {
	return &PaymentMethodController{PaymentMethodService: service}
}

func (controller *PaymentMethodController) FindById(ctx *gin.Context) {

	paymentMethodId := ctx.Param("paymentMethodId")
	id, _ := strconv.Atoi(paymentMethodId)

	paymentMethod := controller.PaymentMethodService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   paymentMethod,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PaymentMethodController) FindAll(ctx *gin.Context) {
	paymentMethods := controller.PaymentMethodService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   paymentMethods,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
