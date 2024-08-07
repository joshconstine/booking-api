package controllers

import (
	"net/http"
	"strconv"

	"booking-api/services"

	responses "booking-api/data/response"

	requests "booking-api/data/request"

	"github.com/gin-gonic/gin"
)

type BookingPaymentController struct {
	bookingPaymentService services.BookingPaymentService
}

func NewBookingPaymentController(service services.BookingPaymentService) *BookingPaymentController {
	return &BookingPaymentController{bookingPaymentService: service}
}

func (controller *BookingPaymentController) FindById(ctx *gin.Context) {
	bookingPaymentId := ctx.Param("bookingPaymentId")
	id, _ := strconv.Atoi(bookingPaymentId)

	bookingPayment := controller.bookingPaymentService.FindById(uint(id))

	webResponse := responses.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingPayment,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingPaymentController) FindAll(ctx *gin.Context) {
	bookingPayments := controller.bookingPaymentService.FindAll()

	webResponse := responses.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingPayments,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingPaymentController) Create(ctx *gin.Context) {
	createBookingPaymentRequest := requests.CreateBookingPaymentRequest{}
	err := ctx.ShouldBindJSON(&createBookingPaymentRequest)

	if err != nil {
		panic(err)
	}

	result, err := controller.bookingPaymentService.Create(createBookingPaymentRequest)

	if err != nil {
		webResponse := responses.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return

	}

	webResponse := responses.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingPaymentController) FindByBookingId(ctx *gin.Context) {
	bookingId := ctx.Param("bookingId")

	bookingPayments := controller.bookingPaymentService.FindByBookingId(bookingId)

	webResponse := responses.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingPayments,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *BookingPaymentController) FindTotalAmountByBookingId(ctx *gin.Context) {
	bookingId := ctx.Param("bookingId")

	totalAmount := controller.bookingPaymentService.FindTotalPaidByBookingId(bookingId)

	webResponse := responses.Response{
		Code:   200,
		Status: "Ok",
		Data:   totalAmount,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
