package models

import (
	"database/sql"

	"github.com/MrAndreiL/go-rest-api/database"
)

type Student struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Email string  `json:"email"`
	Gpa   float64 `json:"gpa"`
}

func GetStudent(id int) ([]byte, int) {
	db := database.GetDbConnection()

	res := db.QueryRow("SELECT * FROM students WHERE id = ?", id)

	var student Student
	err := res.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Gpa)

	if err == sql.ErrNoRows { // the given id does not exist or it's invalid.
		return JsonErrorResponseMessage("The item does not exist"), 404
	}

	if err != nil { // database error
		return JsonErrorResponseMessage("Server error occured when processing request"), 500
	}

	return JsonStudentEncoding(&student), 200
}
