package models

import (
	"database/sql"
	"fmt"
)

type Student struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Email string  `json:"email"`
	Gpa   float64 `json:gpa"`
}
