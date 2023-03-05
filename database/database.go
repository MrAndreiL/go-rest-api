package database

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


func Connect() *sql.DB {
    Db, err := sql.Open("mysql", "user:user_password@tcp(127.0.0.1:6033)/app_db")

    if err != nil {
        fmt.Println("Error occured when connecting to database.")
        panic(err.Error())
    }
    fmt.Println("Connection established successfully.")

    return Db
}

func CloseConnection(db *sql.DB) {
    err := db.Close()

    if err != nil {
        fmt.Println("Error occured when closing database connection")
        panic(err.Error())
    }
    fmt.Println("Connection closed successfully.")
}
