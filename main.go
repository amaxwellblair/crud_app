package main

import (
	"net/http"

	"github.com/amaxwellblair/crud_app/app/handlers"
)

func main() {
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/robots/", handler.RobotsHandler)
	http.ListenAndServe(":8080", nil)
}
