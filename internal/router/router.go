package router

import "net/http"

func New() *http.ServeMux {
	mux := http.NewServeMux()

	registerHealthRoutes(mux)

	return mux
}
