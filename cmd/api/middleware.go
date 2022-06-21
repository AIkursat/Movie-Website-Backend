package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // It means allow all requests.
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		next.ServeHTTP(w, r)
	})


} 

func (app *application) checkToken(next http.Handler) http.Handler { // it returns http.http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("vary", "Authorization",)

		authHeader := r.Header.Get("Authorization")

		if authHeader == ""{
			// could set an anonymous user
		}

		headerParts := strings.Split(authHeader, " ")

		if len(headerParts) != 2{
			app.errorJSON(w, errors.New("invalid auth header"))
			return
		}
		
		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("unauthorized - no bearer"))
			return
		}


		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret)) // HMAC Check
		if err != nil{
			app.errorJSON(w, errors.New("unauthorized - failed hmac check"))
		}

		if !claims.Valid(time.Now()){ // Is Token still Valid at this moment?
          app.errorJSON(w, errors.New("unauthorized - token expired"))
		}

		if !claims.AcceptAudience("mydomain.com"){ 
		app.errorJSON(w, errors.New("unauthorized - invalid"))
		return
		}

		if claims.Issuer != "mydomain.com"{
			app.errorJSON(w, errors.New("invalid issuer"))
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64) // 64 bit
		if err != nil{
			app.errorJSON(w, errors.New("unauthorized"))
			return
		}

		log.Println("valid user:", userID)

		next.ServeHTTP(w, r)
	})

}