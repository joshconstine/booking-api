package main

import (
	"booking-api/config"
	"booking-api/controllers"
	"booking-api/pkg/database"
	"booking-api/pkg/objectStorage"
	"booking-api/pkg/sb"
	"booking-api/repositories"
	"booking-api/router"
	"booking-api/services"
	"fmt"
	"log"
	"net/http"
	"os"

	"booking-api/pkg/shutdown"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stripe/stripe-go/v78"
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
	// jobs.VerifyBookingStatuses() // Call the function to run the jobs

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
	//database.Migrate()

	objectStorage.CreateSession()

	stripe.Key = env.STRIPE_KEY
	// Creates a new Supabase client
	// accessable via sb.ClientInstance
	sb.CreateAuthClient(env)

	// Create a new paypal client

	// pay.CreatePaypalClient()
	// payments.CreatePaypalClient(env)

	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		port := env.PORT
		log.Println("application running", "port", port)
		http.ListenAndServe(port, app)
	}()

	return func() {
		cleanup()
		// app.Shutdown(nil)
	}, nil

}

// func buildServer(env config.EnvVars) (*gin.Engine, func(), error) {
func buildServer(env config.EnvVars) (*chi.Mux, func(), error) {

	validate := validator.New()

	//******************************* Init Repositories ************************************
	//Boats

	//Rentals

	//Entities

	//Bookings

	//Users

	//SAS
	bookingRepository := repositories.NewBookingRepositoryImplementation(database.Instance)
	boatRepository := repositories.NewBoatRepositoryImplementation(database.Instance)
	userRepository := repositories.NewUserRepositoryImplementation(database.Instance)
	bookingDetailsRepository := repositories.NewBookingDetailsRepositoryImplementation(database.Instance)
	// bookingStatusRepository := repositories.NewBookingStatusRepositoryImplementation(database.Instance)
	// bookingCostTypeRepository := repositories.NewBookingCostTypeRepositoryImplementation(database.Instance)
	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(database.Instance)
	rentalRepository := repositories.NewRentalRepositoryImplementation(database.Instance, timeblockRepository)
	// amenityRepository := repositories.NewAmenityRepositoryImplementation(database.Instance)
	// amenityTypeRepository := repositories.NewAmenityTypeRepositoryImplementation(database.Instance)
	// bedTypeRepository := repositories.NewBedTypeRepositoryImplementation(database.Instance)
	// bookingCostItemRepository := repositories.NewBookingCostItemRepositoryImplementation(database.Instance)
	// paymentMethodRepository := repositories.NewPaymentMethodRepositoryImplementation(database.Instance)
	// bookingPaymentRepository := repositories.NewBookingPaymentRepositoryImplementation(database.Instance)
	// rentalStatusRepository := repositories.NewRentalStatusRepositoryImplementation(database.Instance)
	photoRepository := repositories.NewPhotoRepositoryImplementation(objectStorage.Client, database.Instance)
	entitiyPhotoRepository := repositories.NewEntityPhotoRepositoryImplementation(database.Instance)
	// locationRepository := repositories.NewLocationRepositoryImplementation(database.Instance)
	// roomTypeRepository := repositories.NewRoomTypeRepositoryImplementation(database.Instance)
	// rentalRoomRepository := repositories.NewRentalRoomRepositoryImplementation(database.Instance)
	// entityBookingDurationRuleRepository := repositories.NewEntityBookingDurationRuleRepositoryImplementation(database.Instance)
	// entityBookingRepository := repositories.NewEntityBookingRepositoryImplementation(database.Instance)
	// userRoleRepository := repositories.NewUserRoleRepositoryImplementation(database.Instance)
	// accountRepository := repositories.NewAccountRepositoryImplementation(database.Instance)
	// inquiryRepository := repositories.NewInquiryRepositoryImplementation(database.Instance)
	// entityBookingDocumentRepository := repositories.NewEntityBookingDocumentRepositoryImplementation(database.Instance)
	// entityBookingRuleRepository := repositories.NewEntityBookingRuleRepositoryImplementation(database.Instance)
	// entityBookingCostRepository := repositories.NewEntityBookingCostRepositoryImplementation(database.Instance)
	// entityBookingCostAdjustmentRepository := repositories.NewEntityBookingCostAdjustmentRepositoryImplementation(database.Instance)
	chatRepository := repositories.NewChatRepositoryImplementation(database.Instance)
	accountRepository := repositories.NewAccountRepositoryImplementation(database.Instance)
	entityBookingPermissionRepository := repositories.NewEntityBookingPermissionRepositoryImplementation(database.Instance)

	//Init Service
	userService := services.NewUserServiceImplementation(userRepository, validate)
	bookingDetailsService := services.NewBookingDetailsServiceImplementation(bookingDetailsRepository)
	bookingService := services.NewBookingServiceImplementation(bookingRepository, validate, userService)
	boatService := services.NewBoatServiceImplementation(boatRepository, validate)
	// bookingStatusService := services.NewBookingStatusService(bookingStatusRepository, validate)
	// bookingCostTypeService := services.NewBookingCostTypeServiceImplementation(bookingCostTypeRepository, validate)
	rentalService := services.NewRentalServiceImplementation(rentalRepository, validate)
	// amenityService := services.NewAmenityServiceImplementation(amenityRepository, validate)
	// amenityTypeService := services.NewAmenityTypeServiceImplementation(amenityTypeRepository, validate)
	// bedTypeService := services.NewBedTypeServiceImplementation(bedTypeRepository, validate)
	// bookingCostItemService := services.NewBookingCostItemServiceImplementation(bookingCostItemRepository, validate)
	// paymentMethodService := services.NewPaymentMethodServiceImplementation(paymentMethodRepository, validate)
	// bookingPaymentService := services.NewBookingPaymentServiceImplementation(bookingPaymentRepository, validate)
	// rentalStatusService := services.NewRentalStatusServiceImplementation(rentalStatusRepository, validate)
	photoService := services.NewPhotoServiceImplementation(photoRepository, validate)
	entityPhotoService := services.NewEntityPhotoServiceImplementation(entitiyPhotoRepository, validate)
	// locationService := services.NewLocationServiceImplementation(locationRepository, validate)
	// roomTypeService := services.NewRoomTypeServiceImplementation(roomTypeRepository)
	// rentalRoomService := services.NewRentalRoomServiceImplementation(rentalRoomRepository, validate)
	// entityBookingDurationRuleService := services.NewEntityBookingDurationRuleServiceImplementation(entityBookingDurationRuleRepository)
	// entityBookingService := services.NewEntityBookingServiceImplementation(entityBookingRepository)
	// userRoleService := services.NewUserRoleServiceImplementation(userRoleRepository)
	// entityBookingDocumentService := services.NewEntityBookingDocumentServiceImplementation(entityBookingDocumentRepository)
	// entityBookingRuleService := services.NewEntityBookingRuleServiceImplementation(entityBookingRuleRepository)
	// entityBookingCostService := services.NewEntityBookingCostServiceImplementation(entityBookingCostRepository)
	// entityBookingCostAdjustmentService := services.NewEntityBookingCostAdjustmentServiceImplementation(entityBookingCostAdjustmentRepository)
	chatService := services.NewChatServiceImplementation(chatRepository)
	accountService := services.NewAccountServiceImplementation(accountRepository)
	entityBookingPermissionService := services.NewEntityBookingPermissionServiceImplementation(entityBookingPermissionRepository)
	invoiceService := services.NewInvoiceServiceImplementation(bookingRepository)

	//Init controller
	authController := controllers.NewAuthController(userService, sb.ClientInstance)
	bookingController := controllers.NewBookingController(bookingService, bookingDetailsService, invoiceService)
	boatController := controllers.NewBoatController(boatService)
	// userController := controllers.NewUserController(userService)
	// bookingStatusController := controllers.NewBookingStatusController(bookingStatusService)
	// bookingCostTypeController := controllers.NewBookingCostTypeController(bookingCostTypeService)
	rentalController := controllers.NewRentalController(rentalService)
	// amenityController := controllers.NewAmenityController(amenityService)
	// amenityTypeController := controllers.NewAmenityTypeController(amenityTypeService)
	// bedTypeController := controllers.NewBedTypeController(bedTypeService)
	// bookingCostItemController := controllers.NewBookingCostItemController(bookingCostItemService)
	// paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	// bookingPaymentController := controllers.NewBookingPaymentController(bookingPaymentService)
	// rentalStatusController := controllers.NewRentalStatusController(rentalStatusService)
	photoController := controllers.NewPhotoController(photoService, entityPhotoService)
	// locationController := controllers.NewLocationController(locationService)
	// roomTypeController := controllers.NewRoomTypeController(roomTypeService)
	// rentalRoomController := controllers.NewRentalRoomController(rentalRoomService)
	// entityBookingDurationRuleController := controllers.NewEntityBookingDurationRuleController(entityBookingDurationRuleService)
	// entityBookingController := controllers.NewEntityBookingController(entityBookingService)
	// userRoleController := controllers.NewUserRoleController(userRoleService)
	accountController := controllers.NewAccountController(accountRepository)
	// inquiryController := controllers.NewInquiryController(inquiryRepository)
	// entityBookingDocumentController := controllers.NewEntityBookingDocumentController(entityBookingDocumentService)
	// entityBookingRuleController := controllers.NewEntityBookingRuleController(entityBookingRuleService)
	// entityBookingCostController := controllers.NewEntityBookingCostController(entityBookingCostService)
	// entityBookingCostAdjustmentController := controllers.NewEntityBookingCostAdjustmentController(entityBookingCostAdjustmentService)
	userSettingsController := controllers.NewUserSettingsController(userService)
	adminController := controllers.NewAdminController(userService, bookingService, accountService)
	entityBookingPermissionController := controllers.NewEntityBookingPermissionController(entityBookingPermissionService)
	//Router
	// router := router.NewRouter(boatController, bookingController, userController,
	// 	bookingStatusController, bookingCostTypeController, rentalController, amenityController, bedTypeController, amenityTypeController, bookingCostItemController, paymentMethodController, bookingPaymentController, rentalStatusController, photoController, locationController, rentalRoomController, roomTypeController, entityBookingDurationRuleController, entityBookingController, userRoleController, accountController, inquiryController, entityBookingDocumentController, entityBookingRuleController, entityBookingCostController, entityBookingCostAdjustmentController)
	// router.StaticFS("/public", http.Dir("public"))

	chatController := controllers.NewChatController(chatService, userService, accountService)

	router := router.NewChiRouter(authController, rentalController, bookingController, boatController, userSettingsController, &userService, adminController, chatController, entityBookingPermissionController, photoController, accountController)

	// ginRouter := router.InitRouter(routes)

	return router, func() {

	}, nil
}
