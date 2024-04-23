package controllers

import (
	"booking-api/services"
	admin "booking-api/view/admin"
	"net/http"
)

type AdminController struct {
	userService    services.UserService
	bookingService services.BookingService
}

func NewAdminController(service services.UserService, bookingService services.BookingService) *AdminController {
	return &AdminController{userService: service, bookingService: bookingService}
}

func (usc *AdminController) HandleAdminIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	bookings := usc.bookingService.GetSnapshot()
	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(user, bookings).Render(r.Context(), w)
}
