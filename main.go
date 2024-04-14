package main

import (
	"booking-api/api"
	"booking-api/config"
	"booking-api/jobs"
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
	// setup exit code for graceful shutdown

	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig(".")
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}
	//test

	// run the server
	cleanup, err := run(env)

	// Run jobs
	jobs.VerifyBookingStatuses() // Call the function to run the jobs
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
	corsOpts := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // Allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Application/json"}),        // Allowed headers
		handlers.AllowCredentials(),                                                  // Credentials
	)
	// Open a connection to PlanetScale
	db, err := config.ConnectToDB(env)
	if err != nil {
		return nil, nil, err
	}

	api.InitRoutes(r, db)

	server := &http.Server{
		Addr:    ":" + env.PORT,
		Handler: corsOpts(r), // Wrap the router with the CORS handler
	}

	return server, func() {
		if err := server.Close(); err != nil {
			log.Printf("error: %v", err)
		}
	}, nil
}
