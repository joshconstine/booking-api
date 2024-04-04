package router

import (
	"booking-api/controllers"
	"booking-api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	boatController *controllers.BoatController,
	bookingController *controllers.BookingController,
	userController *controllers.UserController,
	bookingStatusController *controllers.BookingStatusController,
	bookingCostTypeController *controllers.BookingCostTypeController,
	rentalController *controllers.RentalController,
	amenityController *controllers.AmenityController,
	bedTypeController *controllers.BedTypeController,
	amenityTypeController *controllers.AmenityTypeController,
	bookingCostItemController *controllers.BookingCostItemController,
	paymentMethodController *controllers.PaymentMethodController,
	bookingPaymentController *controllers.BookingPaymentController,
) *gin.Engine {

	router := gin.Default()
	//allow cors
	router.Use(middlewares.CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		// api.GET("/boats", controllers.GetBoats)
		// api.GET("/boats/:id", controllers.GetBoat)
		// api.GET("/boats/:id/photos", controllers.GetBoatPhotosForBoat)

		//Amenities

		amenityRouter := api.Group("/amenities")
		amenityRouter.GET("", amenityController.FindAll)
		amenityRouter.GET("/:amenityId", amenityController.FindById)
		amenityRouter.POST("", amenityController.Create)

		amenityTypeRouter := api.Group("/amenityTypes")
		amenityTypeRouter.GET("", amenityTypeController.FindAll)
		amenityTypeRouter.GET("/:amenityTypeId", amenityTypeController.FindById)
		amenityTypeRouter.POST("", amenityTypeController.Create)

		bedTypeRouter := api.Group("/bedTypes")
		bedTypeRouter.GET("", bedTypeController.FindAll)
		bedTypeRouter.GET("/:bedTypeId", bedTypeController.FindById)

		paymentMethodRouter := api.Group("/paymentMethods")
		paymentMethodRouter.GET("", paymentMethodController.FindAll)
		paymentMethodRouter.GET("/:paymentMethodId", paymentMethodController.FindById)

		bookingPaymentRouter := api.Group("/bookingPayments")
		bookingPaymentRouter.GET("", bookingPaymentController.FindAll)
		bookingPaymentRouter.GET("/:bookingPaymentId", bookingPaymentController.FindById)
		bookingPaymentRouter.POST("", bookingPaymentController.Create)

		bookingRouter := api.Group("/bookings")
		bookingRouter.GET("", bookingController.FindAll)
		bookingRouter.GET("/:bookingId/details", bookingController.GetDetailsForBookingID)
		bookingRouter.GET("/:bookingId", bookingController.FindById)
		bookingRouter.GET("/:bookingId/costItems", bookingCostItemController.FindByBookingId)
		bookingRouter.GET("/:bookingId/costItems/total", bookingCostItemController.TotalForBookingId)
		bookingRouter.POST("/ui", bookingController.CreateBookingWithUserInformation)

		bookingCostItemRouter := api.Group("/bookingCostItems")
		bookingCostItemRouter.POST("", bookingCostItemController.Create)
		bookingCostItemRouter.PUT("", bookingCostItemController.Update)
		bookingCostItemRouter.DELETE("/:bookingCostItemId", bookingCostItemController.Delete)

		bookingCostTypeRouter := api.Group("/bookingCostTypes")
		bookingCostTypeRouter.GET("", bookingCostTypeController.FindAll)
		bookingCostTypeRouter.GET("/:costTypeId", bookingCostTypeController.FindById)

		bookingStatusRouter := api.Group("/bookingStatus")
		bookingStatusRouter.GET("", bookingStatusController.FindAll)
		bookingStatusRouter.GET("/:statusId", bookingStatusController.FindById)

		boatRouter := api.Group("/boats")
		boatRouter.GET("", boatController.FindAll)
		boatRouter.GET("/:boatId", boatController.FindById)
		// boatRouter.POST("", boatController.Create)
		// boatRouter.PATCH("/:boatId", boatController.Update)
		// boatRouter.DELETE("/:boatId", boatController.Delete)

		userRouter := api.Group("/users")
		userRouter.GET("", userController.FindAll)
		userRouter.POST("/register", userController.RegisterUser)

		rentalRouter := api.Group("/rentals")
		rentalRouter.GET("", rentalController.FindAll)
		rentalRouter.GET("/:rentalId", rentalController.FindById)

		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
			// secured.POST("/boats", controllers.CreateBoat)
			// secured.POST("/boats/:id/photos", controllers.CreateBoatPhoto)
			secured.DELETE("/boatPhoto/:id", controllers.DeleteBoatPhoto)

		}
	}
	return router
}
