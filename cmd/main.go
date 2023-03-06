package main

import (
	"log"
	"net/http"

	"github.com/MrAndreiL/go-rest-api/database"
	"github.com/MrAndreiL/go-rest-api/router"
)

func main() {
	// Open database connection.
	database.Connect()

	// Close database connection after function ends.
	defer database.CloseConnection()

	// register routing mechanism
	router.HandleRequests()

	// open HTTP server
	log.Fatal(http.ListenAndServe(":3001", nil))
}
