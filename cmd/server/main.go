package main

import (
	"fmt"
	"log"
	"net/http"

	"dnd-sheet-api/internal/config"
	"dnd-sheet-api/internal/database"
	"dnd-sheet-api/internal/proxy"
	"dnd-sheet-api/internal/router"
)

func main() {
	configuration := config.Load()
	databaseConnection := database.Connect(configuration)
	dnd5eClient := proxy.NewDnd5eClient(configuration.DND5eAPIURL)

	httpRouter := router.New(databaseConnection, dnd5eClient)

	address := fmt.Sprintf(":%s", configuration.ServerPort)
	log.Printf("server running on http://localhost%s", address)

	if err := http.ListenAndServe(address, httpRouter); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
