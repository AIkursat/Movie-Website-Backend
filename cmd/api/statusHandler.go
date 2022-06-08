package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Avaible",
		Environment: app.config.env,
		Version:     version,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t") // 1) what do you want to convert into JSON on its current status?
	// 2) do you want to have any kind of prefix?
	// 3 ) how much do you want to indent it?
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
