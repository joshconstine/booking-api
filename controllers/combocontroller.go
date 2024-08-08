package controllers

import (
	"booking-api/services"
	rentals "booking-api/view/rentals"
	"net/http"
)

type ComboController struct {
	boatService   services.BoatService
	rentalService services.RentalService
}

func NewComboController(service services.BoatService, rentalService services.RentalService) *ComboController {
	return &ComboController{boatService: service, rentalService: rentalService}

}

// func (controller *TagController) Create(ctx *gin.Context) {
// 	createTagRequest := request.CreateTagsRequest{}
// 	err := ctx.ShouldBindJSON(&createTagRequest)
// 	helper.ErrorPanic(err)

// 	controller.tagService.Create(createTagRequest)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)
// }

// func (controller *TagController) Update(ctx *gin.Context) {
// 	updateTagRequest := request.UpdateTagsRequest{}
// 	err := ctx.ShouldBindJSON(&updateTagRequest)
// 	helper.ErrorPanic(err)

// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)

// 	updateTagRequest.Id = id

// 	controller.tagService.Update(updateTagRequest)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)
// }

// func (controller *TagController) Delete(ctx *gin.Context) {
// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)
// 	controller.tagService.Delete(id)

// 	webResponse := response.Response{
// 		Code:   200,
// 		Status: "Ok",
// 		Data:   nil,
// 	}

// 	ctx.JSON(http.StatusOK, webResponse)

// }

func (controller *ComboController) HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// user := view.getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", user.Account)

	boatData := controller.boatService.FindAll()
	rentalData := controller.rentalService.FindAll()

	return rentals.Index(boatData, rentalData).Render(r.Context(), w)
}
func (controller *ComboController) HandleAllCards(w http.ResponseWriter, r *http.Request) error {

	boatData := controller.boatService.FindAll()
	rentalData := controller.rentalService.FindAll()

	return rentals.RentalAndBoatCards(boatData, rentalData).Render(r.Context(), w)
}
func (controller *ComboController) HandlePropertyRentalCards(w http.ResponseWriter, r *http.Request) error {

	rentalData := controller.rentalService.FindAll()

	return rentals.PropertyRentalCards(rentalData).Render(r.Context(), w)
}
func (controller *ComboController) HandleBoatRentalCards(w http.ResponseWriter, r *http.Request) error {

	boatData := controller.boatService.FindAll()

	return rentals.BoatRentalCards(boatData).Render(r.Context(), w)
}
