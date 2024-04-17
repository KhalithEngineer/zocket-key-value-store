package main

import (
	"log"
	"net/http"
	"os"
	"zocket-key-value-store/api/router"
)

func main() {
	// Initialize HTTP server
	server := http.NewServeMux()

	// Define API routes
	router.RegisterRoutes(server)

	// Start server
	port := os.Getenv("PORT")
	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}
