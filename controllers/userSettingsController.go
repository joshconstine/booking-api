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
	return render(r, w, settings.Index(user, "profile"))
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

func (usc *UserSettingsController) HandleSubscriptions(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Subscription(user))
	// return nil
}

func (usc *UserSettingsController) HandleTeam(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Team(user))
}

func (usc *UserSettingsController) HandleFinances(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Finances(user))
}

func (usc *UserSettingsController) HandleNotifications(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Notifications(user))
}

func (usc *UserSettingsController) HandleCleaners(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Cleaners(user))
}

func (usc *UserSettingsController) HandleSecurity(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Security(user))
}

func (usc *UserSettingsController) HandleProfile(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Profile(user))
}
