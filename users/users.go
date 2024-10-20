package users

import (
	"fmt"
	"net/http"

	"github.com/codeargent/go-cms/utils"
	"github.com/gorilla/mux"
)

func Main(router *mux.Router) {
	initializeRoutes(router)
}

func initializeRoutes(router *mux.Router) {
	router.HandleFunc("/users", utils.MakeHTTPHandleFunc(handleGetAllRequest))
	router.HandleFunc("/users/{id}", utils.MakeHTTPHandleFunc(handleRequestById))
}

func handleGetAllRequest(w http.ResponseWriter, r *http.Request) error {
	return utils.WriteJSON(w, http.StatusOK, map[string]string{"getAllRequest": "AllRequest"})
}

func handleRequestById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		id := mux.Vars(r)["id"]
		return utils.WriteJSON(w, http.StatusOK, map[string]string{"getRequestById": "RequestById: " + id})
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}
