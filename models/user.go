package models

import (
	"booking-api/data/response"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:" primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Username       string         `json:"username" gorm:"unique"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	PreferedName   string         `json:"prefferedName"`
	Email          string         `json:"email" gorm:"unique"`
	PhoneNumber    string         `json:"phoneNumber"`
	Gender         string         `json:"gender"`
	DOB            time.Time      `json:"dob"`
	ProfilePicture string         `json:"profilePicture"`
	AddressID      uint
	Login          Login

	Bookings  []Booking
	Inquiries []EntityBookingPermission
	Chats     []Chat
	Messages  []ChatMessage
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
		UserID:             user.ID,
		Username:           user.Username,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		PrefferedName:      user.PreferedName,
		Bookings:           []response.BookingResponse{},
		Chats:              []response.ChatResponse{},
		PermissionRequests: []response.EntityBookingPermissionResponse{},
	}

	for _, booking := range user.Bookings {
		response.Bookings = append(response.Bookings, booking.MapBookingToResponse())
	}

	for _, chat := range user.Chats {
		response.Chats = append(response.Chats, chat.MapChatToResponse())
	}

	response.ProfilePicture = MakeUrlPathForObjectStorage(user.ProfilePicture)

	return response
}
