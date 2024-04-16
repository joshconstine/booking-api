package middlewares

import (
	"booking-api/auth"
	"booking-api/controllers"
	"booking-api/models"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

func WithAccountSetup(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := controllers.GetAuthenticatedUser(r)
		// account, err := db.GetAccountByUserID(user.ID)
		// The user has not setup his account yet.
		// Hence, redirect him to /account/setup
		// if err != nil {
		// 	if errors.Is(err, sql.ErrNoRows) {
		// 		http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
		// 		return
		// 	}
		// 	next.ServeHTTP(w, r)
		// 	return
		// }
		fmt.Println("in here!!!!!")
		// user.Account = account
		ctx := context.WithValue(r.Context(), models.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
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
