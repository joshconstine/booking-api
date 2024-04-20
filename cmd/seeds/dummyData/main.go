package main

import (
	"booking-api/config"
	"booking-api/data/request"
	"booking-api/pkg/database"
	"booking-api/repositories"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

func GetRandomDateRangeWithenTheNextYear(minLength int, maxLength int) (time.Time, time.Time) {
	min := time.Now()
	max := time.Now().AddDate(0, 0, 365)
	delta := max.Unix() - min.Unix()
	sec := min.Unix() + int64(rand.Intn(int(delta)))
	return time.Unix(sec, 0), time.Unix(sec, 0).AddDate(0, 0, rand.Intn(maxLength-minLength)+minLength)
}

func SeedUsers(db *gorm.DB) {
	// create users
	userRepository := repositories.NewUserRepositoryImplementation(db)
	usersToCreate := request.CreateUserRequest{
		Email:       gofakeit.Email(),
		Username:    gofakeit.Username(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
	}

	userRepository.Create(&usersToCreate)
}
func main() {

	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()
	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	database.Connect(env.DSN)

	// create object storage client
	// objectStorage.CreateSession()
	SeedUsers(database.Instance)

	log.Println("Database seeding Completed!")

}
