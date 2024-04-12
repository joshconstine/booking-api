package view

import (
	"booking-api/models"
	"context"
	"strconv"
)

func AuthenticatedUser(ctx context.Context) models.AuthenticatedUser {
	user, ok := ctx.Value(models.UserContextKey).(models.AuthenticatedUser)
	if !ok {
		return models.AuthenticatedUser{}
	}
	return user
}

func String(i int) string {
	return strconv.Itoa(i)
}
