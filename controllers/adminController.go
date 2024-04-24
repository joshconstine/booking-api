package controllers

import (
	"booking-api/services"
	admin "booking-api/view/admin"
	"net/http"
)

type AdminController struct {
	userService    services.UserService
	bookingService services.BookingService
	accountService services.AccountService
}

func NewAdminController(service services.UserService, bookingService services.BookingService, accountService services.AccountService) *AdminController {
	return &AdminController{userService: service, bookingService: bookingService, accountService: accountService}
}

func (usc *AdminController) HandleAdminIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	bookings := usc.bookingService.GetSnapshot()

	inquiries, err := usc.accountService.GetInquiriesSnapshot(1)
	if err != nil {
		return err
	}

	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(user, bookings, inquiries).Render(r.Context(), w)
}
