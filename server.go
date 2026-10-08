package main

import (
	"net/http"
	"time"

	"github.com/YahyaNashar22/calorie_tracker/internal/config"
	"github.com/YahyaNashar22/calorie_tracker/internal/router"
)

func CreateServer() *http.Server {
	server := &http.Server{
		Addr:              ":" + config.AppConfig.Port,
		Handler:           router.New(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	return server
}
