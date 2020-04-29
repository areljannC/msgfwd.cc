package main

import (
	"log"
	"net/http"
	"server/api"
)

func main() {
	// Initialize router.
	router := api.GetRouter()

	// Setup router handlers.
	api.SetupRouterHandlers(router)

	// Bind to a port and pass our router in.
	log.Fatal(http.ListenAndServe(":8000", router))
}
