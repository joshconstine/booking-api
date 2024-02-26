package payments

import (
	"booking-api/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func GenerateInvoiceNumber(ctx context.Context, client *paypal.Client) (*paypal.InvoiceNumber, error) {
	invoiceNumber, err := client.GenerateInvoiceNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %v", err)
	}
	return invoiceNumber, nil
}

func CreateInvoice(ctx context.Context, client *paypal.Client, r *http.Request) (*string, error) {
	requestURL := fmt.Sprintf("%s%s", client.APIBase, "/v2/invoicing/invoices")
	// Generate the next invoice number.
	invoiceNumberResult, err := GenerateInvoiceNumber(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %v", err)
	}

	// Read the request body.
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %v", err)
	}
	defer r.Body.Close()

	// Update the request body with the new invoice number.
	var requestBody map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %v", err)
	}

	if detail, ok := requestBody["detail"].(map[string]interface{}); ok {
		detail["invoice_number"] = invoiceNumberResult.InvoiceNumberValue
	} else {
		// Create the detail field if it does not exist
		requestBody["detail"] = map[string]interface{}{
			"invoice_number": invoiceNumberResult.InvoiceNumberValue,
		}
	}

	// Re-marshal the modified request body.
	modifiedBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal modified request body: %v", err)
	}

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
		Body: io.NopCloser(bytes.NewReader(modifiedBodyBytes)),
	}

	//create http client
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return &invoiceNumberResult.InvoiceNumberValue, nil

}
