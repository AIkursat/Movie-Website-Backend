package main

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // It means allow all requests.

		next.ServeHTTP(w, r)
	})


} 
