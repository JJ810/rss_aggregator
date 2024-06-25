package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	fmt.Println("PORT is set to", portString)
}
