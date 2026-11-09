// Command api runs the devboard HTTP API.
package main

import (
	"log"

	"github.com/amtzran/devboard/internal/handler"
	"github.com/amtzran/devboard/internal/server"
)

func main() {
	srv := server.New(":8000")

	healthHandler := handler.NewHealth()
	srv.RegisterRoute("GET /health", healthHandler)

	log.Println("Listening on :8000")
	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
