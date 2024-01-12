package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Alcohol struct {
	ID            int
	Name          string
	AlcoholTpyeID int
}

func main() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	alcohols := []Alcohol{
		{1, "Spotted Cow", 1},
		{2, "Miller Lite", 1},
		{3, "Bud Light", 1},
		{4, "Budweiser", 1},
		{5, "Coors Light", 1},
		{6, "Corona Extra", 1},
		{7, "Heineken", 1},
		{8, "Michelob Ultra", 1},
		{9, "Guinness", 1},
		{10, "Blue Moon", 1},
		{11, "Stella Artois", 1},
		{12, "Dos Equis", 1},
		{13, "Fat Tire", 1},
		{14, "Tanqueray", 3},
		{15, "Bombay Sapphire", 3},
		{16, "Hendrick's", 3},
		{17, "Barcardi", 4},
		{18, "Captain Morgan", 4},
		{19, "Malibu", 4},
		{20, "Bacardi", 4},
		{21, "Jose Cuervo", 5},
		{22, "Patron", 5},
		{23, "Don Julio", 5},
		{24, "Tito's", 6},
		{25, "Grey Goose", 6},
		{26, "Ketel One", 6},
		{27, "Smirnoff", 6},
		{28, "Jack Daniel's", 7},
		{29, "Crown Royal", 7},
		{30, "Jameson", 7},
		{31, "Fireball", 7},
		{32, "Jim Beam", 7},
		{33, "Champagne", 8},
		{34, "Prosecco", 8},
		{35, "Cabernet Sauvignon", 9},
		{36, "Pinot Noir", 9},
		{37, "Merlot", 9},
		{38, "Pinot Grigio", 9},
		{39, "Chardonnay", 9},
		{40, "Sauvignon Blanc", 9},
		{41, "Sprite", 10},
		{42, "Coke", 10},
		{43, "Diet Coke", 10},
		{44, "Diet Sprite", 10},
		{45, "Ginger Ale", 10},
		{46, "Tonic Water", 10},
		{47, "Club Soda", 10},
	}

	// Insert the data into the alcohols table
	for _, alcohol := range alcohols {
		_, err = db.Exec("INSERT INTO alcohol (id, name, alcohol_type_id) VALUES (?, ?, ?)", alcohol.ID, alcohol.Name, alcohol.AlcoholTpyeID)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully seeded alcohol table")

	defer db.Close()

}
