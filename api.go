package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIService struct {
	listenAddress string
}

func NewAPIService(listenAddress string) *APIService {
	return &APIService{
		listenAddress: listenAddress,
	}
}

func (s *APIService) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/request", makeHTTPHandleFunc(s.handleRequest))

	log.Println("JSON API server running on port: ", s.listenAddress)

	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIService) handleRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		return s.handleGetRequest(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIService) handleGetRequest(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, map[string]string{"getRequest": "RandomData"})
}
