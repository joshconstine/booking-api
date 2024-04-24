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

	accountId := uint(1)
	//TODO use account middle ware to protect this and get the id
	inquiries, err := usc.accountService.GetInquiriesSnapshot(accountId)
	if err != nil {
		return err
	}
	messages, err := usc.accountService.GetMessagesSnapshot(accountId)
	if err != nil {
		return err
	}

	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(user, bookings, inquiries, messages).Render(r.Context(), w)
}
