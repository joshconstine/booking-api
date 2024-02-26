package payments

import (
	"booking-api/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/plutov/paypal/v4"
)

func CreatePaypalClient() *paypal.Client {
	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil
	}

	// Create a paypal client
	client, err := paypal.NewClient(env.PAYPAL_CLIENT_ID, env.PAYPAL_CLIENT_SECRET, paypal.APIBaseSandBox)
	if err != nil {
		log.Fatalf("failed to create paypal client: %v", err)
	}

	return client

}

func GetInvoiceByID(ctx context.Context, client *paypal.Client, invoiceID string) (*paypal.Invoice, error) {
	invoice, err := client.GetInvoiceDetails(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get invoice details: %v", err)
	}
	return invoice, nil
}

// func CreateInvoice(ctx context.Context, client *paypal.Client, r *http.Request) (*paypal.Invoice, error) {
// 	// 	curl -v -X GET https://api-m.sandbox.paypal.com/v2/invoicing/invoices?total_required=true \
// 	// -H "Content-Type: application/json" \
// 	// -H "Authorization: Bearer <Token>"

// 	//log the request body

// 	requestUrl := fmt.Sprintf("%s%s", client.APIBase, "/v2/invoicing/invoices")
// 	req, err := client.NewRequest(ctx, "POST", requestUrl, r.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create new request: %v", err)
// 	}

// 	// Send the request
// 	err = client.SendWithAuth(req, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to send request: %v", err)
// 	}

// 	return nil, nil

// }

func GenerateInvoiceNumber(ctx context.Context, client *paypal.Client) (*paypal.InvoiceNumber, error) {
	invoiceNumber, err := client.GenerateInvoiceNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %v", err)
	}
	return invoiceNumber, nil
}

// func CreateInvoice(ctx context.Context, client *paypal.Client, r *http.Request) (*paypal.Invoice, error) {
// 	requestURL := fmt.Sprintf("%s%s", client.APIBase, "/v2/invoicing/invoices")

// 	//log the request body
// 	log.Printf("Request Body: %+v", r.Body)

// 	// Assuming r.Body is the correct JSON payload for creating the invoice
// 	// You might want to log or validate r.Body before sending
// 	req, err := client.NewRequest(ctx, "POST", requestURL, r.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create new request: %v", err)
// 	}

// 	var createdInvoice paypal.Invoice
// 	if err = client.SendWithAuth(req, &createdInvoice); err != nil {
// 		return nil, fmt.Errorf("failed to send request: %v", err)
// 	}

// 	// Optionally, log the created invoice
// 	// log.Printf("Created Invoice: %+v", createdInvoice)

// 	return &createdInvoice, nil
// }

func CreateInvoice(ctx context.Context, client *paypal.Client, r *http.Request) (*paypal.Invoice, error) {
	requestURL := fmt.Sprintf("%s%s", client.APIBase, "/v2/invoicing/invoices")

	// requestURL :=

	token, err := client.GetAccessToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %v", err)
	}

	url, err := url.Parse(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request URL: %v", err)
	}

	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + token.Token},
		},
		Body: r.Body,
	}

	var createdInvoice paypal.Invoice

	//create http client
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	//decode the response into

	if resp != nil {
		defer resp.Body.Close()
	}

	return &createdInvoice, nil

}
