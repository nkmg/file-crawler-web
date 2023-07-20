package router

import (
	"webapp/src/router/route"

	"github.com/gorilla/mux"
)

// Return router with the routes
func GenerateRouters() *mux.Router {
	r := mux.NewRouter()
	return route.ConfigRoute(r)
}
