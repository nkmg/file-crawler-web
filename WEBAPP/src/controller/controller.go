package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendSearchForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	research, err := json.Marshal(map[string]string{
		"path": r.FormValue("path"),
		"word": r.FormValue("word"),
	})

	if err != nil {
		log.Println(err)
	}

	response, err := http.Post("http://localhost:3000/searching", "application/json", bytes.NewBuffer(research))

	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	fmt.Println(response.Body)
}
