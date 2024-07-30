package repositories

import "booking-api/data/response"

type EntityRepository interface {
	IsUserAdminOfEntity(userID string, memberships []response.MembershipResponse, entityType string, entityID uint) (bool, error)
}
