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

		bookingRouter := api.Group("/bookings")
		bookingRouter.GET("", bookingController.FindAll)
		bookingRouter.GET("/:bookingId/details", bookingController.GetDetailsForBookingID)
		bookingRouter.GET("/:bookingId", bookingController.FindById)
		bookingRouter.POST("/ui", bookingController.CreateBookingWithUserInformation)

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
