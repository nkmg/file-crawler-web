package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route for Web App
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Add routers for Web app
func ConfigRoute(r *mux.Router) *mux.Router {
	router := routeSearching

	for _, route := range router {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
