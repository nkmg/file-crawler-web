package controller

import (
	"net/http"
	"webapp/src/utils"
)

// Load the initial page for Search
func LoadSearchInitial(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "search.html", nil)
}
