package database

import (
	"booking-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info,
	// 		IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,  // Don't include params in the SQL log
	// 		Colorful:                  false, // Disable color
	// 	},
	// )

	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		// Logger:                                   newLogger,
	})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func Migrate() {
	Instance.Debug()
	// Instance.AutoMigrate((&models.Task{}))
	//**************Helpers**************
	Instance.AutoMigrate(&models.Address{})
	Instance.AutoMigrate(&models.Country{})
	Instance.AutoMigrate(&models.Region{})
	Instance.AutoMigrate(&models.Postal{})
	Instance.AutoMigrate(&models.AddressType{})
	Instance.AutoMigrate(&models.Street{})
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Locality{})
	Instance.AutoMigrate(&models.Location{})
	Instance.AutoMigrate(&models.EntityReview{})
	// Instance.AutoMigrate(&models.BookingStatus{})
	// Instance.AutoMigrate(&models.BookingCostItem{})
	Instance.AutoMigrate(&models.Location{})
	// Instance.AutoMigrate(&models.Photo{})
	// Instance.AutoMigrate(&models.UserRole{})

	// Instance.AutoMigrate(&models.Document{})

	// Instance.AutoMigrate(&models.BookingDocument{})

	// Instance.AutoMigrate(&models.PaymentMethod{})
	// Instance.AutoMigrate(&models.BookingCostType{})
	// Instance.AutoMigrate(&models.Amenity{})
	// Instance.AutoMigrate(&models.TaxRate{})
	// Instance.AutoMigrate(&models.BedType{})
	// Instance.AutoMigrate(&models.RoomType{})

	// Instance.AutoMigrate(&models.Booking{})
	// Instance.AutoMigrate(&models.EntityBooking{})
	// Instance.AutoMigrate(&models.BookingDetails{})
	// Instance.AutoMigrate(&models.BookingPayment{})

	// Instance.AutoMigrate(&models.Boat{})
	// Instance.AutoMigrate(&models.BoatStatus{})

	// Instance.AutoMigrate(&models.Rental{})
	// Instance.AutoMigrate(&models.RentalStatus{})
	// Instance.AutoMigrate(&models.EntityPhoto{})
	// Instance.AutoMigrate(&models.RentalRoom{})

	// Instance.AutoMigrate(&models.EntityBookingCost{})
	// Instance.AutoMigrate(&models.EntityBookingCostAdjustment{})
	// Instance.AutoMigrate(&models.EntityBookingDurationRule{})
	// Instance.AutoMigrate(&models.EntityBookingRule{})
	// Instance.AutoMigrate(&models.EntityBookingDocument{})
	// Instance.AutoMigrate(&models.EntityTimeblock{})

	// Instance.AutoMigrate(&models.Login{})

	//***SAS***
	// Instance.AutoMigrate(&models.Account{})
	// Instance.AutoMigrate(&models.Membership{})
	// Instance.AutoMigrate(&models.AccountSettings{})
	// Instance.AutoMigrate(&models.ServicePlan{})

	// Instance.AutoMigrate(&models.EntityBookingRequest{})
	// Instance.AutoMigrate(&models.InquiryStatus{})
	// Instance.AutoMigrate(&models.ServiceFee{})
	// Instance.AutoMigrate(&models.EntityBookingPermission{})

	// Instance.AutoMigrate(&models.Chat{})
	// Instance.AutoMigrate(&models.ChatMessage{})

	log.Println("Database Migration Completed!")
}
