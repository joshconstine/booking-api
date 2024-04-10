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
	"math/rand"
	"os"
	"time"

	"gorm.io/gorm"
)

func GetRandomDateRangeWithenTheNextYear(minLength int, maxLength int) (time.Time, time.Time) {
	min := time.Now()
	max := time.Now().AddDate(0, 0, 365)
	delta := max.Unix() - min.Unix()
	sec := min.Unix() + int64(rand.Intn(int(delta)))
	return time.Unix(sec, 0), time.Unix(sec, 0).AddDate(0, 0, rand.Intn(maxLength-minLength)+minLength)
}

func SeedInquiryStatuses(db *gorm.DB) {
	inquiryStatuses := []models.InquiryStatus{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name: "New",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name: "Approved",
		},
		{
			Model: gorm.Model{
				ID: 3,
			},
			Name: "Approval Expired",
		},
		{
			Model: gorm.Model{
				ID: 4,
			},
			Name: "Declined",
		},
		{
			Model: gorm.Model{
				ID: 5,
			},
			Name: "Cancelled",
		},
	}

	inquiryStatusRepository := repositories.NewInquiryStatusRepositoryImplementation(db)

	for _, inquiryStatus := range inquiryStatuses {
		inquiryStatusRepository.Create(inquiryStatus)
	}

}

func SeedInquiries(db *gorm.DB) {
	userId := uint(24)
	entityId := uint(29)
	boatEntityId := uint(13)
	inquiries := []models.Inquiry{
		{

			UserID:    userId,
			Note:      "Do you have a toaster?",
			NumGuests: 2,
			BookingRequests: []models.EntityBookingRequest{
				{
					EntityID:   entityId,
					EntityType: "rentals",
					StartTime:  time.Now(),
					EndTime:    time.Now().AddDate(0, 0, 1),
				},
				{
					EntityID:   boatEntityId,
					EntityType: "boats",
					StartTime:  time.Now(),
					EndTime:    time.Now().AddDate(0, 0, 1),
				},
			},
		},
	}

	inquiryRepository := repositories.NewInquiryRepositoryImplementation(db)

	for _, inquiry := range inquiries {
		inquiryRepository.Create(inquiry)
	}

}

func SeedAccounts(db *gorm.DB) {
	accountOwnerEmail := "l@gmail.com"
	cleanerEmail := "k@gmail.com"
	accounts := []models.Account{
		{

			Name: "The Everett Resort",
			Members: []models.Membership{
				{
					User: models.User{
						Email: accountOwnerEmail,
					},
					Role: models.UserRole{
						Model: gorm.Model{
							ID: 3,
						},
					},
					PhoneNumber: "1234567890",
					Email:       accountOwnerEmail,
				},
				{
					User: models.User{
						Email: cleanerEmail,
					},
					Role: models.UserRole{
						Model: gorm.Model{
							ID: 2,
						},
					},
					PhoneNumber: "92026533333",
					Email:       accountOwnerEmail,
				},
			},
			AccountSettings: models.AccountSettings{
				AccountOwner: models.Membership{
					User: models.User{
						Email: accountOwnerEmail,
					},
				},

				ServicePlan: models.ServicePlan{
					Model: gorm.Model{
						ID: 1,
					},
					Fees: []models.ServiceFee{
						{
							FeePercentage:         0.05,
							AppliesToAllCostTypes: true,
						},
					},
				},
			},
		},
		{

			Name: "St Germain Boat Rentals",
			Members: []models.Membership{
				{
					User: models.User{
						Email: accountOwnerEmail,
					},
					Role: models.UserRole{
						Model: gorm.Model{
							ID: 3,
						},
					},
					PhoneNumber: "1234567890",
					Email:       accountOwnerEmail,
				},
			},
			AccountSettings: models.AccountSettings{
				ServicePlan: models.ServicePlan{
					Model: gorm.Model{
						ID: 1,
					},
					Fees: []models.ServiceFee{
						{
							FeePercentage:         0.05,
							AppliesToAllCostTypes: true,
						},
					},
				},
			},
		},
	}

	accountRepository := repositories.NewAccountRepositoryImplementation(db)

	for _, account := range accounts {
		accountRepository.Create(account)

	}

}

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
	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
	threeWeeksFromNow := time.Now().AddDate(0, 0, 21)

	sixWeeksFromNow := time.Now().AddDate(0, 0, 42)
	seventyDaysFromNow := time.Now().AddDate(0, 0, 70)
	//fivePM := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)
	boats := []models.Boat{
		{
			// Model: gorm.Model{
			// 	ID: 1,
			// },

			Name:       "The Big Kahuna",
			Occupancy:  10,
			MaxWeight:  2000,
			Timeblocks: []models.EntityTimeblock{},
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
			BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{
				{
					Amount:            1500,
					BookingCostTypeID: 4,
					TaxRateID:         1,

					StartDate: twoWeeksFromNow,
					EndDate:   threeWeeksFromNow,
				},
				{
					Amount:            2000,
					BookingCostTypeID: 4,
					TaxRateID:         1,

					StartDate: sixWeeksFromNow,
					EndDate:   seventyDaysFromNow,
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

	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
	threeWeeksFromNow := time.Now().AddDate(0, 0, 21)

	sixWeeksFromNow := time.Now().AddDate(0, 0, 42)
	seventyDaysFromNow := time.Now().AddDate(0, 0, 70)

	// startDate, endDate := GetRandomDateRangeWithenTheNextYear(2, 14)

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
			Timeblocks: []models.EntityTimeblock{
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
			BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{
				{
					Amount:            1500,
					BookingCostTypeID: 3,
					TaxRateID:         1,

					StartDate: twoWeeksFromNow,
					EndDate:   threeWeeksFromNow,
				},
				{
					Amount:            2000,
					BookingCostTypeID: 3,
					TaxRateID:         1,

					StartDate: sixWeeksFromNow,
					EndDate:   seventyDaysFromNow,
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

	//database.Migrate()
	// SeedAmenityTypes(database.Instance)
	// SeedBedTypes(database.Instance)
	// SeedBookingCostTypes(database.Instance)
	// SeedUserRoles(database.Instance)
	// SeedBookingStatus(database.Instance)
	// SeedDocuments(objectStorage.Client, database.Instance)
	// SeedRoomTypes(database.Instance)
	// SeedAmenities(database.Instance)
	// SeedPaymentMethods(database.Instance)
	// SeedTaxRates(database.Instance)
	// SeedLocations(database.Instance)

	// SeedInquiryStatuses(database.Instance)
	// SeedAccounts(database.Instance)
	// SeedInquiries(database.Instance)

	// SeedRentals(database.Instance)
	SeedBoats(database.Instance)

	//*****users rbac

	log.Println("Database seeding Completed!")

}
