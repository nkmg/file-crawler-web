package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Api route
type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

// Add routers on API
func ConfigRoute(r *mux.Router) *mux.Router {
	router := routeSearching

	for _, route := range router {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	return r
}
