package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var isCreated bool
var db *sql.DB

func Connect() *sql.DB {
    // make sure a single instance is created
    if isCreated {
        return db
    }

	// Establish connection.
    var err error
	db, err = sql.Open("mysql", "user:user_password@tcp(127.0.0.1:6033)/app_db")
	if err != nil {
		fmt.Println("Error occured when connecting to database.")
		panic(err.Error())
	}
    isCreated = true
	fmt.Println("Connection established successfully.")

	// Create tables if not created and seed them with default values.
	err = seedDatabase(db)
	if err != nil {
		fmt.Println("Error occred when seeding database with default values.")
		panic(err.Error())
	}

	return db
}

func seedDatabase(db *sql.DB) error {
	students_create := `CREATE TABLE IF NOT EXISTS students (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
                                                                 name TEXT NOT NULL,  
                                                                 age INT NOT NULL,  
                                                                 email VARCHAR(50) NOT NULL, 
                                                                 gpa FLOAT NOT NULL)`

	_, err := db.Exec(students_create)
	if err != nil {
		fmt.Println("Error occurred when creating students table")
		return err
	}

	// if the table is empty, seed it with default values.
	res, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Println("Query error")
		return err
	}

	if res.Next() == false {
		students_insert := `INSERT INTO students (name, age, email, gpa) VALUES (?, ?, ?, ?)`

		_, err := db.Exec(students_insert, "Lungu Andrei", 21, "lunguandrei759@gmail.com", 2.7)
		if err != nil {
			fmt.Println("Student seed inserting error")
			return err
		}
	}

	// same for doctors table.
	doctors_create := `CREATE TABLE IF NOT EXISTS doctors (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                                               name TEXT NOT NULL,
                                                               age INT NOT NULL,
                                                               specialty TEXT NOT NULL)`

	_, err = db.Exec(doctors_create)
	if err != nil {
		fmt.Println("Error occurred when creating doctors table")
		return err
	}

	// if the table is empty, seed it with default values
	res, err = db.Query("SELECT * FROM doctors")
	if err != nil {
		fmt.Println("Query error")
		return err
	}

	if res.Next() == false {
		doctors_insert := `INSERT INTO doctors (name, age, specialty) VALUES (?, ?, ?)`

		_, err := db.Exec(doctors_insert, "Tiron Teodor", 19, "neurosurgeon")
		if err != nil {
			fmt.Println("Doctor seed inserting error")
			return err
		}
	}

	return nil
}

func CloseConnection(db *sql.DB) {
	err := db.Close()

	if err != nil {
		fmt.Println("Error occured when closing database connection")
		panic(err.Error())
	}
	fmt.Println("Connection closed successfully.")
}
