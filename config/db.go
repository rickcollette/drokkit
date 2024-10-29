// ./config/db.go
package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"drokkit/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase initializes the database connection with environment variable configuration
func InitDatabase(inMemory bool) (*gorm.DB, *sql.DB) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, proceeding with environment variables")
	}

	// Get database configuration from environment variables
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")

	// Fallback to SQLite in-memory for testing if not provided
	if inMemory {
		dbDriver = "sqlite"
		dbSource = ":memory:"
	} else if dbDriver == "" || dbSource == "" {
		// Default to SQLite if not set in production
		dbDriver = "sqlite"
		dbSource = "game.db"
	}

	var db *gorm.DB
	var err error

	// Switch driver based on environment configuration
	switch dbDriver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dbSource), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	default:
		log.Fatalf("Unsupported DB_DRIVER: %s", dbDriver)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connected!")

	// Migrate the schema
	err = db.AutoMigrate(
		&models.Player{},
		&models.Stats{},
		&models.Match{},
		&models.GameInstance{},
		&models.Faction{},
		&models.FactionMember{},
		&models.Alliance{},
		&models.AllianceMember{},
		&models.AllianceChat{},
		&models.Resource{},
		&models.CombatLog{},
		&models.CombatEvent{},
		&models.VictoryCondition{},
		&models.Admin{},
		&models.Zone{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
	log.Println("Database migrated!")

	// Get the underlying sql.DB instance for connection pooling settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Connection pooling configuration
	maxIdleConns := 10                  // Example: Max idle connections in pool
	maxOpenConns := 100                 // Example: Max open connections in pool
	connMaxLifetime := 30 * time.Minute // Example: Max lifetime of a connection

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	log.Printf("Database connection pooling configured: maxIdleConns=%d, maxOpenConns=%d, connMaxLifetime=%v", maxIdleConns, maxOpenConns, connMaxLifetime)

	return db, sqlDB
}
