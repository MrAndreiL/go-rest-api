package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var isCreated bool
var db *sql.DB

func Connect() {
	// Establish connection.
	var err error
	db, err = sql.Open("mysql", "user:user_password@tcp(127.0.0.1:6033)/app_db")
	if err != nil {
		fmt.Println("Error occurred when connecting to database.")
		panic(err.Error())
	}
	isCreated = true
	fmt.Println("Connection established successfully.")

	// Create tables if not created and seed them with default values.
	err = seedDatabase()
	if err != nil {
		fmt.Println("Error occurred when seeding database with default values.")
		panic(err.Error())
	}
}

func GetDbConnection() *sql.DB {
	if !isCreated {
		Connect()
	}
	return db
}

func seedDatabase() error {
	studentsCreate := `CREATE TABLE IF NOT EXISTS students (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
                                                                 name TEXT NOT NULL,  
                                                                 age INT NOT NULL,  
                                                                 email VARCHAR(50) NOT NULL, 
                                                                 gpa FLOAT NOT NULL)`

	_, err := db.Exec(studentsCreate)
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
		studentsInsert := `INSERT INTO students (name, age, email, gpa) VALUES (?, ?, ?, ?)`

		_, err := db.Exec(studentsInsert, "Lungu Andrei", 21, "lunguandrei759@gmail.com", 2.7)
		if err != nil {
			fmt.Println("Student seed inserting error")
			return err
		}
	}

	// same for doctors table.
	doctorsCreate := `CREATE TABLE IF NOT EXISTS doctors (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                                               name TEXT NOT NULL,
                                                               age INT NOT NULL,
                                                               specialty TEXT NOT NULL)`

	_, err = db.Exec(doctorsCreate)
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
		doctorsInsert := `INSERT INTO doctors (name, age, specialty) VALUES (?, ?, ?)`

		_, err := db.Exec(doctorsInsert, "Tiron Teodor", 19, "neurosurgeon")
		if err != nil {
			fmt.Println("Doctor seed inserting error")
			return err
		}
	}

	return nil
}

func CloseConnection() {
	err := db.Close()

	if err != nil {
		fmt.Println("Error occurred when closing database connection")
		panic(err.Error())
	}
	fmt.Println("Connection closed successfully.")
}
