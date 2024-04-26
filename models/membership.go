package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	AccountID   uint   `gorm:"not null"`
	UserID      string `gorm:"not null"`
	PhoneNumber string
	Email       string
	RoleID      uint `gorm:"not null"`
	Role        UserRole
	User        User
}

func (m *Membership) TableName() string {
	return "memberships"
}

func (m *Membership) MapMembershipToResponse() response.MembershipResponse {
	return response.MembershipResponse{
		ID: m.ID,

		AccountID:   m.AccountID,
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		Role:        m.Role.MapUserRoleToResponse(),
		UserID:      m.UserID,
	}
}

func MapMembershipsToResponses(memberships []Membership) []response.MembershipResponse {
	var responses []response.MembershipResponse
	for _, membership := range memberships {
		responses = append(responses, membership.MapMembershipToResponse())
	}
	return responses
}
