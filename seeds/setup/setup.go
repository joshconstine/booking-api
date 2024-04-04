package main

import (
	"booking-api/config"
	requests "booking-api/data/request"
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
func SeedPaymentMethods(db *gorm.DB) {
	paymentMethods := []requests.CreatePaymentMethodRequest{
		{
			Name: "Cash",
		},
		{
			Name: "Check",
		},
		{
			Name: "PayPal",
		},
	}

	paymentMethodRepository := repositories.NewPaymentMethodRepositoryImplementation(db)

	for _, paymentMethod := range paymentMethods {
		paymentMethodRepository.Create(paymentMethod)
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

func SeedLocations(db *gorm.DB) {
	locations := []models.Location{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "The Everett Resort",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Deer Run Resort",
		},
	}

	locationRepository := repositories.NewLocationRepositoryImplementation(db)

	for _, location := range locations {
		locationRepository.Create(location)
	}

}
func SeedAmenityTypes(db *gorm.DB) {
	amenityTypes := []models.AmenityType{

		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Kitchen",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Bathroom",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "Laundry",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			Name: "Entertainment",
		},
		{
			Model: gorm.Model{
				ID: 5,
			},
			Name: "Outdoor",
		},
		{
			Model: gorm.Model{
				ID: 6,
			},
			Name: "Utilities",
		},
		{
			Model: gorm.Model{
				ID: 7,
			},
			Name: "Safety",
		},
		{
			Model: gorm.Model{
				ID: 8,
			},
			Name: "Miscellaneous",
		},
	}

	amenityTypeRepository := repositories.NewAmenityTypeRepositoryImplementation(db)

	for _, amenityType := range amenityTypes {
		amenityTypeRepository.Create(amenityType)
	}
}

func SeedBedTypes(db *gorm.DB) {
	bedTypes := []models.BedType{

		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "Twin",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Full",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "Queen",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			Name: "King",
		},
		{
			Model: gorm.Model{
				ID: 5,
			},
			Name: "California King",
		},

		{
			Model: gorm.Model{
				ID: 6,
			},
			Name: "Bunk Bed",
		},
		{
			Model: gorm.Model{
				ID: 7,
			},
			Name: "Sofa Bed",
		},
		{
			Model: gorm.Model{
				ID: 8,
			},
			Name: "Futon",
		},
		{
			Model: gorm.Model{
				ID: 9,
			},
			Name: "Crib",
		},
		{
			Model: gorm.Model{
				ID: 10,
			},
			Name: "Toddler Bed",
		},
		{
			Model: gorm.Model{
				ID: 11,
			},
			Name: "Day Bed",
		},
	}

	bedTypeRepository := repositories.NewBedTypeRepositoryImplementation(db)

	for _, bedType := range bedTypes {
		bedTypeRepository.Create(bedType)
	}

}
func SeedAmenities(db *gorm.DB) {
	amenitiesToSeed := []requests.CreateAmenityRequest{
		{AmenityTypeId: 1, Name: "Refrigerator"},
		{AmenityTypeId: 1, Name: "Microwave"},
		{AmenityTypeId: 1, Name: "Oven"},
		{AmenityTypeId: 1, Name: "Stove"},
		{AmenityTypeId: 1, Name: "Dishwasher"},
		{AmenityTypeId: 1, Name: "Coffee Maker"},
		{AmenityTypeId: 1, Name: "Toaster"},
		{AmenityTypeId: 1, Name: "Blender"},
		{AmenityTypeId: 1, Name: "Food Processor"},
		{AmenityTypeId: 1, Name: "Slow Cooker"},
		{AmenityTypeId: 1, Name: "Stand Mixer"},
		{AmenityTypeId: 1, Name: "Waffle Iron"},
		{AmenityTypeId: 1, Name: "Rice Cooker"},
		{AmenityTypeId: 1, Name: "Electric Kettle"},
		{AmenityTypeId: 2, Name: "Hair Dryer"},
		{AmenityTypeId: 2, Name: "Cleaning Supplies"},
		{AmenityTypeId: 2, Name: "Toilet Paper"},
		{AmenityTypeId: 2, Name: "Shampoo"},
		{AmenityTypeId: 2, Name: "Conditioner"},
		{AmenityTypeId: 2, Name: "Body Wash"},
		{AmenityTypeId: 2, Name: "Hand Soap"},
		{AmenityTypeId: 2, Name: "Towels"},
		{AmenityTypeId: 3, Name: "Washer"},
		{AmenityTypeId: 3, Name: "Dryer"},
		{AmenityTypeId: 3, Name: "Iron"},
		{AmenityTypeId: 3, Name: "Ironing Board"},
		{AmenityTypeId: 4, Name: "TV"},
		{AmenityTypeId: 4, Name: "Cable"},
		{AmenityTypeId: 4, Name: "Netflix"},
		{AmenityTypeId: 4, Name: "Hulu"},
		{AmenityTypeId: 5, Name: "Amazon Prime"},
		{AmenityTypeId: 5, Name: "Apple TV"},
		{AmenityTypeId: 7, Name: "WiFi"},
		{AmenityTypeId: 6, Name: "Patio"},
		{AmenityTypeId: 6, Name: "Balcony"},
		{AmenityTypeId: 6, Name: "Grill"},
		{AmenityTypeId: 6, Name: "Fire Pit"},
		{AmenityTypeId: 7, Name: "Central Air Conditioning"},
		{AmenityTypeId: 7, Name: "Central Heating"},
		{AmenityTypeId: 7, Name: "Fan"},
		{AmenityTypeId: 7, Name: "Space Heater"},
		{AmenityTypeId: 8, Name: "Smoke Detector"},
		{AmenityTypeId: 8, Name: "Carbon Monoxide Detector"},
		{AmenityTypeId: 8, Name: "First Aid Kit"},
		{AmenityTypeId: 8, Name: "Fire Extinguisher"},
		{AmenityTypeId: 9, Name: "Luggage Dropoff Allowed"},
		{AmenityTypeId: 9, Name: "Long Term Stays Allowed"},
		{AmenityTypeId: 9, Name: "Private Entrance"},
	}

	amenityRepository := repositories.NewAmenityRepositoryImplementation(db)

	for _, amenity := range amenitiesToSeed {
		amenityRepository.Create(amenity)
	}

}

func SeedRentals(db *gorm.DB) {
	rentals := []models.Rental{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:        "The Lodge",
			LocationID:  1,
			Bedrooms:    13,
			Bathrooms:   5,
			Description: "cozy up north cabin",
			Amenities:   []models.Amenity{},
			// Timeblocks:  []models.Timeblock{},
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:        "The Morey",
			LocationID:  1,
			Bedrooms:    2,
			Bathrooms:   1,
			Description: "cozy up north cabin",
			Amenities:   []models.Amenity{},
			// Timeblocks:  []models.Timeblock{},
		},
	}

	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(db)
	rentalRepository := repositories.NewRentalRepositoryImplementation(db, timeblockRepository)

	for _, rental := range rentals {
		rentalRepository.Create(rental)
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

	// SeedBookingStatus(database.Instance)
	// SeedBookingCostTypes(database.Instance)
	// SeedRentals(database.Instance)
	// SeedAmenityTypes(database.Instance)
	// SeedLocations(database.Instance)
	// SeedAmenities(database.Instance)
	// SeedBedTypes(database.Instance)
	SeedPaymentMethods(database.Instance)

	log.Println("Database seeding Completed!")

}
