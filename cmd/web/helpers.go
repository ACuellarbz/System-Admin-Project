package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Defining global type
type envelope map[string]interface{}

// Constant string containing characters for generating random strings
const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

// Generates a JSON response
func (myApp *App) createJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	// Convert the data to a JSON-formatted string
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// Set the provided headers in the response.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// Reads the "seed" parameter from the request and converts it to an int64.
func (myApp *App) readUserInput(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	// Retrieve the "seed" parameter from the request.
	userInput, err := strconv.ParseInt(params.ByName("seed"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid seed, please ensure that the seed is a base64 int")
	}

	return userInput, nil
}

// Generates a random string of the specified length.
func (myApp *App) generateRandomString(length int) string {
	// Create a slice of runes with the specified length.
	slice := make([]rune, length)

	// Convert the randomStringSource to a slice of runes.
	r := []rune(randomStringSource)

	// Generate a random prime number using the length of the rune slice.
	for i := range slice {
		primeNum, _ := rand.Prime(rand.Reader, len(r))
		unint := primeNum.Uint64()
		len := uint64(len(r))
		slice[i] = r[unint%len]
	}

	return string(slice)
}

