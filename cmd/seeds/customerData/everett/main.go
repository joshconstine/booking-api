package main

import (
	"booking-api/config"
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/models"
	"booking-api/objectStorage"
	"booking-api/pkg/database"
	"booking-api/repositories"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type MemberInput struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Role        string
}

func (m *MemberInput) MapMemberInputToMember() models.Membership {
	return models.Membership{
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		User: models.User{
			Email: m.Email,
		},
		Role: models.UserRole{
			Name: m.Role,
		},
	}
}

type CreateAccountRequest struct {
	AccountName   string
	ServicePlanID uint
	Members       []MemberInput
	Rentals       []request.CreateRentalRequest
	boats         []request.CreateBoatRequest
}

func (c *CreateAccountRequest) MapAccountRequestToAccount() models.Account {
	account := models.Account{
		Name: c.AccountName,
		AccountSettings: models.AccountSettings{
			ServicePlanID: c.ServicePlanID,
		},
		Members: []models.Membership{},
		Rentals: []models.Rental{},
		Boats:   []models.Boat{},
	}
	for _, memberInput := range c.Members {
		account.Members = append(account.Members, memberInput.MapMemberInputToMember())
	}
	for _, rentalInput := range c.Rentals {
		account.Rentals = append(account.Rentals, rentalInput.MapCreateRentalRequestToRental())
	}
	for _, boatInput := range c.boats {
		account.Boats = append(account.Boats, boatInput.MapCreateBoatRequestToBoat())

	}

	return account
}

var accountsToCreate = []CreateAccountRequest{
	{
		AccountName:   "The Everett Reosort",
		ServicePlanID: constants.SERVICE_PLAN_BASIC_ID,
		Members: []MemberInput{
			{
				FirstName:   "Jim",
				LastName:    "Constine",
				PhoneNumber: "7155259214",
				Email:       "Everettmarinabar@outlook.com",
				Role:        constants.USER_ROLE_ACCOUNT_OWNER_NAME,
			},
		},

		Rentals: []request.CreateRentalRequest{
			{
				Name:        "The Lodge",
				LocationID:  1,
				Bedrooms:    13,
				Bathrooms:   7,
				NightlyRate: 500,
				Description: "The Lodge is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Lodge is located on the shores of Lake Everett and is a short walk to the marina. The Lodge is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Lodge is located on the shores of Lake Everett and is a short walk to the marina. The Lodge is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Lodge is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Morey",
				LocationID:  1,
				Bedrooms:    3,
				Bathrooms:   2,
				NightlyRate: 200,
				Description: "The Morey is a cozy up north cabin with 3 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Morey is located on the shores of Lake Everett and is a short walk to the marina. The Morey is a cozy up north cabin with 3 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Morey is located on the shores of Lake Everett and is a short walk to the marina. The Morey is a cozy up north cabin with 3 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Morey is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Gables",
				LocationID:  1,
				Bedrooms:    7,
				Bathrooms:   4,
				NightlyRate: 300,
				Description: "The Gables is a cozy up north cabin with 7 bedrooms and 4 bathrooms. It is perfect for medium groups and family gatherings. The Gables is located on the shores of Lake Everett and is a short walk to the marina. The Gables is a cozy up north cabin with 7 bedrooms and 4 bathrooms. It is perfect for medium groups and family gatherings. The Gables is located on the shores of Lake Everett and is a short walk to the marina. The Gables is a cozy up north cabin with 7 bedrooms and 4 bathrooms. It is perfect for medium groups and family gatherings. The Gables is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Clubhouse",
				LocationID:  1,
				Bedrooms:    5,
				Bathrooms:   3,
				NightlyRate: 250,
				Description: "The Clubhouse is a cozy up north cabin with 5 bedrooms and 3 bathrooms. It is perfect for medium groups and family gatherings. The Clubhouse is located on the shores of Lake Everett and is a short walk to the marina. The Clubhouse is a cozy up north cabin with 5 bedrooms and 3 bathrooms. It is perfect for medium groups and family gatherings. The Clubhouse is located on the shores of Lake Everett and is a short walk to the marina. The Clubhouse is a cozy up north cabin with 5 bedrooms and 3 bathrooms. It is perfect for medium groups and family gatherings. The Clubhouse is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Eisenhower",
				LocationID:  1,
				Bedrooms:    4,
				Bathrooms:   2,
				NightlyRate: 150,
				Description: "The Eisenhower is a cozy up north cabin with 4 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Eisenhower is located on the shores of Lake Everett and is a short walk to the marina. The Eisenhower is a cozy up north cabin with 4 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Eisenhower is located on the shores of Lake Everett and is a short walk to the marina. The Eisenhower is a cozy up north cabin with 4 bedrooms and 2 bathrooms. It is perfect for small groups and family gatherings. The Eisenhower is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Musky Inn",
				LocationID:  2,
				Bedrooms:    13,
				Bathrooms:   7,
				NightlyRate: 500,
				Description: "The Musky Inn is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Musky Inn is located on the shores of Lake Everett and is a short walk to the marina. The Musky Inn is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Musky Inn is located on the shores of Lake Everett and is a short walk to the marina. The Musky Inn is a cozy up north cabin with 13 bedrooms and 7 bathrooms. It is perfect for large groups and family gatherings. The Musky Inn is located on the shores of Lake Everett and is a short walk to the marina.",
			},
			{
				Name:        "The Little Guy",
				LocationID:  2,
				Bedrooms:    1,
				Bathrooms:   1,
				NightlyRate: 100,
				Description: "The Little Guy is a cozy up north cabin with 1 bedroom and 1 bathroom. It is perfect for small groups and family gatherings. The Little Guy is located on the shores of Lake Everett and is a short walk to the marina. The Little Guy is a cozy up north cabin with 1 bedroom and 1 bathroom. It is perfect for small groups and family gatherings. The Little Guy is located on the shores of Lake Everett and is a short walk to the marina. The Little Guy is a cozy up north cabin with 1 bedroom and 1 bathroom. It is perfect for small groups and family gatherings. The Little Guy is located on the shores of Lake Everett and is a short walk to the marina.",
			},
		},

		boats: []request.CreateBoatRequest{},
	},
	{
		AccountName:   "St Germain Boat Rentals",
		ServicePlanID: constants.SERVICE_PLAN_BASIC_ID,
		Members: []MemberInput{
			{
				FirstName:   "Jim",
				LastName:    "Constine",
				PhoneNumber: "7155259214",
				Email:       "Everettmarinabar@outlook.com",
				Role:        constants.USER_ROLE_ACCOUNT_OWNER_NAME,
			},
		},
		Rentals: []request.CreateRentalRequest{},
		boats: []request.CreateBoatRequest{
			{
				Name:        "the party barge",
				NightlyRate: 500,
				Description: "be the king of the chain with this party barge",
			},
		},
	},
}

