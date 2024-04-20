package models

import (
	"booking-api/data/response"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID      string `gorm:" primaryKey"`
	Username    string `json:"username" gorm:"unique"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Login       Login

	Bookings  []Booking `gorm:"foreignKey:UserID"`
	Inquiries []Inquiry `gorm:"foreignKey:UserID"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	//TODO:VERIFY THIS IS probs wronf
	user.UserID = uuid.New().String()

	fmt.Println("BeforeCreate")
	fmt.Println(user.UserID)
	return nil
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
		UserID:      user.UserID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Bookings:    []response.BookingResponse{},
		Inquiries:   []response.InquiryResponse{},
	}

	for _, booking := range user.Bookings {
		response.Bookings = append(response.Bookings, booking.MapBookingToResponse())
	}

	for _, inquiry := range user.Inquiries {
		response.Inquiries = append(response.Inquiries, inquiry.MapInquiryToResponse())
	}

	return response
}
