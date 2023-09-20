package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON format
	bytes, err := json.Marshal(payload)
	if err != nil {
		// Log the error and send a 500 Internal Server Error response
		log.Printf("respondWithJSON: error: %v\nfailed to marshal JSON response: %v", err, payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal JSON response"))
		return
	}
	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("respondWithError: error: %s\nresponding with 5xx error", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