func SeedAccountsFromInput(db *gorm.DB) {
	accountRepository := repositories.NewAccountRepositoryImplementation(db)
	for _, accountInput := range accountsToCreate {
		accountRepository.Create(accountInput.MapAccountRequestToAccount())

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
			Status: models.BoatStatus{
				IsClean:    true,
				LocationID: 1,
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

// func SeedRentals(db *gorm.DB) {
// 	nineAM := time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC)
// 	//elevenAM := time.Date(2021, 1, 1, 11, 0, 0, 0, time.UTC)
// 	threePM := time.Date(2021, 1, 1, 15, 0, 0, 0, time.UTC)

// 	twoWeeksFromNow := time.Now().AddDate(0, 0, 14)
// 	threeWeeksFromNow := time.Now().AddDate(0, 0, 21)

// 	sixWeeksFromNow := time.Now().AddDate(0, 0, 42)
// 	seventyDaysFromNow := time.Now().AddDate(0, 0, 70)

// 	// startDate, endDate := GetRandomDateRangeWithenTheNextYear(2, 14)

// 	//fivePM := time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC)
// 	// rentals := []models.Rental{
// 	// 	{
// 	// 		// Model: gorm.Model{
// 	// 		// 	ID: 13,
// 	// 		// },
// 	// 		Name:        "The Lodge",
// 	// 		LocationID:  1,
// 	// 		Bedrooms:    13,
// 	// 		Bathrooms:   5,
// 	// 		Description: "cozy up north cabin",
// 	// 		Amenities: []models.Amenity{
// 	// 			{
// 	// 				Model: gorm.Model{
// 	// 					ID: 1,
// 	// 				},
// 	// 			},
// 	// 			{
// 	// 				Model: gorm.Model{
// 	// 					ID: 40,
// 	// 				},
// 	// 			},
// 	// 		},
// 	// 		Timeblocks: []models.EntityTimeblock{
// 	// 			{
// 	// 				StartTime: nineAM,
// 	// 				EndTime:   threePM,
// 	// 			},
// 	// 		},

// 	// 		RentalStatus: models.RentalStatus{
// 	// 			IsClean: true,
// 	// 		},
// 	// 		RentalRooms: []models.RentalRoom{
// 	// 			{
// 	// 				Name:        "Main bedroom",
// 	// 				Description: "Master bedroom",
// 	// 				Floor:       1,
// 	// 				RoomTypeID:  1,
// 	// 				Beds: []models.BedType{
// 	// 					{
// 	// 						Model: gorm.Model{
// 	// 							ID: 1,
// 	// 						},
// 	// 					},
// 	// 					{
// 	// 						Model: gorm.Model{
// 	// 							ID: 2,
// 	// 						},
// 	// 					},
// 	// 				},
// 	// 				Photos: []models.EntityPhoto{
// 	// 					{
// 	// 						Photo: models.Photo{
// 	// 							URL: "rental_photos/3/078c6a16-2076-4d1b-88b7-b6e466763aff.PNG",
// 	// 						},
// 	// 					},
// 	// 				},
// 	// 			},
// 	// 		},
// 	// 		BookingDurationRule: models.EntityBookingDurationRule{
// 	// 			MinimumDuration: 2,
// 	// 			MaximumDuration: 14,
// 	// 			BookingBuffer:   2,
// 	// 			StartTime:       nineAM,
// 	// 			EndTime:         threePM,
// 	// 		},
// 	// 		BookingRule: models.EntityBookingRule{
// 	// 			AdvertiseAtAllLocations: true,
// 	// 			AllowPets:               false,
// 	// 			AllowInstantBooking:     true,
// 	// 			OfferEarlyCheckIn:       true,
// 	// 		},
// 	// 		BookingDocuments: []models.EntityBookingDocument{
// 	// 			{
// 	// 				Document: models.Document{
// 	// 					Model: gorm.Model{
// 	// 						ID: 2,
// 	// 					},
// 	// 				},
// 	// 				RequiresSignature: true,
// 	// 			},
// 	// 		},
// 	// 		BookingCostItems: []models.EntityBookingCost{
// 	// 			{
// 	// 				BookingCostType: models.BookingCostType{
// 	// 					Model: gorm.Model{
// 	// 						ID: 3,
// 	// 					},
// 	// 				},
// 	// 				TaxRateID: 1,
// 	// 				Amount:    1000,
// 	// 			},
// 	// 			{
// 	// 				BookingCostType: models.BookingCostType{
// 	// 					Model: gorm.Model{
// 	// 						ID: 2,
// 	// 					},
// 	// 				},
// 	// 				TaxRateID: 2,
// 	// 				Amount:    100,
// 	// 			},
// 	// 		},
// 	// 		BookingCostItemAdjustments: []models.EntityBookingCostAdjustment{
// 	// 			{
// 	// 				Amount:            1500,
// 	// 				BookingCostTypeID: 3,
// 	// 				TaxRateID:         1,

// 	// 				StartDate: twoWeeksFromNow,
// 	// 				EndDate:   threeWeeksFromNow,
// 	// 			},
// 	// 			{
// 	// 				Amount:            2000,
// 	// 				BookingCostTypeID: 3,
// 	// 				TaxRateID:         1,

// 	// 				StartDate: sixWeeksFromNow,
// 	// 				EndDate:   seventyDaysFromNow,
// 	// 			},
// 	// 		},
// 	// 		EntityPhotos: []models.EntityPhoto{
// 	// 			{
// 	// 				Photo: models.Photo{
// 	// 					URL: "boat_photos/1/https://bookingapp.us-ord-1.linodeobjects.com/boat_photos/1/5a1ab150-1ef3-4959-8b5b-085263d9b831.jpeg",
// 	// 				},
// 	// 			},
// 	// 		},
// 	// 	},
// 	// }

// 	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(db)
// 	rentalRepository := repositories.NewRentalRepositoryImplementation(db, timeblockRepository)

// 	for _, rental := range rentals {
// 		rentalRepository.Create(rental)
// 	}

// }
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

	// SeedAmenityTypes(database.Instance)
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
	//SeedInquiries(database.Instance)
	SeedAccountsFromInput(database.Instance)

	// SeedRentals(database.Instance)
	// SeedBoats(database.Instance)

	//*****users rbac

	log.Println("Database seeding Completed!")

}
