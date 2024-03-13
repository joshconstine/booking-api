package main

import (
	"booking-api/config"
	"booking-api/controllers"
	"booking-api/database"
	"booking-api/jobs"
	"booking-api/repositories"
	"booking-api/router"
	"booking-api/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"booking-api/pkg/shutdown"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

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

	database.Connect(env.DSN)
	database.Migrate()
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.Run()
		log.Println("server started")
	}()

	return func() {
		cleanup()
		// app.Shutdown(nil)
	}, nil
}

func buildServer(env config.EnvVars) (*gin.Engine, func(), error) {

	validate := validator.New()

	//Init Repositories
	bookingRepository := repositories.NewBookingRepositoryImplementation(database.Instance)
	boatRepository := repositories.NewBoatRepositoryImplementation(database.Instance)

	//Init Service
	bookingService := services.NewBookingServiceImplementation(bookingRepository, validate)
	boatService := services.NewBoatServiceImplementation(boatRepository, validate)

	//Init controller
	bookingController := controllers.NewBookingController(bookingService)
	boatController := controllers.NewBoatController(boatService)

	//Router
	router := router.NewRouter(boatController, bookingController)

	// ginRouter := router.InitRouter(routes)

	return router, func() {

	}, nil
}
