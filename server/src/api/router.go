package api

import (
	"github.com/gorilla/mux"
)

// GetRouter returns a new instance of a 'mux' router.
func GetRouter() *mux.Router {
	return mux.NewRouter()
}

// SetupRouterHandlers attaches the handlers to a router.
func SetupRouterHandlers(router *mux.Router) {
	router.HandleFunc(PingRoute, PingHandler).Methods("GET")
	router.HandleFunc(MessageRoute, ReceiveMessageHandler).Methods("POST")
	router.HandleFunc(MessageRoute, PreviousMessageHandler).Methods("GET")
}
