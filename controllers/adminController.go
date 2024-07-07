package controllers

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	admin "booking-api/view/admin"
	"net/http"
	"strconv"
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

	// Default values for pagination
	limit := 10 // Default limit
	page := 1   // Default page
	sort := ""  // Default sort

	// Parse query parameters
	query := r.URL.Query()

	// Get 'limit' from query parameters, if present
	if limitParam := query.Get("limit"); limitParam != "" {
		limit, _ = strconv.Atoi(limitParam) // Convert to int
	}

	// Get 'page' from query parameters, if present
	if pageParam := query.Get("page"); pageParam != "" {
		page, _ = strconv.Atoi(pageParam) // Convert to int
	}

	// Get 'sort' from query parameters, if present
	sort = query.Get("sort")

	request := request.GetBookingSnapshotRequest{
		SearchString: "",
		Statuses:     []int{},
		PaginationRequest: request.PaginationRequest{
			Limit: limit,
			Page:  page,
			Sort:  sort,
		},
	}
	bookings := usc.bookingService.GetSnapshot(request)

	userAccountRoles, err := usc.accountService.GetUserAccountRoles(user.User.UserID)
	if err != nil {
		return err
	}

	uniqueAccountIDs := []uint{}

	for _, role := range userAccountRoles {
		unique := true
		for _, id := range uniqueAccountIDs {
			if id == role.AccountID {
				unique = false
				break
			}
		}
		if unique {
			uniqueAccountIDs = append(uniqueAccountIDs, role.AccountID)
		}
	}

	for _, accountID := range uniqueAccountIDs {
		accinquiries, err := usc.accountService.GetInquiriesSnapshot(accountID)
		if err != nil {
			return err
		}
		accmessages, err := usc.accountService.GetMessagesSnapshot(accountID)
		if err != nil {
			return err
		}

		inquiries.Inquiries = append(inquiries.Inquiries, accinquiries.Inquiries...)
		inquiries.Notifications += accinquiries.Notifications

		messages.Chats = append(messages.Chats, accmessages.Chats...)
		messages.Notifications += accmessages.Notifications

	}

	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return admin.Index(bookings, inquiries, messages).Render(r.Context(), w)
}
