package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
	}

	envVar := os.Getenv("ENV_VAR_NAME")
	apiKey := os.Getenv("API_KEY")
	serviceURL := os.Getenv("SERVICE_URL")

	fmt.Println("Environment Variable:", envVar)
	fmt.Println("API Key:", apiKey)
	fmt.Println("Service URL:", serviceURL)
}