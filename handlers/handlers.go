// handlers/handlers.go

package handlers

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db     *gorm.DB
	JwtKey = []byte(os.Getenv("JWT_SECRET_KEY")) // Load JWT key from env var
	nc     *nats.Conn
)

// Claims structure with UserID included for authentication
type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"id"`
	jwt.RegisteredClaims
}

// InitHandlers sets up the shared database and NATS connection instances
func InitHandlers(database *gorm.DB, natsConn *nats.Conn) {
	db = database
	nc = natsConn
	if len(JwtKey) == 0 {
		log.Fatal("JWT_SECRET_KEY environment variable is required but not set")
	}
}

// PublishAlert sends an alert message to a specific NATS subject
func PublishAlert(subject, message string) {
	if nc != nil {
		err := nc.Publish(subject, []byte(message))
		if err != nil {
			log.Printf("Failed to send alert: %v", err)
		} else {
			log.Printf("Alert sent to %s: %s", subject, message)
		}
	}
}
