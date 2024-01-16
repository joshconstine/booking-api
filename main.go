package main

import (
	"booking-api/api"
	"booking-api/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"booking-api/pkg/shutdown"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env)

	// run the cleanup after the server is terminated
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	shutdown.Gracefully()

}

func run(env config.EnvVars) (func(), error) {

	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.ListenAndServe()
		log.Println("server started")
	}()

	return func() {
		cleanup()
		app.Shutdown(nil)
	}, nil
}

func buildServer(env config.EnvVars) (*http.Server, func(), error) {
	r := mux.NewRouter()

	// Configure CORS
	corsOpts := handlers.AllowedOrigins([]string{"http://localhost:3000"}) // Update this with the allowed origin(s)
	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", env.DSN)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("connected to PlanetScale")

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping: %v", err)
	}

	api.InitRoutes(r, db)

	server := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: handlers.CORS(corsOpts)(r), // Wrap the router with the CORS handler
	}

	return server, func() {
		if err := server.Close(); err != nil {
			log.Printf("error: %v", err)
		}
	}, nil
}
