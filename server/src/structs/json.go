package structs

import "time"

// PingJSON represents an HTTP response to the client.
type PingJSON struct {
	StatusCode uint16    `json:"statusCode"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}

// MessageJSON represents an HTTP request from the client.
type MessageJSON struct {
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	Location  string    `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}
