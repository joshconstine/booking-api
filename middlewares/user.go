package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		if strings.Contains(context.Request.URL.Path, "/public") {
			context.Next()
			return
		}
		fmt.Println("from the user middleware")
		context.Next()
	}
}

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
// 		store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
// 		session, err := store.Get(r, sessionUserKey)
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		accessToken := session.Values[sessionAccessTokenKey]
// 		if accessToken == nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		resp, err := sb.Client.Auth.User(r.Context(), accessToken.(string))
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		user := types.AuthenticatedUser{
// 			ID:          uuid.MustParse(resp.ID),
// 			Email:       resp.Email,
// 			LoggedIn:    true,
// 			AccessToken: accessToken.(string),
// 		}
// 		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// 	return http.HandlerFunc(fn)
// }
