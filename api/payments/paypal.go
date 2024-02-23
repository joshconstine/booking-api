package payments

import (
	"booking-api/config"
	"context"
	"fmt"
	"log"

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
