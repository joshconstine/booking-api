package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"unique"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phoneNumber"`
	Login       Login
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
