package models

import (
	"booking-api/data/response"

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

	Bookings  []Booking                 `gorm:"foreignKey:UserID"`
	Inquiries []EntityBookingPermission `gorm:"foreignKey:UserID"`
	Chats     []Chat                    `gorm:"foreignKey:UserID"`
	Messages  []ChatMessage             `gorm:"foreignKey:UserID"`
}

// func (user *User) BeforeCreate(scope *gorm.DB) error {
// 	//TODO:VERIFY THIS IS probs wronf
// 	user.UserID = uuid.New().String()

// 	fmt.Println("BeforeCreate")
// 	fmt.Println(user.UserID)
// 	return nil
// }

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
	}

	for _, booking := range user.Bookings {
		response.Bookings = append(response.Bookings, booking.MapBookingToResponse())
	}

	return response
}
