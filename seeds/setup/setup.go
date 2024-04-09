package main

import (
	"booking-api/config"
	requests "booking-api/data/request"
	"booking-api/database"
	"booking-api/models"
	"booking-api/objectStorage"
	"booking-api/repositories"
	"fmt"
	"log"
	"os"
	"time"

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

func SeedTaxRates(db *gorm.DB) {
	taxRatesToSeed := []requests.CreateTaxRateRequest{
		{
			Percentage: 0.10,
			Name:       "Short Term Rental Tax",
		},
		{
			Percentage: 0.06,
			Name:       "Sales Tax",
		},
	}

	taxRateRepository := repositories.NewTaxRateRepositoryImplementation(db)

	for _, taxRate := range taxRatesToSeed {
		taxRateRepository.Create(taxRate)
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
	locations := []string{
		"The Everett Resort",

		"Deer Run Resort",
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

func SeedRoomTypes(db *gorm.DB) {
	roomTypes := []string{

		"Bedroom",
		"Bathroom",

		"Kitchen",

		"Living Room",

		"Dining Room",

		"Patio",

		"Balcony",

		"Theater Room",
	}

	roomTypeRepository := repositories.NewRoomTypeRepositoryImplementation(db)

	for _, roomType := range roomTypes {
		roomTypeRepository.Create(roomType)

	}

}

func SeedUserRoles(db *gorm.DB) {
	userRoles := []string{
		"Admin",
		"Cleaner",
		"Account Owner"}

	userRoleRepository := repositories.NewUserRoleRepositoryImplementation(db)

	for _, userRole := range userRoles {
		userRoleRepository.Create(userRole)
	}

}

func SeedBoats(db *gorm.DB) {
	nineAM := time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC)
	//elevenAM := time.Date(2021, 1, 1, 11, 0, 0, 0, time.UTC)
	threePM := time.Date(2021, 1, 1, 15, 0, 0, 0, time.UTC)

	//fivePM := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)
	boats := []models.Boat{
		{
			// Model: gorm.Model{
			// 	ID: 1,
			// },

			Name:       "The Big Kahuna",
			Occupancy:  10,
			MaxWeight:  2000,
			Timeblocks: []models.Timeblock{},
			Photos: []models.EntityPhoto{
				{
					Photo: models.Photo{
						URL: "boat_photos/1/https://bookingapp.us-ord-1.linodeobjects.com/boat_photos/1/5a1ab150-1ef3-4959-8b5b-085263d9b831.jpeg",
					},
				},
			},
			BookingDurationRule: models.EntityBookingDurationRule{
				MinimumDuration: 2,
				MaximumDuration: 14,
				BookingBuffer:   2,
				StartTime:       nineAM,
				EndTime:         threePM,
			},
			BookingDocuments: []models.EntityBookingDocument{
				{
					Document: models.Document{
						Model: gorm.Model{
							ID: 2,
						},
					},
					RequiresSignature: true,
				},
			},
			BookingCostItems: []models.EntityBookingCost{
				{
					BookingCostType: models.BookingCostType{
						Model: gorm.Model{
							ID: 4,
						},
					},
					TaxRateID: 2,
					Amount:    250,
				},
				{
					BookingCostType: models.BookingCostType{
						Model: gorm.Model{
							ID: 2,
						},
					},
					TaxRateID: 2,
					Amount:    80,
				},
			},
			BookingRule: models.EntityBookingRule{
				AdvertiseAtAllLocations: true,
				AllowPets:               false,
				AllowInstantBooking:     true,
				OfferEarlyCheckIn:       true,
			},
		},
		// {
		// 	// Model: gorm.Model{
		// 	// 	ID: 2,
		// 	// },
		// 	Name:      "The Little Dipper",
		// 	Occupancy: 4,
		// 	MaxWeight: 1000,
		// 	Timeblocks: []models.Timeblock{
		// 		{
		// 			StartTime: elevenAM,
		// 			EndTime:   fivePM,
		// 		},
		// 	},
		// 	Photos: []models.EntityPhoto{
		// 		{
		// 			Photo: models.Photo{

		// 				URL: "boat_photos/2/https://bookingapp.us-ord-1.linodeobjects.com/boat_photos/2/5a1ab150-1ef3-4959-8b5b-085263d9b831.jpeg",
		// 			},
		// 		},
		// 	},
		// 	BookingDurationRule: models.EntityBookingDurationRule{
		// 		MinimumDuration: 3,
		// 		MaximumDuration: 18,
		// 		BookingBuffer:   3,
		// 		StartTime:       elevenAM,
		// 		EndTime:         fivePM,
		// 	},
		// 	BookingCostItems: []models.EntityBookingCost{
		// 		{
		// 			BookingCostType: models.BookingCostType{
		// 				Model: gorm.Model{
		// 					ID: 4,
		// 				},
		// 			},
		// 			TaxRateID: 2,
		// 			Amount:    150,
		// 		},
		// 		{
		// 			BookingCostType: models.BookingCostType{
		// 				Model: gorm.Model{
		// 					ID: 2,
		// 				},
		// 			},
		// 			TaxRateID: 2,
		// 			Amount:    130,
		// 		},
		// 	},
		// },
	}

	boatRepository := repositories.NewBoatRepositoryImplementation(db)

	for _, boat := range boats {
		boatRepository.Create(boat)
	}

}

func SeedRentals(db *gorm.DB) {
	nineAM := time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC)
	//elevenAM := time.Date(2021, 1, 1, 11, 0, 0, 0, time.UTC)
	threePM := time.Date(2021, 1, 1, 15, 0, 0, 0, time.UTC)

	//fivePM := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)
	rentals := []models.Rental{
		{
			// Model: gorm.Model{
			// 	ID: 13,
			// },
			Name:        "The Lodge",
			LocationID:  1,
			Bedrooms:    13,
			Bathrooms:   5,
			Description: "cozy up north cabin",
			Amenities: []models.Amenity{
				{
					Model: gorm.Model{
						ID: 1,
					},
				},
				{
					Model: gorm.Model{
						ID: 40,
					},
				},
			},
			Timeblocks: []models.Timeblock{
				{
					StartTime: nineAM,
					EndTime:   threePM,
				},
			},

			RentalStatus: models.RentalStatus{
				IsClean: true,
			},
			RentalRooms: []models.RentalRoom{
				{
					Name:        "Main bedroom",
					Description: "Master bedroom",
					Floor:       1,
					RoomTypeID:  1,
					Beds: []models.BedType{
						{
							Model: gorm.Model{
								ID: 1,
							},
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
						},
					},
					Photos: []models.EntityPhoto{
						{
							Photo: models.Photo{
								URL: "rental_photos/3/078c6a16-2076-4d1b-88b7-b6e466763aff.PNG",
							},
						},
					},
				},
			},
			BookingDurationRule: models.EntityBookingDurationRule{
				MinimumDuration: 2,
				MaximumDuration: 14,
				BookingBuffer:   2,
				StartTime:       nineAM,
				EndTime:         threePM,
			},
			BookingRule: models.EntityBookingRule{
				AdvertiseAtAllLocations: true,
				AllowPets:               false,
				AllowInstantBooking:     true,
				OfferEarlyCheckIn:       true,
			},
			BookingDocuments: []models.EntityBookingDocument{
				{
					Document: models.Document{
						Model: gorm.Model{
							ID: 2,
						},
					},
					RequiresSignature: true,
				},
			},
			BookingCostItems: []models.EntityBookingCost{
				{
					BookingCostType: models.BookingCostType{
						Model: gorm.Model{
							ID: 3,
						},
					},
					TaxRateID: 1,
					Amount:    1000,
				},
				{
					BookingCostType: models.BookingCostType{
						Model: gorm.Model{
							ID: 2,
						},
					},
					TaxRateID: 2,
					Amount:    100,
				},
			},
			EntityPhotos: []models.EntityPhoto{
				{
					Photo: models.Photo{
						URL: "boat_photos/1/https://bookingapp.us-ord-1.linodeobjects.com/boat_photos/1/5a1ab150-1ef3-4959-8b5b-085263d9b831.jpeg",
					},
				},
			},
		},
	}

	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(db)
	rentalRepository := repositories.NewRentalRepositoryImplementation(db, timeblockRepository)

	for _, rental := range rentals {
		rentalRepository.Create(rental)
	}

}
func SeedDocuments(client *objectStorage.S3Client, db *gorm.DB) {
	documents := []models.Document{
		{
			Name: "Boat Rental Agreement",
			URL:  "documents/boatRentalAgreement.docx",
		},
		{
			Name: "Everett Rental Agreement",
			URL:  "documents/rentalContract.docx",
		},
	}

	documentRepository := repositories.NewDocumentRepositoryImplementation(client, db)

	for _, document := range documents {
		documentRepository.AddDocumentWithUrl(document.URL, document.Name)

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

	// create object storage client
	objectStorage.CreateSession()

	// database.Migrate()

	// SeedBookingStatus(database.Instance)
	// SeedBookingCostTypes(database.Instance)
	// SeedRoomTypes(database.Instance)
	//SeedAmenities(database.Instance)
	//SeedRentals(database.Instance)
	//SeedBoats(database.Instance)
	//SeedDocuments(objectStorage.Client, database.Instance)
	// SeedTaxRates(database.Instance)
	// SeedAmenityTypes(database.Instance)
	// SeedLocations(database.Instance)
	// SeedBedTypes(database.Instance)
	// SeedPaymentMethods(database.Instance)

	//*****users rbac

	SeedUserRoles(database.Instance)

	log.Println("Database seeding Completed!")

}
