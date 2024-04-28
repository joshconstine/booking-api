package request

import "github.com/plutov/paypal/v4"

type CreateInvoiceRequest struct {
	Detail            paypal.InvoiceDetail          `json:"detail"`
	Invoicer          paypal.InvoicerInfo           `json:"invoicer"`
	PrimaryRecipients []paypal.InvoiceRecipientInfo `json:"primary_recipients"`
	Items             []paypal.InvoiceItem          `json:"items"`
	Configuration     paypal.InvoiceConfiguration   `json:"configuration"`
	Amount            paypal.AmountSummaryDetail    `json:"amount"`
}
