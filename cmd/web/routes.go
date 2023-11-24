package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Restful http router
func (myApp *App) routes() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/randomstring/:seed", myApp.makeRandomPassword)
	return router
}
