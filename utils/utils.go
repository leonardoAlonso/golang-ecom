package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("Missing request body")
	}
	// Decode the request body into the payload struct
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status_code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status_code int, err error) {
	WriteJson(w, status_code, map[string]string{"error": err.Error()})
}
