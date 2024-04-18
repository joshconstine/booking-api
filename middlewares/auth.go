package middlewares

import (
	"booking-api/auth"
	"booking-api/controllers"
	"booking-api/models"
	"booking-api/services"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
func NewWithAccountSetupMiddleWare(userService services.UserService) func(http.Handler) http.Handler {
	withAccountSetup := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authenticatedUser := controllers.GetAuthenticatedUser(r)
			user := userService.FindByUserID(authenticatedUser.User.UserID)
			if user.UserID == uuid.Nil {
				http.Redirect(w, r, "/account/setup", http.StatusSeeOther)
				return
			}
			fmt.Println("in here!!!!!")
			authenticatedUser.User = user
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
