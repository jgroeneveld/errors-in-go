package main

import (
	"net/http"
)

func myWebHandler(w http.ResponseWriter, r *http.Request) {
	model, err := getSomethingFromDB()
	if err != nil {
		writeError(w, err) // provide a general "writeError" function that deals with the errors your app has
		return // dont forget to stop function execution
	}

	writeJSON(w, model)
}
