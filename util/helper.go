package util

import (
	"encoding/json"
	"net/http"
)

// RespondWithError sends an error response in JSON format
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON sends a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "JSON Encoding Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
