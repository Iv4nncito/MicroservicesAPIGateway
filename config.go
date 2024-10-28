package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// logMessage logs messages to the console with a specified type (e.g., INFO, WARNING, ERROR)
func logMessage(messageType, message string) {
	fmt.Printf("[%s] %s\n", messageType, message)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logMessage("ERROR", fmt.Sprintf("Error loading .env file: %s", err))
		return
	}

	envVar := os.Getenv("ENV_VAR_NAME")
	apiKey := os.Getenv("API_KEY")
	serviceURL := os.Getenv("SERVICE_URL")

	logMessage("INFO", fmt.Sprintf("Environment Variable: %s", envVar))
	logMessage("INFO", fmt.Sprintf("API Key: %s", apiKey))
	logMessage("INFO", fmt.Sprintf("Service URL: %s", serviceURL))
}