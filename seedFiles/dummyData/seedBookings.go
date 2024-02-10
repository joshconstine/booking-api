package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type RentalBookingCreate struct {
	RentalID  int
	StartTime time.Time
	EndTime   time.Time
	BookingID int
}

type BoatBookingCreate struct {
	BoatID     int
	StartTime  time.Time
	EndTime    time.Time
	BookingID  int
	LocationID int
}

// get random time between 1 and 100 days
func randomTime() time.Time {
	rand.Seed(time.Now().UnixNano())
	min := time.Now()
	max := min.AddDate(0, 0, 100)
	return min.Add(time.Duration(rand.Int63n(max.Unix() - min.Unix())))

}

func RandomDateRangeBetweenNowAnd180Days() (time.Time, time.Time) {
	rand.Seed(time.Now().UnixNano())
	//get a random number of days between 1 and 180
	days := rand.Intn(180) + 1
	min := time.Now()
	min = min.AddDate(0, 0, days)
	max := min.AddDate(0, 0, rand.Intn(7)+1)
	return min, max
}

func GenerateRandomRentalBooking(bookingID int) RentalBookingCreate {

	start, end := RandomDateRangeBetweenNowAnd180Days()
	return RentalBookingCreate{
		RentalID:  rand.Intn(11),
		StartTime: start,
		EndTime:   end,

		BookingID: bookingID,
	}
}

func GenerateRandomBoatBooking(bookingID int) BoatBookingCreate {

	start, end := RandomDateRangeBetweenNowAnd180Days()
	return BoatBookingCreate{
		BoatID:     rand.Intn(4),
		StartTime:  start,
		EndTime:    end,
		BookingID:  bookingID,
		LocationID: rand.Intn(2),
	}
}

func main() {
	// Mock API endpoint
	bookingCreateApiEndpoint := "http://localhost:8080/bookings/ui"
	rentalBookingCreateApiEndpoint := "http://localhost:8080/rentalBooking"
	boatBookingCreateApiEndpoint := "http://localhost:8080/boatBooking"

	// Seed some bookings
	numBookings := 10 // Change as needed
	bookingCreateInformation := make([]User, numBookings)

	for i := 0; i < numBookings; i++ {
		bookingCreateInformation[i] = User{
			FirstName:   faker.FirstName(),
			LastName:    faker.LastName(),
			Email:       faker.Email(),
			PhoneNumber: faker.Phonenumber(),
		}

	}

	createdBookingIDs := []int{}

	//loop through the bookings and send them to the API
	for _, booking := range bookingCreateInformation {
		bookingJSON, err := json.Marshal(booking)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post(bookingCreateApiEndpoint, "application/json", bytes.NewBuffer(bookingJSON))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Status)

		//parse the response to get the booking ID
		var bookingID int
		json.NewDecoder(resp.Body).Decode(&bookingID)
		createdBookingIDs = append(createdBookingIDs, bookingID)

	}

	fmt.Println("Bookings seeded successfully")
	fmt.Println(createdBookingIDs)

	// Seed rental Bookings
	for _, bookingID := range createdBookingIDs {
		booking := GenerateRandomRentalBooking(bookingID)
		bookingJSON, err := json.Marshal(booking)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := http.Post(rentalBookingCreateApiEndpoint, "application/json", bytes.NewBuffer(bookingJSON))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Status)

		boatBooking := GenerateRandomBoatBooking(bookingID)
		boatBookingJSON, err := json.Marshal(boatBooking)
		if err != nil {
			log.Fatal(err)
		}
		resp, err = http.Post(boatBookingCreateApiEndpoint, "application/json", bytes.NewBuffer(boatBookingJSON))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Status)

	}

	fmt.Println("Rental Bookings seeded successfully")

}
