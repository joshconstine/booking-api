package jobs

import (
	"booking-api/api"
	"database/sql"
	"os"

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

	}

	defer db.Close()

}

func VerifyBookingStatuses() {
	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("0 6,12,18,0 * * *", func() { VerifyAndUpdateBookingStatuses() })

	// Start cron with one scheduled job
	log.Info("Start cron")
	c.Start()
	printCronEntries(c.Entries())

}

func printCronEntries(cronEntries []*cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}
