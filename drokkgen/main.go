package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

const envFile = ".env"

// GenerateUUID4 generates a new UUID4 string.
func GenerateUUID4() string {
	return uuid.New().String()
}

// GenerateBase64UUID4 generates a new UUID4 and encodes it as base64.
func GenerateBase64UUID4() string {
	uuidBytes := uuid.New()
	return base64.StdEncoding.EncodeToString(uuidBytes[:])
}

// GenerateEnvContent generates the .env content with the provided configurations.
func GenerateEnvContent() string {
	jwtSecretKey := GenerateUUID4()
	mysqlPassword := GenerateBase64UUID4()
	redisPassword := GenerateBase64UUID4()

	return fmt.Sprintf(`# Generated .env file

# Main Relational DB
DB_DRIVER=mysql
DB_SOURCE=drokkit:%s@tcp(localhost:3306)/megacity?charset=utf8mb4&parseTime=True&loc=Local
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
DB_CONN_MAX_LIFETIME=30m

# Redis DB for NATS and Leaderboards
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=%s

# JWT secret key
JWT_SECRET_KEY=%s

# NATS configuration
NATS_URL=nats://localhost:4222
`, mysqlPassword, redisPassword, jwtSecretKey)
}

// ViewEnvFile displays the existing .env file content.
func ViewEnvFile() error {
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("\nCurrent .env file content:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

// PromptUserChoice displays options if .env file exists and returns the choice.
func PromptUserChoice() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - View existing .env file")
		fmt.Println("2 - Overwrite .env file")
		fmt.Println("3 - Cancel")
		fmt.Print("Enter your choice (1/2/3): ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		choice := strings.TrimSpace(input)
		if choice == "1" || choice == "2" || choice == "3" {
			return choice, nil
		}
		fmt.Println("Invalid choice, please enter 1, 2, or 3.")
	}
}

// WriteEnvFile writes the generated content to the .env file.
func WriteEnvFile(content string) error {
	return os.WriteFile(envFile, []byte(content), 0600)
}

func main() {
	if _, err := os.Stat(envFile); err == nil {
		for {
			choice, err := PromptUserChoice()
			if err != nil {
				log.Fatalf("Error reading choice: %v", err)
			}

			switch choice {
			case "1":
				if err := ViewEnvFile(); err != nil {
					log.Fatalf("Error viewing .env file: %v", err)
				}
			case "2":
				content := GenerateEnvContent()
				if err := WriteEnvFile(content); err != nil {
					log.Fatalf("Error writing .env file: %v", err)
				}
				fmt.Println("\n.env file has been overwritten with new configurations.")
				return
			case "3":
				fmt.Println("Operation canceled.")
				return
			}
		}
	} else {
		content := GenerateEnvContent()
		if err := WriteEnvFile(content); err != nil {
			log.Fatalf("Error writing .env file: %v", err)
		}
		fmt.Println(".env file created successfully with new configurations.")
	}
}
