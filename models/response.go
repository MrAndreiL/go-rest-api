package models

import (
	"encoding/json"
	"errors"
	"log"
)

type RequestError struct {
	ErrorCode int
	Error     error
}

func CreateError(code int, message string) *RequestError {
	return &RequestError{
		ErrorCode: code,
		Error:     errors.New(message),
	}
}

func JsonErrorResponseMessage(message string) []byte {
	resp := make(map[string]string)

	resp["message"] = message

	jsonMessage, err := json.Marshal(resp)
	if err != nil {
		log.Fatal("Error occurred when marshalling json.")
	}

	return jsonMessage
}

func JsonStudentEncoding(student *Student) []byte {
	jsonMessage, err := json.Marshal(student)

	if err != nil {
		log.Fatal("Error occured when marshalling json.")
	}

	return jsonMessage
}
