package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type Membership struct {
	gorm.Model
	AccountID   uint `gorm:"not null"`
	UserID      uint `gorm:"not null"`
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

		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,

		Role: m.Role.MapUserRoleToResponse(),
		User: m.User.MapUserToResponse(),
	}
}

func MapMembershipsToResponses(memberships []Membership) []response.MembershipResponse {
	var responses []response.MembershipResponse
	for _, membership := range memberships {
		responses = append(responses, membership.MapMembershipToResponse())
	}
	return responses
}
