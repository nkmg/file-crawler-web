package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"
)

// Generate and return all routes
func GenerateRouters() *mux.Router {
	r := mux.NewRouter()
	return route.ConfigRoute(r)
}
