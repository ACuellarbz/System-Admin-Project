package main

import "net/http"

// logError logs an error message using the application's logger
func (myApp *App) logError(r *http.Request, err error) {
	myApp.logger.Println(err)
}

// Handles JSON formatted error responses
func (myApp *App) jsonError(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// Create an envelope for the error message
	env := envelope{"error": message}

	// Write the JSON response
	err := myApp.createJSON(w, status, env, nil)

	if err != nil {
		myApp.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Handles internal server errors
func (myApp *App) serverError(w http.ResponseWriter, r *http.Request, err error) {
	myApp.logError(r, err)

	message := "the server encountered a problem and could not process the request"

	// Respond with a JSON-formatted error message and Internal Server Error
	myApp.jsonError(w, r, http.StatusInternalServerError, message)
}

// Handles errors where the input exceeds a specified maximum value
func (myApp *App) exceedsMaxValueError(w http.ResponseWriter, r *http.Request, err error) {
	myApp.logError(r, err)

	message := err.Error()

	myApp.jsonError(w, r, http.StatusInternalServerError, message)
}
