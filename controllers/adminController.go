package controllers

import (
	"booking-api/services"
	admin "booking-api/view/admin"
	"net/http"
)

type AdminController struct {
	userService services.UserService
}

func NewAdminController(service services.UserService) *AdminController {
	return &AdminController{userService: service}
}

func (usc *AdminController) HandleAdminIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(user).Render(r.Context(), w)
}
