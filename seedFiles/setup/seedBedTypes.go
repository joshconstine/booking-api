package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type BedType struct {
	ID   int
	Name string
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	log.Println("connected to PlanetScale")

	err = db.Ping()

	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	bedTypes := []BedType{
		{1, "Twin"},
		{2, "Full"},
		{3, "Queen"},
		{4, "King"},
		{5, "California King"},
		{6, "Bunk Bed"},
		{7, "Sofa Bed"},
		{8, "Futon"},
		{9, "Crib"},
		{10, "Toddler Bed"},
		{11, "Day Bed"},
	}

	for _, bedType := range bedTypes {
		_, err := db.Exec("INSERT INTO bed_type (id, name) VALUES (?, ?)", bedType.ID, bedType.Name)
		if err != nil {
			log.Fatalf("failed to insert: %v", err)
		}
	}

	fmt.Println("seeded bed_type")

	defer db.Close()

}
