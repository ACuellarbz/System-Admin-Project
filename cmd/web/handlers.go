package main

import (
	"errors"
	"net/http"
)

// Generates a random string of specified length 
func (myApp *App) makeRandomPassword(w http.ResponseWriter, r *http.Request) {
	userInput, err := myApp.readUserInput(r) // Getting the specified number.
	if err != nil {
		myApp.serverError(w, r, err)
		return
	}

	// Check if the user input exceeds the maximum value.
	if userInput > 600 {
		myApp.exceedsMaxValueError(w, r, errors.New("the int provided exceeds 600"))
		return
	}

	// Generate a random string based on the user input.
	randomString := myApp.generateRandomString(int(userInput))

	// Write the generated random string in JSON format to the response.
	err = myApp.createJSON(w, http.StatusOK, envelope{"Random String": randomString}, nil)
	if err != nil {
		myApp.serverError(w, r, err)
	}
}
