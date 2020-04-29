package api

import (
	"encoding/json"
	"net/http"
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
	var parseInfo structs.MessageJSON
	json.NewDecoder(r.Body).Decode(&parseInfo)
	info := structs.MessageJSON{Name: parseInfo.Name, Message: parseInfo.Message, Location: parseInfo.Location, Timestamp: time.Now()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// PreviousMessageHandler parses a GET a request and sends back info about the last message to client.
func PreviousMessageHandler(w http.ResponseWriter, r *http.Request) {
	var parseInfo structs.MessageJSON
	json.NewDecoder(r.Body).Decode(&parseInfo)
	info := structs.MessageJSON{Name: parseInfo.Name, Message: parseInfo.Message, Location: parseInfo.Location, Timestamp: time.Now()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}
