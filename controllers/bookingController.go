package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	bookings "booking-api/view/bookings"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
)

type BookingController struct {
	bookingService        services.BookingService
	bookingDetailsService services.BookingDetailsService
	invoiceService        services.InvoiceService
}

func NewBookingController(service services.BookingService, detailsService services.BookingDetailsService, invoiceService services.InvoiceService) *BookingController {
	return &BookingController{bookingService: service, bookingDetailsService: detailsService, invoiceService: invoiceService}
}

func (t BookingController) FindAll(ctx *gin.Context) {
	bookingResponse := t.bookingService.FindAll()
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
func convertStringToUint(s string) uint {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return uint(id)
}
func (t BookingController) FindById(ctx *gin.Context) {
	id := ctx.Param("bookingId")

	bookingResponse := t.bookingService.FindById(id)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (t BookingController) CreateBookingWithUserInformation(ctx *gin.Context) {
	var request request.CreateUserRequest
	ctx.BindJSON(&request)

	// bookingResponse, err := t.bookingService.Create(request)

	// bookingResponse, err :=
	// if err != nil {
	// 	webResponse := response.Response{
	// 		Code:   http.StatusBadRequest,
	// 		Status: http.StatusText(http.StatusBadRequest),
	// 		Data:   err.Error(),
	// 	}

	// 	ctx.Header("Content-Type", "application/json")
	// 	ctx.JSON(http.StatusBadRequest, webResponse)
	// 	return
	// }

	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   "not implemented",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (t BookingController) GetDetailsForBookingID(ctx *gin.Context) {
	bookingId := ctx.Param("bookingId")
	id := convertStringToUint(bookingId)

	bookingDetailsResponse := t.bookingDetailsService.FindById(id)
	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   bookingDetailsResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *BookingController) HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// user := view.getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", user.Account)

	bookingData := controller.bookingService.FindAll()

	return bookings.Index(bookingData).Render(r.Context(), w)
}
func (controller *BookingController) HandleCreateBookingPage(w http.ResponseWriter, r *http.Request) error {
	return bookings.CreateBookingPage().Render(r.Context(), w)
}

func (controller *BookingController) HandleBookingInformation(w http.ResponseWriter, r *http.Request) error {

	bookingId := chi.URLParam(r, "bookingId")

	booking := controller.bookingService.FindById(bookingId)

	return bookings.BookingInformationTemplate(booking).Render(r.Context(), w)
}

func (controller *BookingController) HandleCreateInvoiceForBooking(w http.ResponseWriter, r *http.Request) error {

	// vars := mux.Vars(r)
	// bookingId := vars["id"]
	bookingId := chi.URLParam(r, "bookingId")

	invoiceId, err := controller.invoiceService.CreateInvoiceForBooking(bookingId)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create invoice: %v", err), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(invoiceId))
	return nil

}
