package main

import (
	"booking-api/config"
	"booking-api/controllers"
	"booking-api/database"
	"booking-api/jobs"
	"booking-api/objectStorage"
	"booking-api/repositories"
	"booking-api/router"
	"booking-api/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"booking-api/pkg/shutdown"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// setup exit code for graceful shutdown
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
	//test

	// run the server
	cleanup, err := run(env)

	// Run jobs
	jobs.VerifyBookingStatuses() // Call the function to run the jobs
	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()

}

func run(env config.EnvVars) (func(), error) {

	database.Connect(env.DSN)
	database.Migrate()
	objectStorage.CreateSession()
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.Run()
		log.Println("server started")
	}()

	return func() {
		cleanup()
		// app.Shutdown(nil)
	}, nil
}

func buildServer(env config.EnvVars) (*gin.Engine, func(), error) {

	validate := validator.New()

	//Init Repositories
	bookingRepository := repositories.NewBookingRepositoryImplementation(database.Instance)
	boatRepository := repositories.NewBoatRepositoryImplementation(database.Instance)
	userRepository := repositories.NewUserRepositoryImplementation(database.Instance)
	bookingDetailsRepository := repositories.NewBookingDetailsRepositoryImplementation(database.Instance)
	bookingStatusRepository := repositories.NewBookingStatusRepositoryImplementation(database.Instance)
	bookingCostTypeRepository := repositories.NewBookingCostTypeRepositoryImplementation(database.Instance)
	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(database.Instance)
	rentalRepository := repositories.NewRentalRepositoryImplementation(database.Instance, timeblockRepository)
	amenityRepository := repositories.NewAmenityRepositoryImplementation(database.Instance)
	amenityTypeRepository := repositories.NewAmenityTypeRepositoryImplementation(database.Instance)
	bedTypeRepository := repositories.NewBedTypeRepositoryImplementation(database.Instance)
	bookingCostItemRepository := repositories.NewBookingCostItemRepositoryImplementation(database.Instance)
	paymentMethodRepository := repositories.NewPaymentMethodRepositoryImplementation(database.Instance)
	bookingPaymentRepository := repositories.NewBookingPaymentRepositoryImplementation(database.Instance)
	rentalStatusRepository := repositories.NewRentalStatusRepositoryImplementation(database.Instance)
	photoRepository := repositories.NewPhotoRepositoryImplementation(objectStorage.Client, database.Instance)
	entitiyPhotoRepository := repositories.NewEntityPhotoRepositoryImplementation(database.Instance)

	//Init Service
	userService := services.NewUserServiceImplementation(userRepository, validate)
	bookingDetailsService := services.NewBookingDetailsServiceImplementation(bookingDetailsRepository)
	bookingService := services.NewBookingServiceImplementation(bookingRepository, validate, userService, bookingDetailsService)
	boatService := services.NewBoatServiceImplementation(boatRepository, validate)
	bookingStatusService := services.NewBookingStatusService(bookingStatusRepository, validate)
	bookingCostTypeService := services.NewBookingCostTypeServiceImplementation(bookingCostTypeRepository, validate)
	rentalService := services.NewRentalServiceImplementation(rentalRepository, validate)
	amenityService := services.NewAmenityServiceImplementation(amenityRepository, validate)
	amenityTypeService := services.NewAmenityTypeServiceImplementation(amenityTypeRepository, validate)
	bedTypeService := services.NewBedTypeServiceImplementation(bedTypeRepository, validate)
	bookingCostItemService := services.NewBookingCostItemServiceImplementation(bookingCostItemRepository, validate)
	paymentMethodService := services.NewPaymentMethodServiceImplementation(paymentMethodRepository, validate)
	bookingPaymentService := services.NewBookingPaymentServiceImplementation(bookingPaymentRepository, validate)
	rentalStatusService := services.NewRentalStatusServiceImplementation(rentalStatusRepository, validate)
	photoService := services.NewPhotoServiceImplementation(photoRepository, validate)
	entityPhotoService := services.NewEntityPhotoServiceImplementation(entitiyPhotoRepository, validate)

	//Init controller
	bookingController := controllers.NewBookingController(bookingService, bookingDetailsService)
	boatController := controllers.NewBoatController(boatService)
	userController := controllers.NewUserController(userService)
	bookingStatusController := controllers.NewBookingStatusController(bookingStatusService)
	bookingCostTypeController := controllers.NewBookingCostTypeController(bookingCostTypeService)
	rentalController := controllers.NewRentalController(rentalService)
	amenityController := controllers.NewAmenityController(amenityService)
	amenityTypeController := controllers.NewAmenityTypeController(amenityTypeService)
	bedTypeController := controllers.NewBedTypeController(bedTypeService)
	bookingCostItemController := controllers.NewBookingCostItemController(bookingCostItemService)
	paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	bookingPaymentController := controllers.NewBookingPaymentController(bookingPaymentService)
	rentalStatusController := controllers.NewRentalStatusController(rentalStatusService)
	photoController := controllers.NewPhotoController(photoService, entityPhotoService)

	//Router
	router := router.NewRouter(boatController, bookingController, userController,
		bookingStatusController, bookingCostTypeController, rentalController, amenityController, bedTypeController, amenityTypeController, bookingCostItemController, paymentMethodController, bookingPaymentController, rentalStatusController, photoController)

	// ginRouter := router.InitRouter(routes)

	return router, func() {

	}, nil
}
