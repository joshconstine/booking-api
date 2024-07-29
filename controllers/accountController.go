package controllers

import (
	"booking-api/constants"
	"booking-api/repositories"
	"booking-api/view/settings"
	"bytes"
	"encoding/json"
	"github.com/stripe/stripe-go/v78"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78/account"
	"github.com/stripe/stripe-go/v78/accountsession"
	session "github.com/stripe/stripe-go/v78/checkout/session"
)

type AccountController struct {
	AccountRepository repositories.AccountRepository
}

func NewAccountController(accountRepository repositories.AccountRepository) *AccountController {
	return &AccountController{
		AccountRepository: accountRepository,
	}
}

func (controller *AccountController) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	account, err := controller.AccountRepository.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type RequestBody struct {
	Account string `json:"account"`
}

func (ac *AccountController) CreateAccountSession(w http.ResponseWriter, r *http.Request) error {
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
			Payments: &stripe.AccountSessionComponentsPaymentsParams{
				Enabled: stripe.Bool(true),
				Features: &stripe.AccountSessionComponentsPaymentsFeaturesParams{
					RefundManagement:  stripe.Bool(true),
					DisputeManagement: stripe.Bool(true),
					CapturePayments:   stripe.Bool(true),
				},
			},
			Payouts: &stripe.AccountSessionComponentsPayoutsParams{
				Enabled: stripe.Bool(true),
				Features: &stripe.AccountSessionComponentsPayoutsFeaturesParams{
					InstantPayouts:     stripe.Bool(true),
					StandardPayouts:    stripe.Bool(true),
					EditPayoutSchedule: stripe.Bool(true),
					//ExternalAccountCollection: stripe.Bool(true),
				},
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

func (ac *AccountController) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) error {

	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				Price:    stripe.String("{{PRICE_ID}}"),
				Quantity: stripe.Int64(1),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			ApplicationFeeAmount: stripe.Int64(123),
			TransferData: &stripe.CheckoutSessionPaymentIntentDataTransferDataParams{
				Destination: stripe.String("{{CONNECTED_ACCOUNT_ID}}"),
			},
		},
		UIMode:    stripe.String(string(stripe.CheckoutSessionUIModeEmbedded)),
		ReturnURL: stripe.String("https://example.com/checkout/return?session_id={CHECKOUT_SESSION_ID}"),
	}
	accountSession, err := session.New(params)

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

func (ac *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return nil
	}
	accountId := 0

	userId := GetAuthenticatedUser(r).User.UserID
	userAccountRoles, err := ac.AccountRepository.GetUserAccountRoles(userId)
	if err != nil {
		log.Printf("An error occurred when getting user account roles: %v", err)
		handleError(w, err)
		return err
	}

	//TODO:: this is bad, add accountid from a paramater as this will caise bugs for hosts with multiple accounts

	for _, userAccountRole := range userAccountRoles {
		if userAccountRole.Role.ID == constants.USER_ROLE_ACCOUNT_OWNER_ID {
			accountId = int(userAccountRole.AccountID)
			break
		}
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

	err = ac.AccountRepository.AddStripeIDToAccountSettings(uint(accountId), account.ID)

	if err != nil {
		log.Printf("An error occurred when adding the Stripe ID to the account settings: %v", err)
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

func (ac *AccountController) HandleAccountFinance(w http.ResponseWriter, r *http.Request) error {
	user := GetAuthenticatedUser(r)
	memberships, err := ac.AccountRepository.GetUserAccountRoles(user.User.UserID)
	if err != nil {
		return err
	}
	var accountID uint
	for _, userAccountRole := range memberships {
		if userAccountRole.Role.ID == constants.USER_ROLE_ACCOUNT_OWNER_ID {
			accountID = userAccountRole.AccountID
			break
		}

	}
	accountSettings, err := ac.AccountRepository.GetAccountSettings(accountID)
	if err != nil {
		return err

	}

	if accountSettings.StripeAccountID == "" {
		return render(r, w, settings.StripeOnboarding())
	}
	//TODO:: handle multople accounts be rendering some sort of parent compoentn.
	return render(r, w, settings.StripeAccountInfo(accountSettings.StripeAccountID))
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
