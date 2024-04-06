package models

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID               string `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	UserID           uint
	BookingStatusID  uint
	BookingDetailsID uint
	User             User
	BookingStatus    BookingStatus
	BookingDetails   BookingDetails
	BookingCostItems []BookingCostItem
	BookingPayments  []BookingPayment
}

func (b *Booking) TableName() string {
	return "bookings"
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
func (b *Booking) BeforeCreate(db *gorm.DB) error {
	rand := RandStringBytesMask(10)

	b.ID = rand

	return nil

}
