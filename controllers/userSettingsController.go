package controllers

import (
	"booking-api/data/request"
	validate "booking-api/pkg/kit"
	"booking-api/services"
	"booking-api/view/settings"
	"net/http"
)

type UserSettingsController struct {
	userService services.UserService
}

func NewUserSettingsController(service services.UserService) *UserSettingsController {
	return &UserSettingsController{userService: service}
}

func (usc *UserSettingsController) HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return render(r, w, settings.Index(user))
}

func (usc *UserSettingsController) HandleSettingsUpdate(w http.ResponseWriter, r *http.Request) error {
	params := settings.ProfileParams{
		Username:    r.FormValue("username"),
		FirstName:   r.FormValue("firstName"),
		LastName:    r.FormValue("lastName"),
		PhoneNumber: r.FormValue("phoneNumber"),
	}
	errors := settings.ProfileErrors{}
	ok := validate.New(&params, validate.Fields{
		"Username": validate.Rules(validate.Min(3), validate.Max(40)),
	}).Validate(&errors)
	if !ok {
		return render(r, w, settings.ProfileForm(params, errors))
	}
	user := GetAuthenticatedUser(r)
	user.User.Username = params.Username

	updateUserRequest := request.UpdateUserRequest{
		UserID:      user.User.UserID,
		Username:    user.User.Username,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		PhoneNumber: params.PhoneNumber,
	}

	// if err := db.UpdateY(&user.User); err != nil {
	// 	return err
	// }

	err := usc.userService.UpdateUser(&updateUserRequest)
	if err != nil {
		return err
	}

	params.Success = true
	return render(r, w, settings.ProfileForm(params, settings.ProfileErrors{}))
}
