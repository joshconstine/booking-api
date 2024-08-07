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

	log.Println("Connecting to the database...")
	database.Connect(env.DSN)
	log.Println("Database connected.")

	// Uncomment if you have migrations
	// log.Println("Running database migrations...")
	// database.Migrate()
	// log.Println("Database migrations completed.")

	log.Println("Creating object storage session...")
	objectStorage.CreateSession()
	log.Println("Object storage session created.")

	stripe.Key = env.STRIPE_KEY

	log.Println("Creating Supabase client...")
	sb.CreateAuthClient(env)
	log.Println("Supabase client created.")

	// Create a new paypal client

	// pay.CreatePaypalClient()
	// payments.CreatePaypalClient(env)

	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		port := env.PORT
		log.Printf("Application running on port %s", port)
		if err := http.ListenAndServe(port, app); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	return func() {
		log.Println("Running cleanup...")
		cleanup()
		log.Println("Cleanup completed.")
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
	bookingCostTypeRepository := repositories.NewBookingCostTypeRepositoryImplementation(database.Instance)
	timeblockRepository := repositories.NewTimeblockRepositoryImplementation(database.Instance)
	rentalRepository := repositories.NewRentalRepositoryImplementation(database.Instance, timeblockRepository)
	amenityRepository := repositories.NewAmenityRepositoryImplementation(database.Instance)
	// amenityTypeRepository := repositories.NewAmenityTypeRepositoryImplementation(database.Instance)
	bedTypeRepository := repositories.NewBedTypeRepositoryImplementation(database.Instance)
	// bookingCostItemRepository := repositories.NewBookingCostItemRepositoryImplementation(database.Instance)
	// paymentMethodRepository := repositories.NewPaymentMethodRepositoryImplementation(database.Instance)
	bedRepository := repositories.NewBedRepositoryImplementation(database.Instance)

	rentalStatusRepository := repositories.NewRentalStatusRepositoryImplementation(database.Instance)
	photoRepository := repositories.NewPhotoRepositoryImplementation(objectStorage.Client, database.Instance)
	entitiyPhotoRepository := repositories.NewEntityPhotoRepositoryImplementation(database.Instance)
	// locationRepository := repositories.NewLocationRepositoryImplementation(database.Instance)
	roomTypeRepository := repositories.NewRoomTypeRepositoryImplementation(database.Instance)
	rentalRoomRepository := repositories.NewRentalRoomRepositoryImplementation(database.Instance, bedRepository)
	// entityBookingDurationRuleRepository := repositories.NewEntityBookingDurationRuleRepositoryImplementation(database.Instance)
	entityBookingRepository := repositories.NewEntityBookingRepositoryImplementation(database.Instance)
	// userRoleRepository := repositories.NewUserRoleRepositoryImplementation(database.Instance)
	// accountRepository := repositories.NewAccountRepositoryImplementation(database.Instance)
	// inquiryRepository := repositories.NewInquiryRepositoryImplementation(database.Instance)
	// entityBookingDocumentRepository := repositories.NewEntityBookingDocumentRepositoryImplementation(database.Instance)
	// entityBookingRuleRepository := repositories.NewEntityBookingRuleRepositoryImplementation(database.Instance)
	entityBookingCostRepository := repositories.NewEntityBookingCostRepositoryImplementation(database.Instance)
	// entityBookingCostAdjustmentRepository := repositories.NewEntityBookingCostAdjustmentRepositoryImplementation(database.Instance)
	chatRepository := repositories.NewChatRepositoryImplementation(database.Instance)
	accountRepository := repositories.NewAccountRepositoryImplementation(database.Instance, bookingRepository)
	entityBookingPermissionRepository := repositories.NewEntityBookingPermissionRepositoryImplementation(database.Instance)
	entityRepository := repositories.NewEntityRepositoryImplementation(database.Instance)
	membershipRepository := repositories.NewMembershipRepositoryImplementation(database.Instance)
	taxRateRepository := repositories.NewTaxRateRepositoryImplementation(database.Instance)
	bookingCostItemRepository := repositories.NewBookingCostItemRepositoryImplementation(database.Instance)

	bookingPaymentRepository := repositories.NewBookingPaymentRepositoryImplementation(bookingCostItemRepository, database.Instance)
	//Init Service
	userService := services.NewUserServiceImplementation(userRepository, entityRepository, membershipRepository, validate)
	entityBookingPermissionService := services.NewEntityBookingPermissionServiceImplementation(entityBookingPermissionRepository)
	entityBookingService := services.NewEntityBookingServiceImplementation(entityBookingRepository, bookingDetailsRepository)
	bookingService := services.NewBookingServiceImplementation(bookingRepository, bookingDetailsRepository, bookingPaymentRepository, validate, userService, entityBookingService)
	bookingDetailsService := services.NewBookingDetailsServiceImplementation(bookingDetailsRepository, bookingPaymentRepository)
	boatService := services.NewBoatServiceImplementation(boatRepository, validate)
	// bookingStatusService := services.NewBookingStatusService(bookingStatusRepository, validate)
	bookingCostTypeService := services.NewBookingCostTypeServiceImplementation(bookingCostTypeRepository, validate)
	bookingCostItemService := services.NewBookingCostItemServiceImplementation(bookingCostItemRepository, validate)
	rentalService := services.NewRentalServiceImplementation(rentalRepository, validate)
	amenityService := services.NewAmenityServiceImplementation(amenityRepository, validate)
	// amenityTypeService := services.NewAmenityTypeServiceImplementation(amenityTypeRepository, validate)
	bedTypeService := services.NewBedTypeServiceImplementation(bedTypeRepository, validate)
	// bookingCostItemService := services.NewBookingCostItemServiceImplementation(bookingCostItemRepository, validate)
	// paymentMethodService := services.NewPaymentMethodServiceImplementation(paymentMethodRepository, validate)
	bookingPaymentService := services.NewBookingPaymentServiceImplementation(bookingPaymentRepository, bookingDetailsService, bookingCostItemService, validate)
	rentalStatusService := services.NewRentalStatusServiceImplementation(rentalStatusRepository, validate)
	photoService := services.NewPhotoServiceImplementation(photoRepository, validate)
	entityPhotoService := services.NewEntityPhotoServiceImplementation(entitiyPhotoRepository, validate)
	// locationService := services.NewLocationServiceImplementation(locationRepository, validate)
	roomTypeService := services.NewRoomTypeServiceImplementation(roomTypeRepository)
	// rentalRoomService := services.NewRentalRoomServiceImplementation(rentalRoomRepository, validate)
	// entityBookingDurationRuleService := services.NewEntityBookingDurationRuleServiceImplementation(entityBookingDurationRuleRepository)
	// userRoleService := services.NewUserRoleServiceImplementation(userRoleRepository)
	// entityBookingDocumentService := services.NewEntityBookingDocumentServiceImplementation(entityBookingDocumentRepository)
	// entityBookingRuleService := services.NewEntityBookingRuleServiceImplementation(entityBookingRuleRepository)
	entityBookingCostService := services.NewEntityBookingCostServiceImplementation(entityBookingCostRepository)
	// entityBookingCostAdjustmentService := services.NewEntityBookingCostAdjustmentServiceImplementation(entityBookingCostAdjustmentRepository)
	chatService := services.NewChatServiceImplementation(chatRepository)
	accountService := services.NewAccountServiceImplementation(accountRepository)
	invoiceService := services.NewInvoiceServiceImplementation(bookingRepository)
	rentalRoomService := services.NewRentalRoomServiceImplementation(rentalRoomRepository, validate)

	//Init controller
	authController := controllers.NewAuthController(userService, sb.ClientInstance)
	bookingController := controllers.NewBookingController(bookingService, bookingDetailsService, invoiceService)
	boatController := controllers.NewBoatController(boatService)
	userController := controllers.NewUserController(userService)
	// bookingStatusController := controllers.NewBookingStatusController(bookingStatusService)
	// bookingCostTypeController := controllers.NewBookingCostTypeController(bookingCostTypeService)
	rentalController := controllers.NewRentalController(rentalService, amenityService, roomTypeService, bedTypeService, rentalRoomService, photoService, entityPhotoService)
	// amenityController := controllers.NewAmenityController(amenityService)
	// amenityTypeController := controllers.NewAmenityTypeController(amenityTypeService)
	// bedTypeController := controllers.NewBedTypeController(bedTypeService)
	// bookingCostItemController := controllers.NewBookingCostItemController(bookingCostItemService)
	// paymentMethodController := controllers.NewPaymentMethodController(paymentMethodService)
	// bookingPaymentController := controllers.NewBookingPaymentController(bookingPaymentService)
	rentalStatusController := controllers.NewRentalStatusController(rentalStatusService)
	photoController := controllers.NewPhotoController(photoService, entityPhotoService)
	// locationController := controllers.NewLocationController(locationService)
	// roomTypeController := controllers.NewRoomTypeController(roomTypeService)
	rentalRoomController := controllers.NewRentalRoomController(rentalRoomService, roomTypeService, bedTypeService, entityPhotoService)
	// entityBookingDurationRuleController := controllers.NewEntityBookingDurationRuleController(entityBookingDurationRuleService)
	entityBookingController := controllers.NewEntityBookingController(entityBookingService, rentalService, boatService)
	// userRoleController := controllers.NewUserRoleController(userRoleService)
	accountController := controllers.NewAccountController(bookingCostItemService, bookingPaymentService, bookingService, accountRepository)
	// inquiryController := controllers.NewInquiryController(inquiryRepository)
	// entityBookingDocumentController := controllers.NewEntityBookingDocumentController(entityBookingDocumentService)
	// entityBookingRuleController := controllers.NewEntityBookingRuleController(entityBookingRuleService)
	entityBookingCostController := controllers.NewEntityBookingCostController(entityBookingCostService, bookingCostTypeService, taxRateRepository)
	// entityBookingCostAdjustmentController := controllers.NewEntityBookingCostAdjustmentController(entityBookingCostAdjustmentService)
	userSettingsController := controllers.NewUserSettingsController(userService, accountService)
	adminController := controllers.NewAdminController(userService, bookingService, accountService)
	entityBookingPermissionController := controllers.NewEntityBookingPermissionController(entityBookingPermissionService)
	comboController := controllers.NewComboController(boatService, rentalService)
	//Router
	// router := router.NewRouter(boatController, bookingController, userController,
	// 	bookingStatusController, bookingCostTypeController, rentalController, amenityController, bedTypeController, amenityTypeController, bookingCostItemController, paymentMethodController, bookingPaymentController, rentalStatusController, photoController, locationController, rentalRoomController, roomTypeController, entityBookingDurationRuleController, entityBookingController, userRoleController, accountController, inquiryController, entityBookingDocumentController, entityBookingRuleController, entityBookingCostController, entityBookingCostAdjustmentController)
	// router.StaticFS("/public", http.Dir("public"))

	chatController := controllers.NewChatController(chatService, userService, accountService)

	chirouter := router.NewChiRouter(authController, rentalController, bookingController, boatController, userSettingsController, &userService, adminController, chatController, entityBookingPermissionController, photoController, accountController, userController, entityBookingController, membershipRepository, entityRepository, entityBookingCostController, rentalStatusController, rentalRoomController, comboController)

	// ginRouter := router.InitRouter(routes)

	return chirouter, func() {

	}, nil
}
