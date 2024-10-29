package main

import (
	"log"
	"net/http"
	"os"

	"drokkit/config"
	"drokkit/handlers"
	"drokkit/routes"
	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func main() {
	// Load environment variables for configuration
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL // Fallback if not set
	}

	// Initialize Database
	db, sqlDB := config.InitDatabase(false)
	defer sqlDB.Close()

	// Initialize NATS connection
	var err error
	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	log.Println("Connected to NATS!")

	// Pass the DB and NATS instance to handlers
	handlers.InitHandlers(db, nc)

	// Initialize router
	router := routes.InitRoutes()

	// Start the server
	log.Println("Starting game server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
