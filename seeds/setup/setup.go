package main

import (
	"booking-api/config"
	"booking-api/database"
	"booking-api/models"
	"booking-api/repositories"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

func SeedBookingStatus(db *gorm.DB) {
	bookingStauses := []models.BookingStatus{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Drafted",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Requested",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "Confirmed",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			Name: "In Progress",
		},
		{
			Model: gorm.Model{
				ID: 5,
			},
			Name: "Completed",
		},
		{
			Model: gorm.Model{
				ID: 6,
			},
			Name: "Cancelled",
		},
	}

	bookingStatusRepository := repositories.NewBookingStatusRepositoryImplementation(db)

	for _, bookingStatus := range bookingStauses {
		bookingStatusRepository.Create(bookingStatus)
	}

}

func SeedBookingCostTypes(db *gorm.DB) {
	bookingCostTypes := []models.BookingCostType{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Tax",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Cleaning Fee",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "Cabin Rental Cost",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			Name: "Boat Rental Cost",
		},
		{
			Model: gorm.Model{
				ID: 5,
			},
			Name: "Gas Refil fee",
		},
		{
			Model: gorm.Model{
				ID: 6,
			},
			Name: "Labor",
		},
		{
			Model: gorm.Model{
				ID: 7,
			},
			Name: "Damage Fee",
		},
		{
			Model: gorm.Model{
				ID: 8,
			},
			Name: "Wedding Fee",
		},
		{
			Model: gorm.Model{
				ID: 9,
			},
			Name: "Event fee",
		},
		{
			Model: gorm.Model{
				ID: 10,
			},
			Name: "Event Fee Flat",
		},
		{
			Model: gorm.Model{
				ID: 11,
			},
			Name: "Open Bar Fee",
		},
		{
			Model: gorm.Model{
				ID: 12,
			},
			Name: "Cancelation Fee",
		},
		{
			Model: gorm.Model{
				ID: 13,
			},
			Name: "Alcohol",
		},
	}

	bookingCostTypeRepository := repositories.NewBookingCostTypeRepositoryImplementation(db)

	for _, bookingCostType := range bookingCostTypes {
		bookingCostTypeRepository.Create(bookingCostType)
	}

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

	// database.Migrate()

	SeedBookingStatus(database.Instance)
	SeedBookingCostTypes(database.Instance)

	log.Println("Database seeding Completed!")

}
