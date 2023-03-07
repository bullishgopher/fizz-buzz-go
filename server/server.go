package server

import (
	"log"
	"os"
)

func Init() {
	os.Setenv("Fizz", "Fizz")
	os.Setenv("Buzz", "Buzz")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := SetupRouter()
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
