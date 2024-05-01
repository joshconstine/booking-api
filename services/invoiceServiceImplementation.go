package services

import (
	payments "booking-api/pkg/payment"
	"booking-api/repositories"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"

	"booking-api/data/request"
	"booking-api/data/response"

	"github.com/plutov/paypal/v4"
)

type InvoiceServiceImplementation struct {
	bookingRepository repositories.BookingRepository
}

func NewInvoiceServiceImplementation(
	bookingRepository repositories.BookingRepository,
) InvoiceService {
	return &InvoiceServiceImplementation{
		bookingRepository: bookingRepository,
	}
}

func CreateInvoiceHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//invoice is a struct that holds the invoice details
	//decode the request body into struct and check for errors

	// Create invoice
	createdInvoice, err := payments.CreateInvoice(r.Context(), r)
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

func AddMerchantDetailsToInvoice(invoice *request.CreateInvoiceRequest) error {
	//get merchant details from db
	//fill in the invoicer field
	invoice.Detail.Note = "Thank you for your business"
	return nil
}

func SetConstantsForInvoice(invoice *request.CreateInvoiceRequest) {
	//fill in the configuration field
	//fill in the amount field
	invoice.Detail.CurrencyCode = "USD"
	invoice.Detail.PaymentTerm.TermType = "NET_30"

	invoice.Configuration.PartialPayment.AllowPartialPayment = true
	invoice.Configuration.PartialPayment.MinimumAmountDue = paypal.Money{
		Currency: "USD",
		Value:    "0.00",
	}

	invoice.Configuration.AllowTip = true
	invoice.Configuration.TaxCalculatedAfterDiscount = true
	invoice.Configuration.TaxInclusive = false
	invoice.Configuration.TemplateId = "TEMP-8M018176VU180072F"
}

func AddInvoicerDetailsToInvoice(invoice *request.CreateInvoiceRequest) error {
	//get invoicer details from db
	//fill in the invoicer field
	invoice.Invoicer.EmailAddress = "joshua.constine97@gmail.com"
	invoice.Invoicer.Phones = []paypal.InvoicerPhoneDetail{
		{
			CountryCode:    "001",
			NationalNumber: "920-265-7335",
		},
	}

	invoice.Invoicer.Phones = []paypal.InvoicerPhoneDetail{
		{
			CountryCode:    "001",
			NationalNumber: "9202657335",
		},
	}

	invoice.Invoicer.Website = "https://www.joshuaconstine.com"
	invoice.Invoicer.LogoUrl = "https://www.joshuaconstine.com/logo.png"
	invoice.Invoicer.TaxId = "XX-XXXXXXX"

	return nil
}

