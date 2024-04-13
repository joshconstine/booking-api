package router

import (
	"booking-api/controllers"
	"booking-api/middlewares"

	"booking-api/view/home"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewChiRouter(authController *controllers.AuthController, rentalsController *controllers.RentalController, bookingController *controllers.BookingController) *chi.Mux {

	router := chi.NewMux()
	router.Use(middlewares.WithUser)

	router.Handle("/*", http.FileServer(http.Dir("public")))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		home.Index().Render(r.Context(), w)
	})

	router.Get("/rentals", func(w http.ResponseWriter, r *http.Request) {
		rentalsController.HandleHomeIndex(w, r)
	})

	router.Get("/rentals/{rentalId}", func(w http.ResponseWriter, r *http.Request) {
		rentalsController.HandleRentalDetail(w, r)
	})

	router.Get("/bookings", func(w http.ResponseWriter, r *http.Request) {
		bookingController.HandleHomeIndex(w, r)
	})

	router.Get("/bookings/{bookingId}", func(w http.ResponseWriter, r *http.Request) {
		bookingController.HandleBookingInformation(w, r)
	})

	router.Get("/login", controllers.Make(authController.HandleLoginIndex))
	// router.Get("/login/provider/google", handler.Make(handler.HandleLoginWithGoogle))
	// router.Get("/signup", handler.Make(handler.HandleSignupIndex))
	// router.Post("/logout", handler.Make(handler.HandleLogoutCreate))
	// router.Post("/login", handler.Make(handler.HandleLoginCreate))
	// router.Get("/auth/callback", handler.Make(handler.HandleAuthCallback))

	return router
}
