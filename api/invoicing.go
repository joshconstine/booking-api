package api

import (
	"booking-api/api/payments"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
