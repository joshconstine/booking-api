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

	bookingResponse, err := t.bookingService.FindById(id)
	if err != nil {
		webResponse := response.Response{
			Code:   404,
			Status: "Not Found",
			Data:   nil,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		return

	}
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
	//return render(r, w, bookings.BookingConfirmation(bookingId))

	//reroute to /bookings/{bookingId}
	http.Redirect(w, r, fmt.Sprintf("/bookings/%s", bookingId), http.StatusFound)
	return nil
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

	// Default values for pagination
	limit := 10 // Default limit
	page := 1   // Default page
	sort := ""  // Default sort

	// Parse query parameters
	query := r.URL.Query()

	// Get 'limit' from query parameters, if present
	if limitParam := query.Get("limit"); limitParam != "" {
		limit, _ = strconv.Atoi(limitParam) // Convert to int
	}

	// Get 'page' from query parameters, if present
	if pageParam := query.Get("page"); pageParam != "" {
		page, _ = strconv.Atoi(pageParam) // Convert to int
	}

	// Get 'sort' from query parameters, if present
	sort = query.Get("sort")

	request := request.GetBookingSnapshotRequest{
		SearchString: "",
		Statuses:     []int{},
		PaginationRequest: request.PaginationRequest{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
	}

	bookingData := controller.bookingService.GetSnapshot(request)

	return bookings.Index(bookingData, request.PaginationRequest).Render(r.Context(), w)
}
func (controller *BookingController) HandleCreateBookingPage(w http.ResponseWriter, r *http.Request) error {
	return bookings.CreateBookingPage().Render(r.Context(), w)
}

func (controller *BookingController) HandleBookingInformation(w http.ResponseWriter, r *http.Request) error {

	bookingId := chi.URLParam(r, "bookingId")

	booking, err := controller.bookingService.FindById(bookingId)

	if err != nil {
		return err
	}

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

func (controller *BookingController) HandleAuditBookingStautsTrigger(w http.ResponseWriter, r *http.Request) error {
	err := controller.bookingService.AuditAllBookingStatus()
	if err != nil {
		return err
	}
	return nil
}
