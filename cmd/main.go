package main

import (
	"app/internal/app"
	"app/internal/boot"
	"github.com/joho/godotenv"
	"log"
)

// fakhimi@gmail.com
func main() {
	// Todo:: move to bootstrapper
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := boot.Boot(); err != nil {
		panic(err)
	}

	if err := app.Serve(); err != nil {
		panic(err)
	}
}
