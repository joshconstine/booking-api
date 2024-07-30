package repositories

import (
	"booking-api/data/response"
)

type MembershipRepository interface {
	FindAllForUser(userID string) []response.MembershipResponse
}
