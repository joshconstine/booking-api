package router

import (
	"booking-api/controllers"
	"booking-api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	boatController *controllers.BoatController,
	bookingController *controllers.BookingController,
) *gin.Engine {

	router := gin.Default()
	//allow cors
	router.Use(middlewares.CORSMiddleware())

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		// api.GET("/boats", controllers.GetBoats)
		// api.GET("/boats/:id", controllers.GetBoat)
		// api.GET("/boats/:id/photos", controllers.GetBoatPhotosForBoat)

		bookingRouter := api.Group("/bookings")
		bookingRouter.GET("", bookingController.FindAll)
		bookingRouter.GET("/:bookingId", bookingController.FindById)

		boatRouter := api.Group("/boats")
		boatRouter.GET("", boatController.FindAll)
		boatRouter.GET("/:boatId", boatController.FindById)
		// boatRouter.POST("", boatController.Create)
		// boatRouter.PATCH("/:boatId", boatController.Update)
		// boatRouter.DELETE("/:boatId", boatController.Delete)

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
