package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

func GetUserForUserID(userId string, db *sql.DB) (User, error) {

	var user User

	err := db.QueryRow("SELECT id, firstName, lastName, email, phoneNumber FROM user WHERE id = ?", userId).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	return user, err
}

func AttemptToInsertUser(user User, db *sql.DB) (int64, error) {

	result, err := db.Exec("INSERT INTO user (firstName, lastName, email, phoneNumber) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.PhoneNumber)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	return result.LastInsertId()
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	id, err := AttemptToInsertUser(user, db)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(id)

}

func GetUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT id, firstName, lastName, email, phoneNumber FROM user")

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
