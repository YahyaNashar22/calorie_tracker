package router

import (
	"net/http"

	"github.com/YahyaNashar22/calorie_tracker/internal/handlers"
)

func registerHealthRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /health", handlers.Health)

}
