package view

import (
	"booking-api/models"
	"context"
	"fmt"
	"strconv"
)

func AuthenticatedUser(ctx context.Context) models.AuthenticatedUser {
	fmt.Println("in authenticatedUser")
	user, _ := ctx.Value(models.UserContextKey).(models.AuthenticatedUser)
	// if !ok {
	// 	// return models.AuthenticatedUser{}

	// }

	fmt.Println("from the AuthenticatedUser")

	fmt.Printf("User: %+v\n", user)

	return user
}

func String(i int) string {
	return strconv.Itoa(i)
}
