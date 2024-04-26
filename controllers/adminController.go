package controllers

import (
	"booking-api/data/response"
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
	var inquiries response.AccountInquiriesSnapshot
	var messages response.AccountMessagesSnapshot

	bookings := usc.bookingService.GetSnapshot()

	userAccountRoles, err := usc.accountService.GetUserAccountRoles(user.User.UserID)
	if err != nil {
		return err
	}

	for _, role := range userAccountRoles {
		accinquiries, err := usc.accountService.GetInquiriesSnapshot(role.AccountID)
		if err != nil {
			return err
		}

		inquiries.Inquiries = append(inquiries.Inquiries, accinquiries.Inquiries...)
		accinquiries.Notifications = inquiries.Notifications + accinquiries.Notifications

		accmessages, err := usc.accountService.GetMessagesSnapshot(role.AccountID)
		if err != nil {
			return err
		}

		messages.Chats = append(messages.Chats, accmessages.Chats...)
		accmessages.Notifications = messages.Notifications + accmessages.Notifications

	}

	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(user, bookings, inquiries, messages).Render(r.Context(), w)
}
