package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.GenerateRouters()

	log.Fatal(http.ListenAndServe(":3000", r))

}
