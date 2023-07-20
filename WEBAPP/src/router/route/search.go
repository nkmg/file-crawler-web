package route

import (
	"net/http"
	"webapp/src/controller"
)

var routeSearching = []Route{
	{
		URI:     "/",
		Method:  http.MethodGet,
		Handler: controller.LoadSearchInitial,
	},
	{
		URI:     "/search",
		Method:  http.MethodGet,
		Handler: controller.LoadSearchInitial,
	},
}
