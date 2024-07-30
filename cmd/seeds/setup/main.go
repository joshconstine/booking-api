package main

import (
	"booking-api/config"
	"booking-api/constants"
	requests "booking-api/data/request"
	"booking-api/models"
	"booking-api/pkg/database"
	"booking-api/pkg/objectStorage"
	"booking-api/repositories"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

func SeedInquiryStatuses(db *gorm.DB) {
	inquiryStatuses := []models.InquiryStatus{
		{
			Model: gorm.Model{
				ID: uint(constants.INQUIRY_STATUS_NEW_ID),
			},
			Name: constants.INQUIRY_STATUS_NEW_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.INQUIRY_STATUS_APPROVED_ID),
			},
			Name: constants.INQUIRY_STATUS_APPROVED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.INQUIRY_STATUS_APPROVAL_EXPIRED_ID),
			},
			Name: constants.INQUIRY_STATUS_APPROVAL_EXPIRED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.INQUIRY_STATUS_DECLINED_ID),
			},
			Name: constants.INQUIRY_STATUS_DECLINED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.INQUIRY_STATUS_CANCELLED_ID),
			},
			Name: constants.INQUIRY_STATUS_CANCELLED_NAME,
		},
	}

	inquiryStatusRepository := repositories.NewInquiryStatusRepositoryImplementation(db)

	for _, inquiryStatus := range inquiryStatuses {
		inquiryStatusRepository.Create(inquiryStatus)
	}

}
func SeedBookingStatus(db *gorm.DB) {
	bookingStauses := []models.BookingStatus{
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_DRAFTED_ID),
			},
			Name: constants.BOOKING_STATUS_DRAFTED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_REQUESTED_ID),
			},
			Name: constants.BOOKING_STATUS_REQUESTED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_CONFIRMED_ID),
			},
			Name: constants.BOOKING_STATUS_CONFIRMED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_IN_PROGRESS_ID),
			},
			Name: constants.BOOKING_STATUS_IN_PROGRESS_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_COMPLETED_ID),
			},
			Name: constants.BOOKING_STATUS_COMPLETED_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_STATUS_CANCELLED_ID),
			},
			Name: constants.BOOKING_STATUS_CANCELLED_NAME,
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
			Percentage: 12,
			Name:       "Short Term Rental Tax",
		},
		{
			Percentage: 06,
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
			Name: "Stripe",
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
				ID: uint(constants.BOOKING_COST_TYPE_TAX_ID),
			},
			Name: constants.BOOKING_COST_TYPE_TAX_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_CLEANING_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_CLEANING_FEE_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_RENTAL_COST_ID),
			},
			Name: constants.BOOKING_COST_TYPE_RENTAL_COST_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_BOAT_RENTAL_COST_ID),
			},
			Name: constants.BOOKING_COST_TYPE_BOAT_RENTAL_COST_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_GAS_REFIL_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_GAS_REFIL_FEE_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_LABOR_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_LABOR_FEE_NAME,
		},
		{

			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_DAMAGE_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_DAMAGE_FEE_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_WEDDING_FEE_ID),
			},

			Name: constants.BOOKING_COST_TYPE_WEDDING_FEE_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_EVENT_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_EVENT_FEE_NAME,
		},
		{
			Model: gorm.Model{

				ID: uint(constants.BOOKING_COST_TYPE_OPEN_BAR_ID),
			},

			Name: constants.BOOKING_COST_TYPE_OPEN_BAR_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.BOOKING_COST_TYPE_CANCELLATION_FEE_ID),
			},
			Name: constants.BOOKING_COST_TYPE_CANCELLATION_FEE_NAME,
		},
	}

	bookingCostTypeRepository := repositories.NewBookingCostTypeRepositoryImplementation(db)

	for _, bookingCostType := range bookingCostTypes {
		bookingCostTypeRepository.Create(bookingCostType)
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
	amenitiesToSeed := []requests.CreateAmenityRequest{
		{AmenityTypeId: 1, Name: "Refrigerator"},
		{AmenityTypeId: 1, Name: "Microwave"},
		{AmenityTypeId: 1, Name: "Oven"},
		{AmenityTypeId: 1, Name: "Stove"},
		{AmenityTypeId: 1, Name: "Dishwasher"},
		{AmenityTypeId: 1, Name: "Coffee Maker"},
		{AmenityTypeId: 1, Name: "Toaster"},
		{AmenityTypeId: 1, Name: "Blender"},
		{AmenityTypeId: 1, Name: "Slow Cooker"},
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
		constants.ROOM_TYPE_BEDROOM_NAME,
		constants.ROOM_TYPE_BATHROOM_NAME,
		constants.ROOM_TYPE_KITCHEN_NAME,
		constants.ROOM_TYPE_LIVING_ROOM_NAME,
		constants.ROOM_TYPE_DINING_ROOM_NAME,
		constants.ROOM_TYPE_GARAGE_NAME,
		constants.ROOM_TYPE_PATIO_NAME,
		constants.ROOM_TYPE_ENTERTAINMENT_ROOM_NAME,
	}

	roomTypeRepository := repositories.NewRoomTypeRepositoryImplementation(db)

	for _, roomType := range roomTypes {
		roomTypeRepository.Create(roomType)

	}

}

