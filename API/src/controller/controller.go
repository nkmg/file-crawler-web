package controller

import "net/http"

// Search the files
func SearchingFiles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST Search is working!!"))

}

// Return the files finded
func ReturnFiles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get search is working!!"))

}
