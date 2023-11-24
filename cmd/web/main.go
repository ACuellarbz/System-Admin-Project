package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// AppConfig represents the configuration for the application.
type AppConfig struct {
	ServerPort int    // Port on which the server will listen.
	Environment  string
}

// App represents the main application structure.
type App struct {
	Config AppConfig        // Configuration for the application.
	logger *log.Logger   // Logger for logging information.
}

func main() {
	// Configuration flags
	var appConfig AppConfig

	// Setting up server configurations
	appConfig.ServerPort = 4000
	appConfig.Environment = "production"

	// Creating a logger that writes to os.Stdout with date and time prefixes
	logger := log.New(os.Stdout, "APP: ", log.Ldate|log.Ltime)

	// Creating an instance of the application structure
	myApp := &App{
		Config: appConfig,
		logger: logger,
	}

	// Creating the HTTP Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.ServerPort),
		Handler: myApp.routes(), // Setting the server's handler using the routes method of the application.
	}

	// Starting the server
	logger.Printf("Starting %s server on %s", appConfig.Environment, server.Addr)
	err := server.ListenAndServe()
	logger.Fatal(err) // Log any errors and exit if there is an error.
}
