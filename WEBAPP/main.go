package main

import (
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {

	utils.LoadTemplates()
	r := router.GenerateRouters()
	log.Fatal(http.ListenAndServe(":5000", r))
}
