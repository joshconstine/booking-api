package controllers

import (
	"booking-api/data/request"
	validate "booking-api/pkg/kit"
	"booking-api/services"
	"booking-api/view/settings"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/account"
	"github.com/stripe/stripe-go/v78/accountsession"
)

type UserSettingsController struct {
	userService services.UserService
}

func NewUserSettingsController(service services.UserService) *UserSettingsController {
	return &UserSettingsController{userService: service}
}

func (usc *UserSettingsController) HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	// user.User = usc.userService.FindByUserID(user.User.UserID)
	return render(r, w, settings.Index(user, "profile"))
}

func (usc *UserSettingsController) HandleSettingsUpdate(w http.ResponseWriter, r *http.Request) error {
	params := settings.ProfileParams{
		Username:    r.FormValue("username"),
		FirstName:   r.FormValue("firstName"),
		LastName:    r.FormValue("lastName"),
		PhoneNumber: r.FormValue("phoneNumber"),
	}
	errors := settings.ProfileErrors{}
	ok := validate.New(&params, validate.Fields{
		"Username": validate.Rules(validate.Min(3), validate.Max(40)),
	}).Validate(&errors)
	if !ok {
		return render(r, w, settings.ProfileForm(params, errors))
	}
	user := GetAuthenticatedUser(r)
	user.User.Username = params.Username

	updateUserRequest := request.UpdateUserRequest{
		UserID:      user.User.UserID,
		Username:    user.User.Username,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		PhoneNumber: params.PhoneNumber,
	}

	// if err := db.UpdateY(&user.User); err != nil {
	// 	return err
	// }

	err := usc.userService.UpdateUser(&updateUserRequest)
	if err != nil {
		return err
	}

	params.Success = true
	return render(r, w, settings.ProfileForm(params, settings.ProfileErrors{}))
}

func (usc *UserSettingsController) HandleSubscriptions(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Subscription(user))
	// return nil
}

func (usc *UserSettingsController) HandleTeam(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Team(user))
}

func (usc *UserSettingsController) HandleFinances(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Finances(user))
}

func (usc *UserSettingsController) HandleNotifications(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Notifications(user))
}

func (usc *UserSettingsController) HandleCleaners(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Cleaners(user))
}

func (usc *UserSettingsController) HandleSecurity(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Security(user))
}

func (usc *UserSettingsController) HandleProfile(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	return render(r, w, settings.Profile(user))
}

type RequestBody struct {
	Account string `json:"account"`
}

func (usc *UserSettingsController) CreateAccountSession(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	params := &stripe.AccountSessionParams{
		Account: stripe.String(requestBody.Account),
		Components: &stripe.AccountSessionComponentsParams{
			AccountOnboarding: &stripe.AccountSessionComponentsAccountOnboardingParams{
				Enabled: stripe.Bool(true),
			},
		},
	}

	accountSession, err := accountsession.New(params)

	if err != nil {
		log.Printf("An error occurred when calling the Stripe API to create an account session: %v", err)
		handleError(w, err)
		return err
	}

	writeJSON(w, struct {
		ClientSecret string `json:"client_secret"`
	}{
		ClientSecret: accountSession.ClientSecret,
	})
	return nil
}

func (usc *UserSettingsController) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
	}

	account, err := account.New(&stripe.AccountParams{
		Controller: &stripe.AccountControllerParams{
			StripeDashboard: &stripe.AccountControllerStripeDashboardParams{
				Type: stripe.String("none"),
			},
			Fees: &stripe.AccountControllerFeesParams{
				Payer: stripe.String("application"),
			},
		},
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		Country: stripe.String("US"),
	})

	if err != nil {
		log.Printf("An error occurred when calling the Stripe API to create an account: %v", err)
		handleError(w, err)
		return err
	}

	writeJSON(w, struct {
		Account string `json:"account"`
	}{
		Account: account.ID,
	})
	return nil
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if stripeErr, ok := err.(*stripe.Error); ok {
		writeJSON(w, struct {
			Error string `json:"error"`
		}{
			Error: stripeErr.Msg,
		})
	} else {
		writeJSON(w, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
	}
	return
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
