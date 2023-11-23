package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Defining global envelope
type envelope map[string]interface {
}

// Source string
const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"


func (myApp *App) createJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//the actual conversion
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func (myApp *App) readUserInput(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	user_input, err := strconv.ParseInt(params.ByName("seed"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid seed, please ensure that the seed is a base64 int")
	}
	return user_input, nil
}

func (myApp *App) generateRandomString(length int) string {
	slice := make([]rune, length) //creates a slice of type run wuth a length of 'length'
	r := []rune(randomStringSource)//creates a slice of runs from randomStringSource

	for i := range slice {
		prime_num, _ := rand.Prime(rand.Reader, len(r))
		unint := prime_num.Uint64()// converts the Prime numbers to Uint64
		len := uint64(len(r)) //length of r is set to 
		slice[i] = r[unint%len] 
	}

	return string(slice)

}