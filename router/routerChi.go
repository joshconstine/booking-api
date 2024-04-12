package router

import (
	"booking-api/controllers"
	"booking-api/middlewares"
	"booking-api/view/home"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewChiRouter(rentalsController *controllers.RentalController) *chi.Mux {

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

	return router
}
