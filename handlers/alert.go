package handlers

import (
	"github.com/nats-io/nats.go"
	"log"
)

// SendAlert publishes an alert message to a NATS subject
func SendAlert(nc *nats.Conn, subject, message string) {
	err := nc.Publish(subject, []byte(message))
	if err != nil {
		log.Printf("Failed to send alert: %v", err)
	}
	log.Printf("Alert sent on subject %s: %s", subject, message)
}
