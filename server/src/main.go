package main

import (
	"fmt"
	"log"
	"net/http"
	"server/api"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables.
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading environment variables.")
	}

	// Initialize router.
	router := api.GetRouter()

	// Setup router handlers.
	api.SetupRouterHandlers(router)

	// Bind to a port and pass our router in.
	log.Fatal(http.ListenAndServe(":8000", router))
}
