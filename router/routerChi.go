package router

import (
	"booking-api/constants"
	"booking-api/controllers"
	"booking-api/middlewares"
	"booking-api/services"
	home "booking-api/view/home"
	"os"
	"strconv"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewChiRouter(authController *controllers.AuthController, rentalsController *controllers.RentalController, bookingController *controllers.BookingController, boatsController *controllers.BoatController, userSettingsController *controllers.UserSettingsController,
	userService *services.UserService, adminController *controllers.AdminController, chatController *controllers.ChatController, entityBookingPermissionController *controllers.EntityBookingPermissionController, photoController *controllers.PhotoController, accountController *controllers.AccountController) *chi.Mux {

	router := chi.NewMux()

	// router.Use(middlewares.WithLogger)
	userMiddleware := middlewares.NewWithUserMiddleWare(*userService)
	router.Use(userMiddleware)
	withAccountSetupMiddleware := middlewares.NewWithAccountSetupMiddleWare(*userService)
	withIsAdminMiddleware := middlewares.NewWithIsAdminMiddleWare(*userService)

	router.Handle("/*", http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public"))))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		home.Index().Render(r.Context(), w)
	})

	/************************ ADMIN ROUTES ************************/
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

	router.Put("/bookings/{bookingId}/invoice", controllers.Make(bookingController.HandleCreateInvoiceForBooking))

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
		auth.Post("/billing/account", controllers.Make(accountController.CreateAccount))
		auth.Post("/billing/session", controllers.Make(accountController.CreateAccountSession))

	})

	router.Group(func(auth chi.Router) {
		auth.Use(middlewares.WithAuth, withAccountSetupMiddleware)
		auth.Get("/settings", controllers.Make(userSettingsController.HandleSettingsIndex))
		// router.Get("/settings", func(w http.ResponseWriter, r *http.Request) {
		// 	userSettingsController.HandleSettingsIndex(w, r)
		// })
		auth.Put("/settings/account/profile", controllers.Make(userSettingsController.HandleSettingsUpdate))
		auth.Get("/settings/account/subscriptions", controllers.Make(userSettingsController.HandleSubscriptions))
		auth.Get("/settings/account/profile", controllers.Make(userSettingsController.HandleProfile))
		auth.Get("/settings/account/team", controllers.Make(userSettingsController.HandleTeam))
		auth.Get("/settings/account/finances", controllers.Make(userSettingsController.HandleFinances))
		auth.Get("/settings/account/notifications", controllers.Make(userSettingsController.HandleNotifications))
		auth.Get("/settings/account/cleaners", controllers.Make(userSettingsController.HandleCleaners))
		auth.Get("/settings/account/security", controllers.Make(userSettingsController.HandleSecurity))

	})

	router.Group(func(auth chi.Router) {
		auth.Use(middlewares.WithAuth, withIsAdminMiddleware)
		auth.Get("/admin", controllers.Make(adminController.HandleAdminIndex))
		router.Post("/chat/message", controllers.Make(chatController.HandleChatMessageCreate))
		router.Delete("/chat/message", controllers.Make(chatController.HandleChatMessageDelete))
		router.Put("/permission/{entityBookingPermissionID}", controllers.Make(entityBookingPermissionController.Update))
		router.Put("/permission/{entityBookingPermissionID}/approve", controllers.Make(entityBookingPermissionController.HandleApproveBookingPermissionRequest))
	})

	apiRouter := chi.NewRouter()

	apiRouter.Get("/rentals", controllers.Make(rentalsController.FindAll))
	apiRouter.Post("/rentals/{rentalId}/photos", func(w http.ResponseWriter, r *http.Request) {
		rentalID := chi.URLParam(r, "rentalId")

		rentalIDInt, err := strconv.Atoi(rentalID)

		if err != nil {
			http.Error(w, "Invalid rental ID", http.StatusBadRequest)
			return
		}

		photoController.AddPhoto(w, r, constants.RENTAL_ENTITY, rentalIDInt)
		// bookingController.HandleBookingInformation(w, r)
	})

	router.Mount("/api", apiRouter)
	return router
}
