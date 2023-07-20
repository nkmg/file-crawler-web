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
		URI:     "/",
		Method:  http.MethodPost,
		Handler: controller.SendSearchForm,
	},
}
