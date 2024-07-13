package middlewares

import (
	"booking-api/config"
	"booking-api/data/response"
	"booking-api/models"
	sb "booking-api/pkg/sb"
	"booking-api/repositories"
	"booking-api/services"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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
				UserID: uuid.MustParse(resp.ID).String(),
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
			userData := userService.FindByUserID(uuid.MustParse(resp.ID).String())
			// if userData.UserID == uuid.Nil {
			// 	// http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
			// 	return
			// }

			user := models.AuthenticatedUser{

				User: response.UserResponse{
					UserID:    uuid.MustParse(resp.ID).String(), //userData.UserID,
					Email:     resp.Email,
					FirstName: userData.FirstName,
					LastName:  userData.LastName,
					Username:  userData.Username,
					Chats:     userData.Chats,

					ProfilePicture: userData.ProfilePicture,
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

func NewWithIsAdminMiddleWare(userService services.UserService) func(http.Handler) http.Handler {
	withIsAdmin := func(next http.Handler) http.Handler {
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
			userData := userService.FindByUserID(uuid.MustParse(resp.ID).String())

			isAdmin := userService.IsAdmin(userData.UserID)
			if !isAdmin {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	return withIsAdmin
}
func NewWithIsOwnerOfEntityMiddleWare(userService services.UserService, membershipRepository repositories.MembershipRepository, entityRepository repositories.EntityRepository) func(http.Handler) http.Handler {
	withIsOwnerOfEntity := func(next http.Handler) http.Handler {
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
			userData := userService.FindByUserID(uuid.MustParse(resp.ID).String())

			//{entityType}/{entityID}/admin
			// Use entityType and entityID from context if they are not provided
			entityType := chi.URLParam(r, "entityType")
			entityIDstr := chi.URLParam(r, "entityID")

			//try form values
			if entityType == "" && entityIDstr == "" {
				entityType = r.FormValue("entityType")
				entityIDstr = r.FormValue("entityID")
			}

			entityID, err := strconv.Atoi(entityIDstr)

			if err != nil {
				// http.Error(w, err.Error(), http.StatusInternalServerError)

			}

			if entityType == "" && entityID == 0 {
				urlParts := strings.Split(r.URL.Path, "/")
				entityType = urlParts[1]
				entityIDstr = urlParts[2]
				entityID, err = strconv.Atoi(entityIDstr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}

			}

			entityIDuint := uint(entityID)

			memberships := membershipRepository.FindAllForUser(userData.UserID)

			if len(memberships) == 0 {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			isOwner, err := entityRepository.IsUserAdminOfEntity(userData.UserID, memberships, entityType, uint(entityIDuint))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return

			}

			if !isOwner {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
	return withIsOwnerOfEntity
}
