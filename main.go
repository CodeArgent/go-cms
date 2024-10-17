package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Println("JSON API server running on port: ", 3000)

	http.ListenAndServe(":3000", router)
}
