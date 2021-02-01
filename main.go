package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	. "api/provider"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
		log.Printf("Defaulting to port %s", port)
	}

	r := Routes()
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CompressHandler(r)))
}
