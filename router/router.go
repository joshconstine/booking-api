package router

import (
	"booking-api/constants"
	"booking-api/controllers"
	"booking-api/middlewares"
	"strconv"

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
	rentalStatusController *controllers.RentalStatusController,
	photoController *controllers.PhotoController,
) *gin.Engine {

	router := gin.Default()
	//allow cors
	router.Use(middlewares.CORSMiddleware())

	api := router.Group("/api")
	{
		/************************ AUTH ************************/
		api.POST("/token", controllers.GenerateToken)
		// api.GET("/boats", controllers.GetBoats)
		// api.GET("/boats/:id", controllers.GetBoat)
		// api.GET("/boats/:id/photos", controllers.GetBoatPhotosForBoat)

		/************************ HELPERS ************************/
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

		photoRouter := api.Group("/photos")
		photoRouter.GET("", photoController.FindAll)
		/************************ BOOKINGS ************************/
		bookingRouter := api.Group("/bookings")
		bookingRouter.GET("", bookingController.FindAll)
		bookingRouter.GET("/:bookingId/details", bookingController.GetDetailsForBookingID)
		bookingRouter.GET("/:bookingId", bookingController.FindById)
		bookingRouter.GET("/:bookingId/costItems", bookingCostItemController.FindByBookingId)
		bookingRouter.GET("/:bookingId/costItems/total", bookingCostItemController.TotalForBookingId)
		bookingRouter.POST("/ui", bookingController.CreateBookingWithUserInformation)

		bookingRouter.GET("/:bookingId/payments", bookingPaymentController.FindByBookingId)
		bookingRouter.GET("/:bookingId/payments/total", bookingPaymentController.FindTotalAmountByBookingId)

		bookingCostItemRouter := api.Group("/bookingCostItems")
		bookingCostItemRouter.POST("", bookingCostItemController.Create)
		bookingCostItemRouter.PUT("", bookingCostItemController.Update)
		bookingCostItemRouter.DELETE("/:bookingCostItemId", bookingCostItemController.Delete)

		bookingCostTypeRouter := api.Group("/bookingCostTypes")
		bookingCostTypeRouter.GET("", bookingCostTypeController.FindAll)
		bookingCostTypeRouter.GET("/:costTypeId", bookingCostTypeController.FindById)

		bookingPaymentRouter := api.Group("/bookingPayments")
		bookingPaymentRouter.GET("", bookingPaymentController.FindAll)
		bookingPaymentRouter.GET("/:bookingPaymentId", bookingPaymentController.FindById)
		bookingPaymentRouter.POST("", bookingPaymentController.Create)

		bookingStatusRouter := api.Group("/bookingStatus")
		bookingStatusRouter.GET("", bookingStatusController.FindAll)
		bookingStatusRouter.GET("/:statusId", bookingStatusController.FindById)

		/************************ BOATS ************************/
		boatRouter := api.Group("/boats")
		boatRouter.GET("", boatController.FindAll)
		boatRouter.GET("/:boatId", boatController.FindById)
		// boatRouter.POST("", boatController.Create)
		// boatRouter.PATCH("/:boatId", boatController.Update)
		// boatRouter.DELETE("/:boatId", boatController.Delete)

		userRouter := api.Group("/users")
		userRouter.GET("", userController.FindAll)
		userRouter.POST("/register", userController.RegisterUser)

		/************************ RENTALS ************************/
		rentalRouter := api.Group("/rentals")
		rentalRouter.GET("", rentalController.FindAll)
		rentalRouter.GET("/:rentalId", rentalController.FindById)
		rentalRouter.GET("/:rentalId/status", rentalStatusController.FindByRentalId)

		rentalRouter.POST("/:rentalId/photos", func(ctx *gin.Context) {

			rentalIdint, _ := strconv.Atoi(ctx.Param("rentalId"))

			photoController.AddPhoto(ctx, constants.REANTAL_ENTITY, rentalIdint)
		})

		rentalStatusRouter := api.Group("/rentalStatus")
		rentalStatusRouter.GET("", rentalStatusController.FindAll)
		rentalStatusRouter.PUT("", rentalStatusController.UpdateStatusForRentalId)

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
