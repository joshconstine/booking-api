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
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Booking{})
	Instance.AutoMigrate(&models.BookingDetails{})
	Instance.AutoMigrate(&models.BookingStatus{})
	Instance.AutoMigrate(&models.BookingCostItem{})
	Instance.AutoMigrate(&models.Location{})

	Instance.AutoMigrate(&models.PaymentMethod{})
	Instance.AutoMigrate(&models.BookingCostType{})

	Instance.AutoMigrate(&models.AmenityType{})
	Instance.AutoMigrate(&models.Amenity{})

	Instance.AutoMigrate(&models.BedType{})

	Instance.AutoMigrate(&models.Boat{})
	Instance.AutoMigrate(&models.BoatDefaultSettings{})
	Instance.AutoMigrate(&models.BoatPhoto{})

	Instance.AutoMigrate(&models.Timeblock{})
	Instance.AutoMigrate(&models.Rental{})
	Instance.AutoMigrate(&models.BookingPayment{})
	Instance.AutoMigrate(&models.RentalStatus{})

	log.Println("Database Migration Completed!")
}
