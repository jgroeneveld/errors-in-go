package main

import "net/http"

func myWebHandler2(w http.ResponseWriter, r *http.Request) {
	model, err := getSomethingFromDB()

	// lets go further by having a dedicated wrapper for all responses. Handlers are the same anyway.
	// and we dont have to use the returns
	respondWithJSON(w, model, err)
}
