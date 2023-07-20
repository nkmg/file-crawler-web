package route

import (
	"api/src/controller"
	"net/http"
)

var routeSearching = []Route{
	{
		URI:     "/searching",
		Method:  http.MethodPost,
		Handler: controller.SearchingFiles,
	},
	{
		URI:     "/searching",
		Method:  http.MethodGet,
		Handler: controller.ReturnFiles,
	},
}
