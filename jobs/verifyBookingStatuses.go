package jobs

import (
	"booking-api/api"
	"database/sql"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func VerifyAndUpdateBookingStatuses() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	bookingIDs, err := api.GetAllBookingIDsThatAreNotCancelledOrCompleted(db)

	if err != nil {
		log.Fatalf("failed to query booking IDs: %v", err)
	}
	for _, bookingID := range bookingIDs {
		updated, err := api.AuditBookingStatus(bookingID, db)
		if err != nil {
			log.Fatalf("failed to verify booking payment status: %v", err)
		}
		if updated {
			log.Infof("Updated booking status for booking ID: %d", bookingID)
		}

		bookingIDString := strconv.Itoa(bookingID)
		rentalBookings, err := api.GetRentalBookingsForBookingId(bookingIDString, db)
		if err != nil {
			log.Fatalf("failed to query rental bookings: %v", err)
		}
		for _, rentalBooking := range rentalBookings {
			updated, err := api.AuditRentalBookingStatus(rentalBooking.ID, db)
			if err != nil {
				log.Fatalf("failed to verify rental booking payment status: %v", err)
			}
			if updated {
				log.Infof("Updated rental booking status for rental booking ID: %d", rentalBooking.ID)
			}
		}

	}

	defer db.Close()

}

func VerifyBookingStatuses() {
	log.Info("Create new cron")
	c := cron.New()
	// c.AddFunc("0 */1 * * *", func() { VerifyAndUpdateBookingStatuses() })

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())

}

func printCronEntries(cronEntries []*cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}
