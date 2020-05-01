package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/services"
	"server/structs"
	"time"
)

// PingHandler returns an OK (200) status code.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	ping := structs.PingJSON{StatusCode: http.StatusOK, Message: "Healthy!", Timestamp: time.Now()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ping)
}

// ReceiveMessageHandler parses a POST a request and sends back data to client.
func ReceiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request body.
	var body structs.MessageJSON
	json.NewDecoder(r.Body).Decode(&body)

	// Create message struct.
	message := structs.MessageJSON{Name: body.Name, Location: body.Location, Message: body.Message, Timestamp: time.Now()}

	// Serialize struct.
	serialized, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
	}

	// Get Redis client and set latest message.
	redis := services.GetRedisClient()
	err = redis.Set("LATEST_MESSAGE", serialized, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// Return a response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// PreviousMessageHandler parses a GET a request and sends back info about the last message to client.
func PreviousMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Get Redis client.
	redis := services.GetRedisClient()

	// Get serialized latest message from Redis and convert it to bytes.
	val, err := redis.Get("LATEST_MESSAGE").Result()
	if err != nil {
		fmt.Println(err)
	}
	serialized := []byte(val)

	// Deserialize the latest message.
	var message structs.MessageJSON
	err = json.Unmarshal(serialized, &message)
	if err != nil {
		fmt.Println(err)
	}

	// Return a response.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
