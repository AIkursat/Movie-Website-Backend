package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router{
	router := httprouter.New()


	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)
	
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie )
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	return router
}