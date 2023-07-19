package route

import (
	"api/src/controller"
	"net/http"
)

var routeSearching = []Route{
	{
		URI:     "/search",
		Method:  http.MethodPost,
		Handler: controller.SearchingFiles,
	},
	{
		URI:     "/search",
		Method:  http.MethodGet,
		Handler: controller.ReturnFiles,
	},
}
