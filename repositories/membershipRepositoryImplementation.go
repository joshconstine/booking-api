package repositories

import (
	"booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
)

type MembershipRepositoryImplementation struct {
	Db *gorm.DB
}

func NewMembershipRepositoryImplementation(db *gorm.DB) MembershipRepository {
	return &MembershipRepositoryImplementation{Db: db}
}

func (m *MembershipRepositoryImplementation) FindAllForUser(userID string) []response.MembershipResponse {
	var memberships []models.Membership
	result := m.Db.Where("user_id = ?", userID).Preload("Role").Find(&memberships)
	if result.Error != nil {
		return []response.MembershipResponse{}
	}

	var membershipResponses []response.MembershipResponse
	for _, membership := range memberships {
		membershipResponses = append(membershipResponses, membership.MapMembershipToResponse())
	}

	return membershipResponses
}
