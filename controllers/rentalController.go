package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	"html/template"
	"net/http"
	"strconv"

	rentals "booking-api/view/rentals"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
)

type RentalController struct {
	rentalService services.RentalService
}

type RentalListTemplateData struct {
	PageTitle string
	Rentals   []response.RentalResponse
}
type RentalTemplateData struct {
	PageTitle string
	Rental    response.RentalInformationResponse
}

func NewRentalController(rentalService services.RentalService) *RentalController {
	return &RentalController{rentalService: rentalService}
}

func (controller *RentalController) FindAll(ctx *gin.Context) {
	rentals := controller.rentalService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rentals,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *RentalController) FindById(ctx *gin.Context) {
	rentalId := ctx.Param("rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rental,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalController) Create(ctx *gin.Context) {
	var request request.CreateRentalRequest
	ctx.BindJSON(&request)

	rental, err := controller.rentalService.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, rental)
}

func (controller *RentalController) Update(ctx *gin.Context) {
	var request request.UpdateRentalRequest
	ctx.BindJSON(&request)

	rental, err := controller.rentalService.Update(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, rental)
}

func (controller *RentalController) GetRentalTemplate(ctx *gin.Context) {
	tmpl := template.Must(template.ParseFiles("public/singleRental.html"))

	rentalId := ctx.Param("rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	data := RentalTemplateData{
		PageTitle: rental.Name,
		Rental:    rental,
	}

	tmpl.Execute(ctx.Writer, data)

}

func (controller *RentalController) HandleRentalDetail(w http.ResponseWriter, r *http.Request) error {

	// dateParam := chi.URLParam(r, "date")

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	return rentals.RentalInformationResponse(rental).Render(r.Context(), w)
}

func (controller *RentalController) HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// user := view.getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", user.Account)

	rentalData := controller.rentalService.FindAll()

	return rentals.Index(rentalData).Render(r.Context(), w)
}
