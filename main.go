package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	portString := os.Getenv("PORT")

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	if portString == "" {
		log.Fatal("PORT is not found on the environment")
	}

	fmt.Println("Port: :", portString)
}
