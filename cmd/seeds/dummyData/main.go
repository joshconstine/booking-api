package main

import (
	"booking-api/config"
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/models"
	"booking-api/pkg/database"
	"booking-api/repositories"
	"booking-api/services"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FakeEntity struct {
	EntityType string
	EntityID   uint
}

var entities = []string{
	constants.BOAT_ENTITY,
	constants.RENTAL_ENTITY,
}

func FindRandomUserIDfromDB(db *gorm.DB) string {
	var userIds []string

	result := db.Model(&models.User{}).Limit(10).Pluck("id", &userIds)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	userID := userIds[gofakeit.Number(0, len(userIds)-1)]
	return userID
}

func SelectRandomIndexFromSlice(slice []uint) uint {
	return uint(gofakeit.Number(0, len(slice)-1))
}

func GetRandomIDForEntity(entity string, db *gorm.DB) uint {

	switch entity {
	case constants.BOAT_ENTITY:
		boatRepository := repositories.NewBoatRepositoryImplementation(db)
		boats := boatRepository.FindAllIDs()
		return uint(boats[SelectRandomIndexFromSlice(boats)])
	case constants.RENTAL_ENTITY:
		timeblockRepository := repositories.NewTimeblockRepositoryImplementation(db)
		rentalRepository := repositories.NewRentalRepositoryImplementation(db, timeblockRepository)
		rentals := rentalRepository.FindAllIDs()
		return uint(rentals[SelectRandomIndexFromSlice(rentals)])

	default:
		return 1
	}
}

func GetRandomEntity(db *gorm.DB) FakeEntity {
	entityType := gofakeit.RandomString(entities)
	entityID := GetRandomIDForEntity(entityType, db)
	return FakeEntity{EntityType: entityType, EntityID: entityID}
}

func GenerateRandomAmmountOfEntityBookings(db *gorm.DB) []request.BookEntityRequest {
	var entityBookings []request.BookEntityRequest

	today := time.Now()
	oneHundredTwentyDaysFromNow := today.AddDate(0, 0, 120)

	startDateOfEntityBookings := gofakeit.DateRange(today, oneHundredTwentyDaysFromNow)

	var rangeForEntityBookingStart time.Time
	var rangeForEntityBookingEnd time.Time
	var randomEntity FakeEntity
	var entityBooking request.BookEntityRequest
	for i := 0; i < gofakeit.Number(1, 2); i++ {

		rangeForEntityBookingStart = gofakeit.DateRange(startDateOfEntityBookings, startDateOfEntityBookings.AddDate(0, 0, 2))
		rangeForEntityBookingEnd = gofakeit.DateRange(rangeForEntityBookingStart.AddDate(0, 0, 3), rangeForEntityBookingStart.AddDate(0, 0, 18))
		randomEntity = GetRandomEntity(db)

		entityBooking = request.BookEntityRequest{
			EntityID:   randomEntity.EntityID,
			EntityType: randomEntity.EntityType,
			StartTime:  rangeForEntityBookingStart,
			EndTime:    rangeForEntityBookingEnd,
		}
		entityBookings = append(entityBookings, entityBooking)
	}
	return entityBookings

}

func GetConflictingStartDateForEntity(entityId uint, entityType string, db *gorm.DB) time.Time {
	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(db)
	timeblocks := timeblockRepository.FindByEntity(entityType, entityId)
	if len(timeblocks) == 0 {
		return time.Now()
	} else {
		return timeblocks[0].StartTime.AddDate(0, 0, 1)
	}
}

func GenerateRandomAmmountOfEntityBookingsWithConflicts(db *gorm.DB) []request.BookEntityRequest {
	var entityBookings []request.BookEntityRequest

	//find a date that would cause a conflict
	var startDateOfEntityBookings time.Time

	var rangeForEntityBookingStart time.Time
	var rangeForEntityBookingEnd time.Time
	var randomEntity FakeEntity
	var entityBooking request.BookEntityRequest
	for i := 0; i < gofakeit.Number(1, 5); i++ {
		randomEntity = GetRandomEntity(db)

		startDateOfEntityBookings = GetConflictingStartDateForEntity(randomEntity.EntityID, randomEntity.EntityType, db)

		rangeForEntityBookingStart = gofakeit.DateRange(startDateOfEntityBookings, startDateOfEntityBookings.AddDate(0, 0, 2))
		rangeForEntityBookingEnd = gofakeit.DateRange(rangeForEntityBookingStart.AddDate(0, 0, 3), rangeForEntityBookingStart.AddDate(0, 0, 18))
		entityBooking = request.BookEntityRequest{
			EntityID:   randomEntity.EntityID,
			EntityType: randomEntity.EntityType,
			StartTime:  rangeForEntityBookingStart,
			EndTime:    rangeForEntityBookingEnd,
		}
		entityBookings = append(entityBookings, entityBooking)
	}
	return entityBookings

}
func SeedBoooking(db *gorm.DB) {
	// create booking

	userRepository := repositories.NewUserRepositoryImplementation(db)
	userService := services.NewUserServiceImplementation(userRepository, nil)
	person := gofakeit.Person()

	/***************** Insert a new user *****************/
	var userToBook = request.CreateUserRequest{
		Email:       gofakeit.Email(),
		Username:    gofakeit.Username(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
		UserID:      uuid.New().String(),
	}
	fmt.Println(userToBook)
	if err := userService.CreateUser(&userToBook); err != nil {
		fmt.Println(err.Error())
		return // handle error
	}

	/***************** Insert a new booking *****************/

	bookingToCreate := request.CreateBookingRequest{
		Email:          userToBook.Email,
		FirstName:      person.FirstName,
		LastName:       person.LastName,
		PhoneNumber:    userToBook.PhoneNumber,
		UserID:         userToBook.UserID,
		Guests:         gofakeit.Number(1, 5),
		EntityRequests: GenerateRandomAmmountOfEntityBookings(db),
	}

	bookingRepository := repositories.NewBookingRepositoryImplementation(db)
	bookingService := services.NewBookingServiceImplementation(bookingRepository, nil, userService)

	bid, err := bookingService.Create(&bookingToCreate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("booking created with id: ", bid)
}

func SeedBoookingWithConflicts(db *gorm.DB) {
	// create booking

	person := gofakeit.Person()

	bookingToCreate := request.CreateBookingRequest{
		Email:          gofakeit.Email(),
		FirstName:      person.FirstName,
		LastName:       person.LastName,
		PhoneNumber:    gofakeit.Phone(),
		Guests:         gofakeit.Number(1, 5),
		EntityRequests: GenerateRandomAmmountOfEntityBookingsWithConflicts(db),
	}

	bookingRepository := repositories.NewBookingRepositoryImplementation(db)
	userRepository := repositories.NewUserRepositoryImplementation(db)
	userService := services.NewUserServiceImplementation(userRepository, nil)
	bookingService := services.NewBookingServiceImplementation(bookingRepository, nil, userService)

	bid, err := bookingService.Create(&bookingToCreate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("booking created with id: ", bid)
}

func SeedMultipleBookings(db *gorm.DB, numberOfBookings int) {
	for i := 0; i < numberOfBookings; i++ {
		SeedBoooking(db)
	}
}

func SeedUsers(db *gorm.DB) {
	// create users
	userRepository := repositories.NewUserRepositoryImplementation(db)
	usersToCreate := request.CreateUserRequest{
		Email:       gofakeit.Email(),
		Username:    gofakeit.Username(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
	}

	userRepository.Create(&usersToCreate)
}

func SeedChat(db *gorm.DB) {
	// create chat
	accountID := 9

	//get 10 user ids from db
	userID := FindRandomUserIDfromDB(db)

	chatRepository := repositories.NewChatRepositoryImplementation(db)
	chatRepository.Create(&models.Chat{
		AccountID: uint(accountID),
		UserID:    userID,
		Messages: []models.ChatMessage{
			{
				Message: gofakeit.Sentence(10),
				UserID:  userID,
			},
			{
				Message: gofakeit.Sentence(10),
				UserID:  userID,
			},
		},
	})
}

func InsertNewUser(db *gorm.DB) (request.CreateUserRequest, error) {
	userRepository := repositories.NewUserRepositoryImplementation(db)
	userService := services.NewUserServiceImplementation(userRepository, nil)

	var userToBook = request.CreateUserRequest{
		Email:       gofakeit.Email(),
		Username:    gofakeit.Username(),
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		PhoneNumber: gofakeit.Phone(),
		UserID:      uuid.New().String(),
	}
	if err := userService.CreateUser(&userToBook); err != nil {
		fmt.Println(err.Error())
		return userToBook, err
	}
	return userToBook, nil
}

func GenerateBookingPermission(db *gorm.DB) {
	// create inquiry

	userID := FindRandomUserIDfromDB(db)

	randomEntity := GetRandomEntity(db)

	var permRequest = models.EntityBookingPermission{
		AccountID:  9,
		EntityID:   randomEntity.EntityID,
		EntityType: randomEntity.EntityType,
		UserID:     userID,
		StartTime:  time.Now(),
		EndTime:    time.Now().AddDate(0, 0, 1),
	}

	result := db.Create(&permRequest)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("Inquiry created with id: ", permRequest.ID)

}

func GrantUserRolesOnAccount(db *gorm.DB, userID string, accountID uint) {
	var membership = models.Membership{
		AccountID: accountID,
		UserID:    userID,
		RoleID:    constants.USER_ROLE_ACCOUNT_OWNER_ID,
	}

	result := db.Create(&membership)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("Membership created with id: ", membership.ID)

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
	// objectStorage.CreateSession()
	// SeedUsers(database.Instance)
	// SeedBookingUI(database.Instance)
	// SeedBoooking(database.Instance)
	// SeedChat(database.Instance)
	SeedMultipleBookings(database.Instance, 10)
	// GenerateInquiry(database.Instance)
	// GenerateBookingPermission(database.Instance)
	GrantUserRolesOnAccount(database.Instance, constants.CURRENT_USER_ID, 15)
	GrantUserRolesOnAccount(database.Instance, constants.CURRENT_USER_ID, 16)

	// SeedBoookingWithConflicts(database.Instance)
	log.Println("Database seeding Completed!")

}
