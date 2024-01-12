package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type RefundStatus struct {
	ID   int
	Name string
}

func GetRefundStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT * FROM refund_status")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var refundStatus []RefundStatus

	for rows.Next() {
		var status RefundStatus
		if err := rows.Scan(&status.ID, &status.Name); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		refundStatus = append(refundStatus, status)
	}

	// Return the data as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(refundStatus)

}
