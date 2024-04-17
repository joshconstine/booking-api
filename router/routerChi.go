package router

import (
	"booking-api/controllers"
	"booking-api/middlewares"
	"booking-api/services"
	home "booking-api/view/home"
	"os"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewChiRouter(authController *controllers.AuthController, rentalsController *controllers.RentalController, bookingController *controllers.BookingController, boatsController *controllers.BoatController, userSettingsController *controllers.UserSettingsController,
	userService *services.UserService) *chi.Mux {

	router := chi.NewMux()
	// router.Use(middlewares.WithLogger)
	router.Use(middlewares.WithUser)
	withAccountSetupMiddleware := middlewares.NewWithAccountSetupMiddleWare(*userService)

	router.Handle("/*", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		home.Index().Render(r.Context(), w)
	})

	router.Get("/rentals", func(w http.ResponseWriter, r *http.Request) {
		rentalsController.HandleHomeIndex(w, r)
	})

	router.Get("/rentals/{rentalId}", func(w http.ResponseWriter, r *http.Request) {
		rentalsController.HandleRentalDetail(w, r)
	})
	router.Get("/boats", func(w http.ResponseWriter, r *http.Request) {
		boatsController.HandleHomeIndex(w, r)
	})

	router.Get("/boats/{boatId}", func(w http.ResponseWriter, r *http.Request) {
		boatsController.HandleBoatDetail(w, r)
	})

	router.Get("/bookings", func(w http.ResponseWriter, r *http.Request) {
		bookingController.HandleHomeIndex(w, r)
	})

	router.Get("/bookings/{bookingId}", func(w http.ResponseWriter, r *http.Request) {
		bookingController.HandleBookingInformation(w, r)
	})

	router.Get("/login", controllers.Make(authController.HandleLoginIndex))
	router.Get("/login/provider/google", controllers.Make(authController.HandleLoginWithGoogle))
	router.Get("/signup", controllers.Make(authController.HandleSignupIndex))
	router.Post("/logout", controllers.Make(authController.HandleLogoutCreate))
	router.Post("/login", controllers.Make(authController.HandleLoginCreate))
	router.Get("/auth/callback", controllers.Make(authController.HandleAuthCallback))

	router.Group(func(auth chi.Router) {
		auth.Use(middlewares.WithAuth)
		auth.Get("/account/setup", controllers.Make(authController.HandleAccountSetupIndex))
		auth.Post("/account/setup", controllers.Make(authController.HandleAccountSetupCreate))
	})

	router.Group(func(auth chi.Router) {
		auth.Use(middlewares.WithAuth, withAccountSetupMiddleware)
		auth.Get("/settings", controllers.Make(userSettingsController.HandleSettingsIndex))
		// router.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
		// 	userSettingsController.HandleSettingsIndex(w, r)
		// })
		auth.Put("/settings/account/profile", controllers.Make(userSettingsController.HandleSettingsUsernameUpdate))
	})

	return router
}
