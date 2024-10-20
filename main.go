package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codeargent/go-cms/users"
	"github.com/codeargent/go-cms/utils"
	"github.com/gorilla/mux"
)

func main() {
	utils.InitializeEnvs()
	port := os.Getenv("PORT")

	router := mux.NewRouter()

	users.Main(router)

	error := http.ListenAndServe(port, router)

	if error == nil {
		fmt.Println("Server initialized sucessfully!")
	} else {
		log.Fatal("Failed to initialize api: \n", error)
	}
}
