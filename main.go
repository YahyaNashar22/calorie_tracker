package main

import (
	"log"

	"github.com/YahyaNashar22/calorie_tracker/internal/config"
	"github.com/YahyaNashar22/calorie_tracker/internal/database"
)

func main() {
	// load all app config
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(config.AppConfig.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := CreateServer()

	log.Printf("Server running on port %s", config.AppConfig.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
