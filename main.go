package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codeargent/go-cms/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	utils.InitializeEnvs()

	port := os.Getenv("PORT")

	log.Println("JSON API server running on port: ", port)

	error := http.ListenAndServe(port, router)

	if error != nil {
		log.Fatal("Failed to initialize api: \n", error)
	}
}
