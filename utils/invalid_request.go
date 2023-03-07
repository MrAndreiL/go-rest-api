package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendBadRequestGeneric(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := make(map[string]string)
	response["message"] = message

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal("Fatal error: Cannot serialize given message.")
	}

	w.Write(jsonResponse)
}
