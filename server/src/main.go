package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize 'mux'.
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/ping", PingHandler).Methods("GET")
	r.HandleFunc("/api/info", InfoHandler).Methods("POST")

	// Bind to a port and pass our router in.
	log.Fatal(http.ListenAndServe(":8000", r))
}

// YourHandler handles basic '/'.
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

// PingHandler returns a 200 status code.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	ping := pingJSON{StatusCode: 200, DateTime: time.Now(), Message: "Healthy!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ping)
}

// InfoHandler parses a POST a request and sends back data to client.
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var parseInfo infoJSON
	_ = json.NewDecoder(r.Body).Decode(&parseInfo)
	info := infoJSON{Name: parseInfo.Name}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

type pingJSON struct {
	StatusCode uint16    `json:"statusCode"`
	DateTime   time.Time `json:"dateTime"`
	Message    string    `json:"message"`
}

type infoJSON struct {
	Name string `json:"name"`
}
