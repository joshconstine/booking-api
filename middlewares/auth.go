package middlewares

import (
	"booking-api/controllers"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/services"
	"context"
	"net/http"
	"strings"
)

func NewWithAccountSetupMiddleWare(userService services.UserService) func(http.Handler) http.Handler {
	withAccountSetup := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authenticatedUser := controllers.GetAuthenticatedUser(r)
			userData := controllers.GetAuthenticatedUser(r)
			if !userData.LoggedIn {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			user, err := userService.FindByUserID(authenticatedUser.User.UserID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			if user.UserID == "" {
				http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
				return
			}
			authenticatedUser.User = response.UserResponse{
				UserID:         authenticatedUser.User.UserID,
				Username:       user.Username,
				Email:          user.Email,
				FirstName:      user.FirstName,
				LastName:       user.LastName,
				PhoneNumber:    user.PhoneNumber,
				Chats:          user.Chats,
				Bookings:       user.Bookings,
				ProfilePicture: user.ProfilePicture,
			}

			ctx := context.WithValue(r.Context(), models.UserContextKey, authenticatedUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	return withAccountSetup
}

func WithAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := controllers.GetAuthenticatedUser(r)
		if !user.LoggedIn {
			path := r.URL.Path
			http.Redirect(w, r, "/login?to="+path, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
