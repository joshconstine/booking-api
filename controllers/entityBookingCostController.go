package controllers

import (
	"booking-api/data/request"
	"booking-api/repositories"
	"booking-api/services"
	ui "booking-api/view/ui"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/gin-gonic/gin"
)

type EntityBookingCostController struct {
	EntityBookingCostService services.EntityBookingCostService
	BookingCostTypeService   services.BookingCostTypeService

	TaxRateRepository repositories.TaxRateRepository
}

func NewEntityBookingCostController(entityBookingCostService services.EntityBookingCostService, bookingCostTypeService services.BookingCostTypeService, taxRateRepository repositories.TaxRateRepository) *EntityBookingCostController {
	return &EntityBookingCostController{EntityBookingCostService: entityBookingCostService, BookingCostTypeService: bookingCostTypeService, TaxRateRepository: taxRateRepository}
}
func (e *EntityBookingCostController) FindAllForEntity(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")

	entityIdInt, _ := strconv.Atoi(entityID)

	response := e.EntityBookingCostService.FindAllForEntity(uint(entityIdInt), entityType)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)

}

// func (e *EntityBookingCostController) Create(w http.ResponseWriter, r *http.Request) error {
//
//	entityID := r.FormValue("entityID")
//	bookingCostTypeID := r.FormValue("costTypeID")
//	amount := r.FormValue("amount")
//	taxRateID := r.FormValue("taxRateID")
//	taxRate := r.FormValue("taxRate")
//	//var errors ui.EntityBookingCostFormErrors
//
//	entityIDInt, _ := strconv.Atoi(entityID)
//	bookingCostTypeIDInt, _ := strconv.Atoi(bookingCostTypeID)
//	amount64, _ := strconv.ParseFloat(amount, 64)
//	taxRateIDInt, _ := strconv.Atoi(taxRateID)
//	taxRate64, _ := strconv.ParseFloat(taxRate, 64)
//
//	params := ui.EntityBookingCostFormParams{
//		EntityType:        r.FormValue("entityType"),
//		EntityID:          uint(entityIDInt),
//		BookingCostTypeID: uint(bookingCostTypeIDInt),
//		Amount:            amount64,
//		TaxRateID:         uint(taxRateIDInt),
//		TaxRatePercentage: taxRate64,
//	}
//	//ok := validate.New(&params, validate.Fields{
//	//	//"Amount":            validate.Rules(minFloat(amount64, 0.01)),
//	//	//"TaxRatePercentage": validate.Rules(minFloat(taxRate64, 0.01)),
//	//}).Validate(&errors)
//	//if !ok {
//	//	return render(r, w, ui.EntityBookingCostForm(params, errors, costTypes, taxRates))
//	//}
//
//	createRequest := request.CreateEntityBookingCostRequest{
//		EntityID:          params.EntityID,
//		EntityType:        params.EntityType,
//		BookingCostTypeID: params.BookingCostTypeID,
//		Amount:            params.Amount,
//		TaxRateID:         params.TaxRateID,
//		TaxRatePercentage: params.TaxRatePercentage,
//	}
//
//	_, err := e.EntityBookingCostService.Create(createRequest)
//	if err != nil {
//		return err
//	}
//
//	return render(r, w, ui.Toast("Entity booking cost created successfully"))
//
// }
func (e *EntityBookingCostController) Create(w http.ResponseWriter, r *http.Request) error {
	// Extract form values and handle errors
	entityID, err := strconv.Atoi(r.FormValue("entityID"))
	if err != nil {
		http.Error(w, "Invalid entityID", http.StatusBadRequest)
		return err
	}

	bookingCostTypeID, err := strconv.Atoi(r.FormValue("bookingCostTypeID"))
	if err != nil {
		http.Error(w, "Invalid costTypeID", http.StatusBadRequest)
		return err
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return err
	}

	taxRateID, err := strconv.Atoi(r.FormValue("taxRateID"))
	if err != nil {
		http.Error(w, "Invalid taxRateID", http.StatusBadRequest)
		return err
	}

	taxRate, err := strconv.ParseFloat(r.FormValue("taxRate"), 64)
	if err != nil {
		http.Error(w, "Invalid taxRate", http.StatusBadRequest)
		return err
	}

	params := ui.EntityBookingCostFormParams{
		EntityType:        r.FormValue("entityType"),
		EntityID:          uint(entityID),
		BookingCostTypeID: uint(bookingCostTypeID),
		Amount:            amount,
		TaxRateID:         uint(taxRateID),
		TaxRatePercentage: taxRate,
	}

	//// Validate the parameters
	//errors := validate.New()
	//validateField(errors, "Amount", amount, minFloat(0.01))
	//validateField(errors, "TaxRatePercentage", taxRate, minFloat(0.01))
	//
	//if errors.HasErrors() {
	//	return render(r, w, ui.EntityBookingCostForm(params, errors, costTypes, taxRates))
	//}

	createRequest := request.CreateEntityBookingCostRequest{
		EntityID:          params.EntityID,
		EntityType:        params.EntityType,
		BookingCostTypeID: params.BookingCostTypeID,
		Amount:            params.Amount,
		TaxRateID:         params.TaxRateID,
		TaxRatePercentage: params.TaxRatePercentage,
	}

	_, err = e.EntityBookingCostService.Create(createRequest)
	if err != nil {
		http.Error(w, "Failed to create entity booking cost", http.StatusInternalServerError)
		return err
	}

	return render(r, w, ui.Toast("Entity booking cost created successfully"))
}

func (e *EntityBookingCostController) Update(c *gin.Context) {
	var request request.UpdateEntityBookingCostRequest
	c.BindJSON(&request)

	response, err := e.EntityBookingCostService.Update(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (e *EntityBookingCostController) Delete(c *gin.Context) {
	entityID := c.Param("entityId")
	entityType := c.Param("entityType")
	bookingCostTypeID := c.Param("bookingCostTypeId")

	entityIdInt, _ := strconv.Atoi(entityID)
	bookingCostTypeIDInt, _ := strconv.Atoi(bookingCostTypeID)

	err := e.EntityBookingCostService.Delete(uint(entityIdInt), entityType, uint(bookingCostTypeIDInt))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Entity booking cost deleted successfully"})
}
func (e *EntityBookingCostController) GetEntityBookingCostForm(w http.ResponseWriter, r *http.Request) error {
	entityType := chi.URLParam(r, "entityType")
	entityID := chi.URLParam(r, "entityID")

	entityIDInt, err := strconv.Atoi(entityID)
	if err != nil {
		// return err
	}
	params := ui.EntityBookingCostFormParams{
		EntityType: entityType,
		EntityID:   uint(entityIDInt),
	}

	errors := ui.EntityBookingCostFormErrors{}
	// render the form

	costTypes := e.BookingCostTypeService.FindAll()
	taxRates := e.TaxRateRepository.FindAll()

	return render(r, w, ui.EntityBookingCostForm(params, errors, costTypes, taxRates))
}
