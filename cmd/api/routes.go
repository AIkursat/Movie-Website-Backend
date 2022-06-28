package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	//"golang.org/x/text/secure"
)

func(app *application) wrap(next http.Handler) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := context.WithValue(r.Context(), "params", ps) // ctx represents the context
    next.ServeHTTP(w, r.WithContext(ctx))
}

}


func (app *application) routes() http.Handler{
	router := httprouter.New()
	secure := alice.New(app.checkToken)



	router.HandlerFunc(http.MethodGet, "/status", app.StatusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/graphql", app.moviesGraphQL)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)
	
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie )
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres )

     router.POST("/v1/admin/editmovie", app.wrap(secure.ThenFunc(app.editmovie)))

	// router.HandlerFunc(http.MethodPost, "/v1/admin/editmovie", app.editmovie )
    
	router.GET("/v1/admin/deletemovie/:id", app.wrap(secure.ThenFunc(app.deleteMovie)))
	//router.HandlerFunc(http.MethodGet, "/v1/admin/deletemovie/:id", app.deleteMovie )

	return app.enableCORS(router)
}