func AddRecipientDetailsToInvoice(invoice *request.CreateInvoiceRequest, bookingInformation *response.BookingInformationResponse) error {
	//get recipient details from db
	//fill in the primary_recipients field

	recipient := paypal.InvoiceRecipientInfo{
		BillingInfo: paypal.InvoiceBillingInfo{
			// EmailAddress: bookingInformation.User.Email,
			EmailAddress: "joshua.constine97@gmail.com",
			Phones: []paypal.InvoicerPhoneDetail{
				{
					CountryCode:    "001",
					NationalNumber: "9202657335",
				},
			},
		},
		ShippingInfo: paypal.InvoiceContactInfo{

			BusinessName: "Joshua Constine",
			RecipientAddress: paypal.InvoiceAddressPortable{
				AddressLine1:   "1234 Elm Street",
				PostalCode:     "95121",
				CountryCode:    "US",
				AddressDetails: paypal.InvoiceAddressDetails{},
			},
		},
	}

	invoice.PrimaryRecipients = append(invoice.PrimaryRecipients, recipient)
	return nil

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func AddItemsToInvoice(invoice *request.CreateInvoiceRequest, bookingInformation *response.BookingInformationResponse) error {
	//get booking details from db
	//fill in the items field
	var items []paypal.InvoiceItem

	itemTotalValue := 0.0
	taxTotalValue := 0.0
	for _, bookingCost := range bookingInformation.CostItems {

		// name, err := GetBookingCostTypeNameFromID(bookingCost.BookingCostTypeID, db)
		// if err != nil {
		// 	return fmt.Errorf("failed to get booking cost type name: %v", err)
		// }

		name := "Booking Cost"
		stringAmmount := fmt.Sprintf("%2f", bookingCost.Amount)
		itemTotalValue += bookingCost.Amount

		taxRateString := fmt.Sprintf("%2f", bookingCost.TaxRate.Percentage)

		taxAmmountForItem := bookingCost.Amount * bookingCost.TaxRate.Percentage / 100
		taxAmmountForItem = toFixed(taxAmmountForItem, 2)

		//truncate to 2 decimal places

		// taxAmmountForItem = math.Ceil(taxAmmountForItem*100) / 100

		taxTotalValue += taxAmmountForItem
		// taxTotalValue += .01

		item := paypal.InvoiceItem{
			Name:        name,
			Description: "",
			Quantity:    "1",
			UnitAmount: paypal.Money{
				Currency: "USD",
				Value:    stringAmmount,
			},

			Tax: paypal.InvoiceTax{
				Name:    bookingCost.TaxRate.Name,
				Percent: taxRateString,
				Amount: paypal.Money{
					Currency: "USD",
					Value:    fmt.Sprintf("%.2f", taxAmmountForItem),
				},
			},
			InvoiceDiscount: paypal.InvoicingDiscount{
				DiscountAmount: paypal.Money{
					Currency: "USD",
					Value:    "0.00",
				},
			},
		}
		items = append(items, item)
	}

	invoice.Items = items

	itemTotalValue = toFixed(itemTotalValue, 2)
	taxTotalValue = toFixed(taxTotalValue, 2)

	value := itemTotalValue + taxTotalValue
	value = toFixed(value, 2)
	valueStr := fmt.Sprintf("%.2f", value)

	ammount := paypal.AmountSummaryDetail{
		Currency: "USD",
		Value:    valueStr,
		Breakdown: paypal.InvoiceAmountWithBreakdown{

			Custom: paypal.CustomAmount{
				Amount: paypal.Money{
					Currency: "USD",
					Value:    "0.00",
				},
				Label: "booking",
			},
			Discount: paypal.AggregatedDiscount{
				InvoiceDiscount: paypal.InvoicingDiscount{
					DiscountAmount: paypal.Money{
						Currency: "USD",
						Value:    "0.00",
					},
					Percent: "0.00",
				},
			},
			Shipping: paypal.InvoiceShippingCost{
				Tax: paypal.InvoiceTax{
					Name:    "Sales Tax",
					Percent: "0.00",
					Amount: paypal.Money{
						Currency: "USD",
						Value:    "0.00",
					},
				},
				Amount: paypal.Money{
					Currency: "USD",
					Value:    "0.00",
				},
			},
			ItemTotal: paypal.Money{
				Currency: "USD",
				Value:    fmt.Sprintf("%.2f", itemTotalValue),
			},
			TaxTotal: paypal.Money{
				Currency: "USD",
				Value:    fmt.Sprintf("%.2f", taxTotalValue),
			},
		},
	}
	invoice.Amount = ammount

	return nil
}

func CreateInvoiceRequestForBooking(booking *response.BookingInformationResponse) (*request.CreateInvoiceRequest, error) {

	paypalInvoice := request.CreateInvoiceRequest{
		Detail:            paypal.InvoiceDetail{},
		Invoicer:          paypal.InvoicerInfo{},
		PrimaryRecipients: []paypal.InvoiceRecipientInfo{},
		Items:             []paypal.InvoiceItem{},
		Configuration:     paypal.InvoiceConfiguration{},
		// Amount:            paypal.AmountSummaryDetail{},
	}

	err := payments.AddNextInvoiceNumberToInvoice(context.Background(), &paypalInvoice)
	if err != nil {
		return nil, fmt.Errorf("failed to add next invoice number to invoice: %v", err)
	}
	//fill in the invoice detail
	err = AddMerchantDetailsToInvoice(&paypalInvoice)
	if err != nil {

		return nil, fmt.Errorf("failed to add merchant details to invoice: %v", err)
	}

	//fill in the configuration and amount fields
	SetConstantsForInvoice(&paypalInvoice)

	//fill in the invoicer field
	err = AddInvoicerDetailsToInvoice(&paypalInvoice)
	if err != nil {
		return nil, fmt.Errorf("failed to add invoicer details to invoice: %v", err)
	}

	//fill in the primary_recipients field
	err = AddRecipientDetailsToInvoice(&paypalInvoice, booking)
	if err != nil {
		return nil, fmt.Errorf("failed to add recipient details to invoice: %v", err)

	}

	//fill in the items field
	err = AddItemsToInvoice(&paypalInvoice, booking)
	if err != nil {
		return nil, fmt.Errorf("failed to add items to invoice: %v", err)
	}

	//turn the invoice into json

	invoiceBytes, err := json.Marshal(paypalInvoice)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal invoice: %v", err)
	}
	log.Println(string(invoiceBytes))

	//log the object
	// Log the struct with fields and values
	// log.Println(fmt.Sprintf("%+v", paypalInvoice))

	//update booking details with invoice id
	return &paypalInvoice, nil

}

func (invs InvoiceServiceImplementation) CreateInvoiceForBooking(bookingId string) (string, error) {

	//create invoice request
	booking := invs.bookingRepository.FindById(bookingId)

	paypalInvoice, err := CreateInvoiceRequestForBooking(&booking)
	if err != nil {
		return "", fmt.Errorf("failed to create invoice request: %v", err)
	}

	invoiceBytes, err := json.Marshal(paypalInvoice)
	if err != nil {
		return "", fmt.Errorf("failed to marshal invoice: %v", err)
	}

	//create invoice
	createdInvoice, err := payments.SendInvoiceToPaypal(context.Background(), invoiceBytes)

	if err != nil {
		return "", fmt.Errorf("failed to send invoice to paypal: %v", err)
	}

	if createdInvoice {
		//log success
		fmt.Printf("created invoice number: %s", paypalInvoice.Detail.InvoiceNumber)

	}

	//update booking details with invoice id
	return paypalInvoice.Detail.InvoiceNumber, nil
}
