package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`   
}

type application struct{
	config config	
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "application environment (development|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	
	app := &application{
		config: cfg,
		logger: logger,	
	}
    srv := &http.Server{
		Addr : fmt.Sprintf(":%d", cfg.port), // We assigned cfg.port in the int format thanks to ":%d"
		Handler : app.routes(),	// what handler do you wanna use?
		IdleTimeout : time.Minute, // How long do you want time out
        ReadTimeout : 10 * time.Second,
		WriteTimeout :  30 * time.Second,
	}
    
	logger.Println("Starting server on port", cfg.port)

	err := srv.ListenAndServe()
	
	if err != nil {
		log.Println(err)
	}

}
