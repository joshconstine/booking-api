package api

import (
	"database/sql"
	"encoding/json"
	"errors"
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

	rows, err := db.Query("SELECT id, first_name, last_name, email, phone_number FROM user WHERE id = ?", userId)

	//check if there was at least one result
	if rows == nil {
		log.Fatalf("failed to query: %v", err)
	}

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber)
	} else {
		return user, errors.New("No user found for id: " + userId)
	}
	if err != nil {
		log.Fatalf("failed to scan row: %v", err)
	}

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	return user, err
}

func GetUserForEmail(email string, db *sql.DB) (User, error) {

	var user User

	rows, err := db.Query("SELECT id, first_name, last_name, email, phone_number FROM user WHERE email = ?", email)

	//check if there was an error
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	//check if there was a result
	if rows == nil {
		log.Fatalf("failed to query: %v", err)
	}

	if rows.Next() {

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PhoneNumber); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
	} else {
		return user, errors.New("No user found for email: " + email)
	}

	return user, err
}

func AttemptToInsertUser(user User, db *sql.DB) (int64, error) {

	result, err := db.Exec("INSERT INTO user (first_name, last_name, email, phone_number) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.PhoneNumber)
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

	rows, err := db.Query("SELECT id, first_name, last_name, email, phone_number FROM user")

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
