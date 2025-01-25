package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nomannaq/simpledbgolang/internal/database"
	"github.com/nomannaq/simpledbgolang/internal/handlers"
)

func main() {
	// Create a new database instance
	dbon := database.NewDatabase()

	// Define HTTP routes
	http.HandleFunc("/set", handlers.SetHandler(dbon))
	http.HandleFunc("/get", handlers.GetHandler(dbon))
	http.HandleFunc("/delete", handlers.DeleteHandler(dbon))

	// Start the HTTP server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
