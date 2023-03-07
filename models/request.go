package models

import (
	"encoding/json"
	"io"
)

func JsonStudentDecoding(body io.ReadCloser) (*Student, error) {
	var student Student

	err := json.NewDecoder(body).Decode(&student)

	return &student, err
}
