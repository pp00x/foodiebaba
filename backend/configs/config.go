package configs

import (
    "log"

    "github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }
}