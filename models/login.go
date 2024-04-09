package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	UserID   uint   `gorm:"ForeignKey:UserID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *Login) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	l.Password = string(bytes)
	return nil
}
func (l *Login) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(l.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
