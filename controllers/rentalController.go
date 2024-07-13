package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	rentals "booking-api/view/rentals"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"

	"encoding/json"
)

type RentalController struct {
	rentalService  services.RentalService
	amenityService services.AmenityService
}

func NewRentalController(rentalService services.RentalService, amenityService services.AmenityService) *RentalController {
	return &RentalController{rentalService: rentalService, amenityService: amenityService}
}

type RentalListTemplateData struct {
	PageTitle string
	Rentals   []response.RentalResponse
}
type RentalTemplateData struct {
	PageTitle string
	Rental    response.RentalInformationResponse
}

func (controller *RentalController) FindAll(w http.ResponseWriter, r *http.Request) error {
	rentals := controller.rentalService.FindAll()

	rentalsJSON, err := json.Marshal(rentals)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rentalsJSON)
	return nil

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

func (controller *RentalController) Update(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	bedrooms := r.FormValue("bedrooms")
	bathrooms := r.FormValue("bathrooms")

	id, _ := strconv.Atoi(rentalId)
	bedroomsInt, _ := strconv.Atoi(bedrooms)
	bathroomsInt, _ := strconv.Atoi(bathrooms)

	params := rentals.RentalFormParams{
		RentalID:    id,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Bedrooms:    bedroomsInt,
		Bathrooms:   bathroomsInt,
	}

	rental, err := controller.rentalService.UpdateRental(params)

	if err != nil {
		return err

	}

	params = rentals.RentalFormParams{
		Name:        rental.Name,
		Description: rental.Description,
		Bedrooms:    int(rental.Bedrooms),
		Bathrooms:   int(rental.Bathrooms),

		Success: true,
	}

	return rentals.RentalDetails(params).Render(r.Context(), w)
}

func (controller *RentalController) HandleRentalDetail(w http.ResponseWriter, r *http.Request) error {

	// dateParam := chi.URLParam(r, "date")

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	return rentals.RentalInformationResponse(rental).Render(r.Context(), w)
}

func (controller *RentalController) HandleRentalAdminDetail(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	amenities := controller.amenityService.FindAllSorted()
	return rentals.RentalAdmin(rental, amenities).Render(r.Context(), w)
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
