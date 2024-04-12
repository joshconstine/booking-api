package middlewares

import (
	"booking-api/models"
	"context"
	"net/http"
	"strings"
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

func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		// store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

		user := models.AuthenticatedUser{
			LoggedIn: true,
			User: models.User{
				Email: "nerd",
			},
		}
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// func WithUserWrapper() gin.HandlerFunc {
// 	return func(context *gin.Context) {

// 		//get http.Handler from gin.HandlerFunc

// 		handler := WithUser()
// 		handler.ServeHTTP(context.Writer, context.Request)
// 	}

// }

// func WithUserWrapper() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Apply the WithUser middleware logic directly here.
// 		if strings.Contains(c.Request.URL.Path, "/public") {
// 			c.Next()
// 			return
// 		}

// 		user := models.AuthenticatedUser{
// 			LoggedIn: true,
// 			User: models.User{
// 				Email: "joshua@gmail.com",
// 			},
// 		}

// 		c.Set(models.UserContextKey, user)
// 		fmt.Println("from the middleware WithUser")
// 		c.Next()
// 	}
// }
