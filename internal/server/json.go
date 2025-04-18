package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	json_data, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return err
	}
	w.WriteHeader(status)
	w.Write(json_data)
	return nil
}

func WriteJSONError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	WriteJSONResponse(w, code, errorResponse{
		Error: msg,
	})
}
