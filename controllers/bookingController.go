package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	validate "booking-api/pkg/kit"
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

func (t BookingController) CreateBookingWithUserInformation(w http.ResponseWriter, r *http.Request) error {
	params := request.CreateBookingWithUserInformationRequest{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
	}

	errors := bookings.BookingUserInformationErrors{}
	ok := validate.New(&params, validate.Fields{
		"FirstName": validate.Rules(validate.Required),
		"LastName":  validate.Rules(validate.Required),
		"Email":     validate.Rules(validate.Required),
	}).Validate(&errors)
	if !ok {
		return render(r, w, bookings.BookingUserInformationForm(params, errors))
	}

	bookingId, err := t.bookingService.CreateBookingWithUserInformation(&params)
	if err != nil {
		return err

	}
	return render(r, w, bookings.BookingConfirmation(bookingId))

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

	bookingData := controller.bookingService.GetSnapshot()

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
