package main

import (
	"log"
	"net/http"

	"github.com/YahyaNashar22/calorie_tracker/internal/config"
)

func main() {
	// load all app config
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	log.Printf("Server running on port %s", config.AppConfig.Port)

	if err := http.ListenAndServe(":"+config.AppConfig.Port, mux); err != nil {
		log.Fatal(err)
	}
}
