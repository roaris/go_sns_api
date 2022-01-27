package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roaris/go_sns_api/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	v1r := r.PathPrefix("/api/v1").Subrouter()
	v1r.Methods(http.MethodGet).Path("/posts/{id:[0-9]+}").HandlerFunc(handlers.PostShow)
	v1r.Methods(http.MethodPost).Path("/posts").HandlerFunc(handlers.PostCreate)
	http.ListenAndServe(":8080", r)
}