func SeedUserRoles(db *gorm.DB) {
	userRoles := []models.UserRole{
		{
			Model: gorm.Model{
				ID: uint(constants.USER_ROLE_ADMIN_ID),
			},
			Name: constants.USER_ROLE_ADMIN_NAME,
		},
		{
			Model: gorm.Model{
				ID: uint(constants.USER_ROLE_ACCOUNT_OWNER_ID),
			},
			Name: constants.USER_ROLE_ACCOUNT_OWNER_NAME,
		},
		{
			Model: gorm.Model{
				ID: constants.USER_ROLE_ACCOUNT_MANAGER_ID,
			},
			Name: constants.USER_ROLE_ACCOUNT_MANAGER_NAME,
		},
		{
			Model: gorm.Model{
				ID: constants.USER_ROLE_CLEANING_STAFF_ID,
			},
			Name: constants.USER_ROLE_CLEANING_STAFF_NAME,
		},
		{
			Model: gorm.Model{
				ID: constants.USER_ROLE_MAINTENANCE_STAFF_ID,
			},
			Name: constants.USER_ROLE_MAINTENANCE_STAFF_NAME,
		},
	}

	userRoleRepository := repositories.NewUserRoleRepositoryImplementation(db)

	for _, userRole := range userRoles {
		userRoleRepository.Create(&userRole)
	}

}

func SeedServicePlans(db *gorm.DB) {
	servicePlans := []models.ServicePlan{
		{
			Model: gorm.Model{
				ID: uint(constants.SERVICE_PLAN_BASIC_ID),
			},
			Name: constants.SERVICE_PLAN_BASIC_NAME,
			Fees: []models.ServiceFee{{

				FeePercentage:         5,
				AppliesToAllCostTypes: true,
			},
				{
					FeePercentage:         2,
					AppliesToAllCostTypes: false,
					BookingCostTypeID:     uint(constants.BOOKING_COST_TYPE_BOAT_RENTAL_COST_ID),
				},
			},
		},
		{
			Model: gorm.Model{
				ID: uint(constants.SERVICE_PLAN_PRO_HOST_ID),
			},
			Name: constants.SERVICE_PLAN_PRO_HOST_NAME,
			Fees: []models.ServiceFee{{

				FeePercentage:         3,
				AppliesToAllCostTypes: true,
			},
			},
		},
		{
			Model: gorm.Model{
				ID: uint(constants.SERVICE_PLAN_FULLY_MANAGED_ID),
			},
			Name: constants.SERVICE_PLAN_FULLY_MANAGED_NAME,
			Fees: []models.ServiceFee{{

				FeePercentage:         15,
				AppliesToAllCostTypes: true,
			},
				{
					FeePercentage:         25,
					AppliesToAllCostTypes: false,
					BookingCostTypeID:     uint(constants.BOOKING_COST_TYPE_WEDDING_FEE_ID),
				},
			},
		},
		{
			Model: gorm.Model{
				ID: uint(constants.SERVICE_PLAN_ENTERPRISE_ID),
			},
			Name: constants.SERVICE_PLAN_ENTERPRISE_NAME,
			Fees: []models.ServiceFee{{

				FeePercentage:         1,
				AppliesToAllCostTypes: true,
			},
			},
		},
	}

	servicePlanRepository := repositories.NewServicePlanRepositoryImplementation(db)

	for _, servicePlan := range servicePlans {
		servicePlanRepository.Create(servicePlan)
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

	SeedUserRoles(database.Instance)
	SeedAmenityTypes(database.Instance)
	SeedBedTypes(database.Instance)
	SeedBookingCostTypes(database.Instance)
	SeedBookingStatus(database.Instance)
	SeedRoomTypes(database.Instance)
	SeedAmenities(database.Instance)
	SeedPaymentMethods(database.Instance)
	SeedInquiryStatuses(database.Instance)
	SeedTaxRates(database.Instance)
	SeedServicePlans(database.Instance)

	log.Println("Database seeding Completed!")

}
