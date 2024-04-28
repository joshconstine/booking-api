package payments

import (
	"booking-api/config"
	"booking-api/data/request"
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

var client *paypal.Client

func CreatePaypalClient(env config.EnvVars) error {
	// load config

	// Create a paypal client
	var err error
	client, err = paypal.NewClient(env.PAYPAL_CLIENT_ID, env.PAYPAL_CLIENT_SECRET, paypal.APIBaseSandBox)
	if err != nil {
		log.Fatalf("failed to create paypal client: %v", err)
	}

	return nil

}

func GetInvoiceByID(ctx context.Context, invoiceID string) (*paypal.Invoice, error) {
	invoice, err := client.GetInvoiceDetails(ctx, invoiceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get invoice details: %v", err)
	}
	return invoice, nil
}

func GenerateInvoiceNumber(ctx context.Context) (*paypal.InvoiceNumber, error) {
	invoiceNumber, err := client.GenerateInvoiceNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %v", err)
	}
	return invoiceNumber, nil
}

func AddNextInvoiceNumberToInvoice(ctx context.Context, invoiceDetails *request.CreateInvoiceRequest) error {
	invoiceNumberResult, err := GenerateInvoiceNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to generate invoice number: %v", err)
	}

	invoiceDetails.Detail.InvoiceNumber = invoiceNumberResult.InvoiceNumberValue

	return nil
}

func AddNextInvoiceNumberToJSON(ctx context.Context, str map[string]interface{}) (*string, error) {
	invoiceNumberResult, err := GenerateInvoiceNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate invoice number: %v", err)
	}

	if detail, ok := str["detail"].(map[string]interface{}); ok {
		detail["invoice_number"] = invoiceNumberResult.InvoiceNumberValue
	} else {
		// Create the detail field if it does not exist
		str["detail"] = map[string]interface{}{
			"invoice_number": invoiceNumberResult.InvoiceNumberValue,
		}
	}

	return &invoiceNumberResult.InvoiceNumberValue, nil
}

func CreateInvoice(ctx context.Context, r *http.Request) (*string, error) {

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

	createdId, err := AddNextInvoiceNumberToJSON(ctx, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to add next invoice number to JSON: %v", err)
	}

	// Re-marshal the modified request body.
	modifiedBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal modified request body: %v", err)
	}

	created, err := SendInvoiceToPaypal(ctx, modifiedBodyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to send invoice to paypal: %v", err)
	}

	if !created {
		return nil, fmt.Errorf("failed to create invoice")
	}

	return createdId, nil

}

func SendInvoiceToPaypal(ctx context.Context, byte []byte) (bool, error) {
	requestURL := fmt.Sprintf("%s%s", client.APIBase, "/v2/invoicing/invoices")

	token, err := client.GetAccessToken(ctx)
	if err != nil {
		return false, fmt.Errorf("failed to get access token: %v", err)
	}

	url, err := url.Parse(requestURL)
	if err != nil {
		return false, fmt.Errorf("failed to parse request URL: %v", err)
	}

	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + token.Token},
		},
		Body: io.NopCloser(bytes.NewReader(byte)),
	}

	//create http client
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return false, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return true, nil

}
