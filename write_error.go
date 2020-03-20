package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func writeError(w http.ResponseWriter, err error) {
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	log.Printf("Error in request: %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		writeError(w, err)
	}
}

func respondWithJSON(w http.ResponseWriter, v interface{}, err error) {
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, v)
}
