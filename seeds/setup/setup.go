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

func SeedAmenities(db *gorm.DB) {
	amenitiesToSeed := []models.Amenity{
		{Model: gorm.Model{
			ID: 1,
		}, AmenityTypeID: 1, Name: "Refrigerator"},
		{Model: gorm.Model{
			ID: 2,
		}, AmenityTypeID: 1, Name: "Microwave"},
		{Model: gorm.Model{
			ID: 3,
		}, AmenityTypeID: 1, Name: "Oven"},
		{Model: gorm.Model{
			ID: 4,
		}, AmenityTypeID: 1, Name: "Stove"},
		{Model: gorm.Model{
			ID: 5,
		}, AmenityTypeID: 1, Name: "Dishwasher"},
		{Model: gorm.Model{
			ID: 6,
		}, AmenityTypeID: 1, Name: "Coffee Maker"},
		{Model: gorm.Model{
			ID: 7,
		}, AmenityTypeID: 1, Name: "Toaster"},
		{Model: gorm.Model{
			ID: 8,
		}, AmenityTypeID: 1, Name: "Blender"},
		{Model: gorm.Model{
			ID: 9,
		}, AmenityTypeID: 1, Name: "Food Processor"},
		{Model: gorm.Model{
			ID: 10,
		}, AmenityTypeID: 1, Name: "Slow Cooker"},
		{Model: gorm.Model{
			ID: 11,
		}, AmenityTypeID: 1, Name: "Stand Mixer"},
		{Model: gorm.Model{
			ID: 12,
		}, AmenityTypeID: 1, Name: "Waffle Iron"},
		{Model: gorm.Model{
			ID: 13,
		}, AmenityTypeID: 1, Name: "Rice Cooker"},
		{Model: gorm.Model{
			ID: 14,
		}, AmenityTypeID: 1, Name: "Electric Kettle"},
		{Model: gorm.Model{
			ID: 15,
		}, AmenityTypeID: 2, Name: "Hair Dryer"},
		{Model: gorm.Model{
			ID: 16,
		}, AmenityTypeID: 2, Name: "Cleaning Supplies"},
		{Model: gorm.Model{
			ID: 17,
		}, AmenityTypeID: 2, Name: "Toilet Paper"},
		{Model: gorm.Model{
			ID: 18,
		}, AmenityTypeID: 2, Name: "Shampoo"},
		{Model: gorm.Model{
			ID: 19,
		}, AmenityTypeID: 2, Name: "Conditioner"},
		{Model: gorm.Model{
			ID: 20,
		}, AmenityTypeID: 2, Name: "Body Wash"},
		{Model: gorm.Model{
			ID: 21,
		}, AmenityTypeID: 2, Name: "Hand Soap"},
		{Model: gorm.Model{
			ID: 22,
		}, AmenityTypeID: 2, Name: "Towels"},
		{Model: gorm.Model{
			ID: 23,
		}, AmenityTypeID: 3, Name: "Washer"},
		{Model: gorm.Model{
			ID: 24,
		}, AmenityTypeID: 3, Name: "Dryer"},
		{Model: gorm.Model{
			ID: 25,
		}, AmenityTypeID: 3, Name: "Iron"},
		{Model: gorm.Model{
			ID: 26,
		}, AmenityTypeID: 3, Name: "Ironing Board"},
		{Model: gorm.Model{
			ID: 27,
		}, AmenityTypeID: 4, Name: "TV"},
		{Model: gorm.Model{
			ID: 28,
		}, AmenityTypeID: 4, Name: "Cable"},
		{Model: gorm.Model{
			ID: 29,
		}, AmenityTypeID: 4, Name: "Netflix"},
		{Model: gorm.Model{
			ID: 30,
		}, AmenityTypeID: 4, Name: "Hulu"},
		{Model: gorm.Model{
			ID: 31,
		}, AmenityTypeID: 4, Name: "Amazon Prime"},
		{Model: gorm.Model{
			ID: 32,
		}, AmenityTypeID: 4, Name: "Apple TV"},
		{Model: gorm.Model{
			ID: 33,
		}, AmenityTypeID: 6, Name: "WiFi"},
		{Model: gorm.Model{
			ID: 34,
		}, AmenityTypeID: 5, Name: "Patio"},
		{Model: gorm.Model{
			ID: 35,
		}, AmenityTypeID: 5, Name: "Balcony"},
		{Model: gorm.Model{
			ID: 36,
		}, AmenityTypeID: 5, Name: "Grill"},
		{Model: gorm.Model{
			ID: 37,
		}, AmenityTypeID: 5, Name: "Fire Pit"},
		{Model: gorm.Model{
			ID: 38,
		}, AmenityTypeID: 6, Name: "Central Air Conditioning"},
		{Model: gorm.Model{
			ID: 39,
		}, AmenityTypeID: 6, Name: "Central Heating"},
		{Model: gorm.Model{
			ID: 40,
		}, AmenityTypeID: 6, Name: "Fan"},
		{Model: gorm.Model{
			ID: 41,
		}, AmenityTypeID: 6, Name: "Space Heater"},
		{Model: gorm.Model{
			ID: 42,
		}, AmenityTypeID: 7, Name: "Smoke Detector"},
		{Model: gorm.Model{
			ID: 43,
		}, AmenityTypeID: 7, Name: "Carbon Monoxide Detector"},
		{Model: gorm.Model{
			ID: 44,
		}, AmenityTypeID: 7, Name: "First Aid Kit"},
		{Model: gorm.Model{
			ID: 45,
		}, AmenityTypeID: 7, Name: "Fire Extinguisher"},
		{Model: gorm.Model{
			ID: 46,
		}, AmenityTypeID: 8, Name: "Luggage Dropoff Allowed"},
		{Model: gorm.Model{
			ID: 47,
		}, AmenityTypeID: 8, Name: "Long Term Stays Allowed"},
		{Model: gorm.Model{
			ID: 48,
		}, AmenityTypeID: 8, Name: "Private Entrance"},
	}

	amenityRepository := repositories.NewAmenityRepositoryImplementation(db)

	for _, amenity := range amenitiesToSeed {
		amenityRepository.Create(amenity)
	}

}

// {1, "The Lodge", 1, 13, 5, "cozy up north cabin"},
// {2, "The Morey", 1, 2, 1, "cozy up north cabin"},
// {3, "The Gables", 1, 7, 3, "cozy up north cabin"},
// {4, "The Clubhouse", 1, 5, 2, "cozy up north cabin"},
// {5, "The Eisenhower", 1, 4, 2, "cozy up north cabin"},
// {6, "The Musky Inn", 2, 13, 7, "cozy up north cabin"},
// // {7, "The Musky Inn North", 2, 6, 4, "cozy up north cabin"},
// // {8, "The Musky Inn North + middle", 2, 9, 4, "cozy up north cabin"},
// // {9, "The Musky Inn South", 2, 4, 3, "cozy up north cabin"},
// // {10, "The Musky Inn South + middle", 2, 7, 5, "cozy up north cabin"},
// {7, "The Little Guy", 2, 1, 1, "cozy up north cabin"},

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
	SeedAmenities(database.Instance)

	log.Println("Database seeding Completed!")

}
