package main

import (
	"log"

	"github.com/YahyaNashar22/calorie_tracker/internal/config"
)

func main() {
	// load all app config
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	server := CreateServer()

	log.Printf("Server running on port %s", config.AppConfig.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
