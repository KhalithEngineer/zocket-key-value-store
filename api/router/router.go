package router

import (
	"log"
	"net/http"
	"os"
	"zocket-key-value-store/api/handlers"
	"zocket-key-value-store/api/middleware"

	"github.com/hashicorp/consul/api"
)

// RegisterRoutes registers API routes
func RegisterRoutes(mux *http.ServeMux) {
	log.Print("Route")
	consulConfig := api.DefaultConfig()
	consulConfig.Address = os.Getenv("CONSUL_ENDPOINT")
	log.Println(consulConfig.Address)
	consulClient, _ := api.NewClient(consulConfig)

	// Create a new KeyValueHandler instance with the memberlist
	handler := handlers.NewKeyValueHandler(consulClient)

	// Apply middleware
	mux.Handle("/put", middleware.LoggerMiddleware(http.HandlerFunc(handler.PutHandler)))
	mux.Handle("/get", middleware.LoggerMiddleware(http.HandlerFunc(handler.GetHandler)))
	mux.Handle("/delete", middleware.LoggerMiddleware(http.HandlerFunc(handler.DeleteHandler)))
}
