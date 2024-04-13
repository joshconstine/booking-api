package middlewares

import (
	"booking-api/config"
	"booking-api/models"
	sb "booking-api/pkg/sb"
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

// func WithUser() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		if strings.Contains(context.Request.URL.Path, "/public") {
// 			context.Next()
// 			return
// 		}
// 		user := models.AuthenticatedUser{
// 			LoggedIn: true,
// 			User: models.User{
// 				Email: "joshua@gmail.com",
// 			},
// 		}

// 		context.Set(models.UserContextKey, user)
// 		fmt.Println("from the middleware WithUser")

// 		// fmt.Printf("User: %+v\n", user)
// 		context.Next()
// 	}
// }

// func WithAccountSetup(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		user := getAuthenticatedUser(r)
// 		account, err := db.GetAccountByUserID(user.ID)
// 		// The user has not setup his account yet.
// 		// Hence, redirect him to /account/setup
// 		if err != nil {
// 			if errors.Is(err, sql.ErrNoRows) {
// 				http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
// 				return
// 			}
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		fmt.Println("in here!!!!!")
// 		user.Account = account
// 		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// 	return http.HandlerFunc(fn)
// }

// func WithUser(next http.Handler) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		if strings.Contains(r.URL.Path, "/public") {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		// store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

// 		// user := models.AuthenticatedUser{
// 		// 	LoggedIn: true,
// 		// 	User: models.User{
// 		// 		Email: "nerd",
// 		// 	},
// 		// }
// 		ctx := context.WithValue(r.Context(), models.UserContextKey, nil)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// 	return http.HandlerFunc(fn)
// }

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
			User: models.User{
				UserID: uuid.MustParse(resp.ID),
				Email:  resp.Email,
			},
			LoggedIn:    true,
			AccessToken: accessToken.(string),
		}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
