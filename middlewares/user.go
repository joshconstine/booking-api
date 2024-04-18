package middlewares

import (
	"booking-api/config"
	"booking-api/data/response"
	"booking-api/models"
	sb "booking-api/pkg/sb"
	"booking-api/services"
	"context"
	"net/http"

	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		// load config
		env, _ := config.LoadConfig(".")
		store := sessions.NewCookieStore([]byte(env.SESSION_SECRET))
		session, err := store.Get(r, sessionUserKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		accessToken := session.Values[sessionAccessTokenKey]
		if accessToken == nil {
			next.ServeHTTP(w, r)
			return
		}
		resp, err := sb.ClientInstance.Auth.User(r.Context(), accessToken.(string))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user := models.AuthenticatedUser{
			User: response.UserResponse{
				UserID: uuid.MustParse(resp.ID),
				Email:  resp.Email,

				//TODO ADD user service to get username here
			},
			LoggedIn:    true,
			AccessToken: accessToken.(string),
		}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func NewWithUserMiddleWare(userService services.UserService) func(http.Handler) http.Handler {
	withUser := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)

				return
			}
			// load config
			env, _ := config.LoadConfig(".")
			store := sessions.NewCookieStore([]byte(env.SESSION_SECRET))
			session, err := store.Get(r, sessionUserKey)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			accessToken := session.Values[sessionAccessTokenKey]
			if accessToken == nil {
				next.ServeHTTP(w, r)
				return
			}
			resp, err := sb.ClientInstance.Auth.User(r.Context(), accessToken.(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			userData := userService.FindByUserID(uuid.MustParse(resp.ID))
			// if userData.UserID == uuid.Nil {
			// 	// http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
			// 	return
			// }

			user := models.AuthenticatedUser{

				User: response.UserResponse{
					UserID:    uuid.MustParse(resp.ID), //userData.UserID,
					Email:     resp.Email,
					FirstName: userData.FirstName,
					LastName:  userData.LastName,

					//TODO ADD user service to get username here
				},
				LoggedIn:    true,
				AccessToken: accessToken.(string),
			}
			ctx := context.WithValue(r.Context(), models.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	return withUser
}
