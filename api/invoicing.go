package api

import (
	"booking-api/api/payments"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/plutov/paypal/v4"
)

func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//invoice is a struct that holds the invoice details
	//decode the request body into struct and check for errors

	// Create a paypal client
	client := payments.CreatePaypalClient()

	// Create invoice
	createdInvoice, err := payments.CreateInvoice(r.Context(), client, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create invoice: %v", err), http.StatusInternalServerError)

		return
	}

	// //log the invoice number

	//log the invoice number
	// fmt.Printf("created invoice number: %s", createdInvoice)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdInvoice)
}

func AddMerchantDetailsToInvoice(invoice *payments.CreateInvoiceRequest, db *sql.DB) error {
	//get merchant details from db
	//fill in the invoicer field
	invoice.Detail.Note = "Thank you for your business"
	return nil
}

func SetConstantsForInvoice(invoice *payments.CreateInvoiceRequest) {
	//fill in the configuration field
	//fill in the amount field
	invoice.Detail.CurrencyCode = "USD"
	invoice.Detail.PaymentTerm.TermType = "no refunds with 30 days of booking"

	invoice.Configuration.PartialPayment.AllowPartialPayment = true
	invoice.Configuration.PartialPayment.MinimumAmountDue = paypal.Money{
		Currency: "USD",
		Value:    "50.00",
	}

	invoice.Configuration.AllowTip = true
	invoice.Configuration.TaxCalculatedAfterDiscount = true
	invoice.Configuration.TaxInclusive = false
	invoice.Configuration.TemplateId = ""
}

func AddInvoicerDetailsToInvoice(invoice *payments.CreateInvoiceRequest, db *sql.DB) error {
	//get invoicer details from db
	//fill in the invoicer field
	invoice.Invoicer.EmailAddress = "joshua.constine97@gmail.com"
	invoice.Invoicer.Phones = []paypal.InvoicerPhoneDetail{
		{
			CountryCode:    "001",
			NationalNumber: "920-265-7335",
		},
	}

	invoice.Invoicer.Website = "https://www.joshuaconstine.com"
	invoice.Invoicer.LogoUrl = "https://www.joshuaconstine.com/logo.png"
	invoice.Invoicer.TaxId = "XX-XXXXXXX"

	return nil
}

func AddRecipientDetailsToInvoice(invoice *payments.CreateInvoiceRequest, db *sql.DB, bookingInformation BookingInformation) error {
	//get recipient details from db
	//fill in the primary_recipients field

	recipient := paypal.InvoiceRecipientInfo{
		BillingInfo: paypal.InvoiceBillingInfo{
			EmailAddress: bookingInformation.User.Email,
			Phones: []paypal.InvoicerPhoneDetail{
				{
					CountryCode:    "001",
					NationalNumber: "920-265-7335",
				},
			},
		},
	}

	invoice.PrimaryRecipients = append(invoice.PrimaryRecipients, recipient)
	return nil

}

func AddItemsToInvoice(invoice *payments.CreateInvoiceRequest, db *sql.DB, bookingInformation BookingInformation) error {
	//get booking details from db
	//fill in the items field
	var items []paypal.InvoiceItem

	for _, bookingCost := range bookingInformation.CostItems {

		name, err := GetBookingCostTypeNameFromID(bookingCost.BookingCostTypeID, db)
		if err != nil {
			return fmt.Errorf("failed to get booking cost type name: %v", err)
		}

		stringAmmount := fmt.Sprintf("%f", bookingCost.Amount)

		item := paypal.InvoiceItem{
			Name:        name,
			Description: "",
			Quantity:    "1",
			UnitAmount: paypal.Money{
				Currency: "USD",
				Value:    stringAmmount,
			},
		}
		items = append(items, item)
	}

	invoice.Items = items

	return nil
}

func CreateInvoiceRequestForBooking(bookingID string, db *sql.DB, client *paypal.Client) (*payments.CreateInvoiceRequest, error) {
	//get booking details from db
	information, err := GetInformationForBookingID(bookingID, db)
	if err != nil {
		return nil, fmt.Errorf("failed to get booking details: %v", err)
	}

	paypalInvoice := payments.CreateInvoiceRequest{
		Detail:            paypal.InvoiceDetail{},
		Invoicer:          paypal.InvoicerInfo{},
		PrimaryRecipients: []paypal.InvoiceRecipientInfo{},
		Items:             []paypal.InvoiceItem{},
		Configuration:     paypal.InvoiceConfiguration{},
		Amount:            paypal.AmountSummaryDetail{},
	}

	err = payments.AddNextInvoiceNumberToInvoice(context.Background(), client, &paypalInvoice)
	if err != nil {
		return nil, fmt.Errorf("failed to add next invoice number to invoice: %v", err)
	}
	//fill in the invoice detail
	err = AddMerchantDetailsToInvoice(&paypalInvoice, db)
	if err != nil {

		return nil, fmt.Errorf("failed to add merchant details to invoice: %v", err)
	}

	//fill in the configuration and amount fields
	SetConstantsForInvoice(&paypalInvoice)

	//fill in the invoicer field
	err = AddInvoicerDetailsToInvoice(&paypalInvoice, db)
	if err != nil {
		return nil, fmt.Errorf("failed to add invoicer details to invoice: %v", err)
	}

	//fill in the primary_recipients field
	err = AddRecipientDetailsToInvoice(&paypalInvoice, db, information)
	if err != nil {
		return nil, fmt.Errorf("failed to add recipient details to invoice: %v", err)

	}

	//fill in the items field
	err = AddItemsToInvoice(&paypalInvoice, db, information)
	if err != nil {
		return nil, fmt.Errorf("failed to add items to invoice: %v", err)
	}

	//log the object
	// Log the struct with fields and values
	log.Println(fmt.Sprintf("%+v", paypalInvoice))

	//update booking details with invoice id
	return &paypalInvoice, nil

}

func HandleCreateInvoiceForBooking(bookingId string, db *sql.DB) error {
	//create invoice request
	client := payments.CreatePaypalClient()

	paypalInvoice, err := CreateInvoiceRequestForBooking(bookingId, db, client)
	if err != nil {
		return fmt.Errorf("failed to create invoice request: %v", err)
	}

	invoiceBytes, err := json.Marshal(paypalInvoice)
	if err != nil {
		return fmt.Errorf("failed to marshal invoice: %v", err)
	}

	//create invoice
	createdInvoice, err := payments.SendInvoiceToPaypal(context.Background(), client, invoiceBytes)

	if err != nil {
		return fmt.Errorf("failed to create invoice: %v", err)
	}

	if createdInvoice {
		//log success
		fmt.Printf("created invoice number: %s", paypalInvoice.Detail.InvoiceNumber)

	}

	//update booking details with invoice id
	return nil
}

func HandleCreateInvoiceForBookingHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	vars := mux.Vars(r)
	bookingId := vars["id"]

	err := HandleCreateInvoiceForBooking(bookingId, db)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create invoice: %v", err), http.StatusInternalServerError)

		return
	}

}
