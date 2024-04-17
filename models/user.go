package models

import (
	"booking-api/data/response"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      uuid.UUID `gorm:" primaryKey"`
	Username    string    `json:"username" gorm:"unique"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email" gorm:"unique"`
	PhoneNumber string    `json:"phoneNumber"`
	Login       Login

	Bookings  []Booking `gorm:"foreignKey:UserID"`
	Inquiries []Inquiry `gorm:"foreignKey:UserID"`
}

func (user *User) CheckPassword(providedPassword string) error {
	err := user.Login.CheckPassword(providedPassword)
	if err != nil {
		return err
	}
	return nil
}
func (user *User) HashPassword(password string) error {
	err := user.Login.HashPassword(password)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) MapUserToResponse() response.UserResponse {
	response := response.UserResponse{
		ID:       user.ID,
		Username: user.Username,

		Email: user.Email,
	}
	return response
}
