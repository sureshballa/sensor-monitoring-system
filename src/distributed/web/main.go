package main

import (
	"distributed/web/controller"
	"log"
	"net/http"
)

func main() {
	controller.Initialize()

	log.Fatal(http.ListenAndServe(":3000", nil))
}
