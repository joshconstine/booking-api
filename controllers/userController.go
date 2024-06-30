package controllers

import (
	"booking-api/data/request"
	validate "booking-api/pkg/kit"
	"booking-api/services"
	bookings "booking-api/view/bookings"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{userService: service}
}

//func (controller *UserController) RegisterUser(context *gin.Context) {
//	var user models.User
//	if err := context.ShouldBindJSON(&user); err != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		context.Abort()
//		return
//	}
//	if err := user.HashPassword(user.Login.Password); err != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		context.Abort()
//		return
//	}
//	record := database.Instance.Create(&user)
//	if record.Error != nil {
//		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
//		context.Abort()
//		return
//	}
//	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
//}

//func (controller *UserController) FindAll(context *gin.Context) {
//	users := controller.userService.FindAll()
//
//	webResponse := response.Response{
//		Code:   200,
//		Status: "Ok",
//		Data:   users,
//	}
//
//	context.JSON(http.StatusOK, webResponse)
//}

func (controller *UserController) FindOrCreateUser(w http.ResponseWriter, r *http.Request) error {
	params := request.CreateBookingWithUserInformationRequest{
		FirstName: r.FormValue("firstName"),
		LastName:  r.FormValue("lastName"),
		Email:     r.FormValue("email"),
	}

	errors := bookings.BookingUserInformationErrors{}
	ok := validate.New(&params, validate.Fields{
		"Email": validate.Rules(validate.Required),
	}).Validate(&errors)
	if !ok {
		return render(r, w, bookings.BookingUserInformationForm(params, errors))
	}
	//reroute to /bookings/{bookingId}
	//http.Redirect(w, r, fmt.Sprintf("/bookings/%s", bookingId), http.StatusFound)
	//return nil

	user, err := controller.userService.FindByEmailPublic(params.Email)

	if err != nil {
		return err

	}
	if user.Email == "" {
		return render(r, w, bookings.BookingUserInformationForm(params, errors))
	}

	return render(r, w, bookings.GuestConfirmationDialog(params, user))

}
