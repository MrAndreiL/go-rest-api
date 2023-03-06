package main

import (
	"github.com/MrAndreiL/go-rest-api/database"
)

func main() {
	// Open database connection.
	db := database.Connect()

	// Close database connection after function ends.
	defer database.CloseConnection(db)
}
