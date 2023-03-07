package models

import (
	"database/sql"
	"io"

	"github.com/MrAndreiL/go-rest-api/database"
	"github.com/MrAndreiL/go-rest-api/utils"
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
		return JsonErrorResponseMessage("Server error occurred when processing request"), 500
	}

	return JsonStudentEncoding(&student), 200
}

func PutStudent(id int, body io.ReadCloser) ([]byte, int) {
	student, err := JsonStudentDecoding(body)

	if err != nil { // bad serialization due to redundant or erroneous fields
		return JsonErrorResponseMessage("Cannot update specified resource."), 500
	}

	db := database.GetDbConnection()

	if !isItemValid(id, db) {
		return JsonErrorResponseMessage("The item does not exist"), 404
	}

	// now, validate given data
	if !utils.IsEmailValid(student.Email) {
		return JsonErrorResponseMessage("Invalid email."), 400
	}
	if !utils.IsAgeValid(student.Age) {
		return JsonErrorResponseMessage("Invalid age"), 400
	}
	if !utils.IsGpaValid(student.Gpa) {
		return JsonErrorResponseMessage("Invalid GPA"), 400
	}

	// update database
	cmd := "UPDATE students SET name = ?, age = ?, email = ?, gpa = ? WHERE id = ?"

	res, err := db.Exec(cmd, student.Name, student.Age, student.Email, student.Gpa, id)
	if err != nil {
		return JsonErrorResponseMessage("Server error occurred when processing request"), 500
	}

	// check how many rows were affected, if 0, then the body matches the existent item.
	numAffected, err := res.RowsAffected()
	if err != nil {
		return JsonErrorResponseMessage("Server error occurred when processing request"), 500
	}

	if numAffected == 0 {
		return JsonErrorResponseMessage("The item already exists in the given form"), 409
	}

	return nil, 204
}

func DeleteStudent(id int) ([]byte, int) {
	db := database.GetDbConnection()

	if !isItemValid(id, db) {
		return JsonErrorResponseMessage("The item does not exist"), 404
	}

	cmd := "DELETE FROM students WHERE id = ?"

	_, err := db.Exec(cmd, id)
	if err != nil {
		return JsonErrorResponseMessage("Server error occurred when processing request"), 500
	}

	return JsonErrorResponseMessage("Item was deleted successfully"), 200
}

func isItemValid(id int, db *sql.DB) bool {
	res := db.QueryRow("SELECT * FROM students WHERE id = ?", id)

	var student Student
	err := res.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Gpa)

	if err == sql.ErrNoRows {
		return false
	}
	return true
}
