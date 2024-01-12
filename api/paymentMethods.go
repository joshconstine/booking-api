package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type PaymentMethod struct {
	ID   int
	Name string
}

func GetPaymentMethods(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM payment_method")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var paymentMethods []PaymentMethod

	for rows.Next() {
		var paymentMethod PaymentMethod
		if err := rows.Scan(&paymentMethod.ID, &paymentMethod.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentMethods)

}
